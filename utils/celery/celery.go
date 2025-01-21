package celery

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	redisDB "go-ent-project/utils/redis"
	"log"
	"sync"
	"time"
)

const (
	MaxWorkers = 5  // Number of workers in the pool
	MaxQueue   = 10 // Maximum number of tasks in the queue
)

// BackgroundTask defines a function type for tasks with specific parameters.
type BackgroundTask[T TaskParams] func(params T) (interface{}, error)

// Task represents a task's metadata, status, result, or error.
type Task struct {
	ID        string      `json:"id"`
	Status    string      `json:"status"` // pending, running, completed, failed
	Result    interface{} `json:"result,omitempty"`
	Error     string      `json:"error,omitempty"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Attempts  int         `json:"attempts"`
}

// TaskManager manages tasks and their statuses.
type TaskManager struct {
	mu           sync.RWMutex
	tasks        map[string]*Task
	resultExpiry time.Duration
	taskQueue    chan func() // Task queue to hold tasks
	wg           sync.WaitGroup
}

var (
	tm *TaskManager
)

// NewTaskManager initializes a TaskManager with Redis client.
func NewTaskManager() *TaskManager {
	tm := &TaskManager{
		tasks:        make(map[string]*Task),
		resultExpiry: 15 * time.Minute,
		taskQueue:    make(chan func(), MaxQueue),
	}

	// Start worker pool
	tm.startWorkerPool(MaxWorkers)
	return tm
}

// startWorkerPool initializes the worker pool with a fixed number of workers.
func (tm *TaskManager) startWorkerPool(numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		tm.wg.Add(1)
		go tm.worker(i)
	}
}

// worker continuously processes tasks from the queue.
func (tm *TaskManager) worker(workerID int) {
	defer tm.wg.Done()
	log.Printf("Worker %d started", workerID)

	for task := range tm.taskQueue {
		task() // Execute the task
	}

	log.Printf("Worker %d stopped", workerID)
}

// EnqueueTask adds a task to the task queue for execution.
func (tm *TaskManager) EnqueueTask(task func()) {
	select {
	case tm.taskQueue <- task:
		log.Printf("Task added to the queue")
	default:
		log.Printf("Task queue is full, dropping task")
	}
}

// Shutdown waits for all workers to finish and closes the task queue.
func (tm *TaskManager) Shutdown() {
	close(tm.taskQueue)
	tm.wg.Wait()
	log.Println("All workers stopped")
}

// CreateTask initializes a new task and returns its ID.
func (tm *TaskManager) CreateTask() string {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	taskID := uuid.New().String()
	task := &Task{
		ID:        taskID,
		Status:    "pending",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Attempts:  0,
	}
	tm.tasks[taskID] = task
	return taskID
}

// UpdateTaskStatus updates a task's status, result, or error.
func (tm *TaskManager) UpdateTaskStatus(id, status string, result interface{}, err error) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	task, exists := tm.tasks[id]
	if !exists {
		return
	}

	task.Status = status
	if err != nil {
		task.Error = err.Error()
	} else {
		task.Result = result
	}
	task.Attempts++
	task.UpdatedAt = time.Now()

	// Save task to Redis
	tm.saveTaskToRedis(task)
}

func (tm *TaskManager) saveTaskToRedis(task *Task) {
	if redisDB.Client == nil {
		log.Printf("Redis client is not initialized. Cannot save task %s to Redis.", task.ID)
		return
	}

	ctx := context.Background()
	key := fmt.Sprintf("task:%s", task.ID)
	fmt.Printf("Saving task %s to Redis\n", task.ID)
	fmt.Printf("Status: %s, Result: %v, Error: %v\n", task.Status, task.Result, task.Error)

	// Convert result to string if not nil
	var result string
	if task.Result != nil {
		result = fmt.Sprintf("%v", task.Result)
	}

	// Save to Redis
	err := redisDB.Client.HSet(ctx, key, map[string]interface{}{
		"status":    task.Status,
		"result":    result,
		"error":     task.Error,
		"createdAt": task.CreatedAt.Format(time.RFC3339),
		"updatedAt": task.UpdatedAt.Format(time.RFC3339),
		"attempts":  task.Attempts,
	}).Err()

	redisDB.Client.Expire(ctx, key, tm.resultExpiry*time.Minute)

	if err != nil {
		log.Printf("Failed to save task to Redis: %v", err)
	}
}

// GetTaskFromRedis queries Redis by task ID and retrieves the task details.
func (tm *TaskManager) GetTaskFromRedis(taskID string) (*Task, error) {
	if redisDB.Client == nil {
		return nil, fmt.Errorf("Redis client is not initialized")
	}

	ctx := context.Background()
	key := fmt.Sprintf("task:%s", taskID)

	// Retrieve task details from Redis
	data, err := redisDB.Client.HGetAll(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("task with ID %s not found", taskID)
		}
		return nil, fmt.Errorf("failed to fetch task from Redis: %v", err)
	}

	// Check if the task data exists
	if len(data) == 0 {
		return nil, fmt.Errorf("task with ID %s not found", taskID)
	}

	// Parse the task details
	task := &Task{
		ID:        taskID,
		Status:    data["status"],
		Result:    data["result"],
		Error:     data["error"],
		CreatedAt: parseTime(data["createdAt"]),
		UpdatedAt: parseTime(data["updatedAt"]),
	}
	if attempts, ok := data["attempts"]; ok {
		_, err := fmt.Sscanf(attempts, "%d", &task.Attempts)
		if err != nil {
			return nil, err
		}
	}

	return task, nil
}

// Helper function to parse time strings from Redis
func parseTime(timeStr string) time.Time {
	if timeStr == "" {
		return time.Time{}
	}
	t, _ := time.Parse(time.RFC3339, timeStr)
	return t
}

// ExecuteTask executes a BackgroundTask with its parameters in a goroutine.
func ExecuteTask[T TaskParams](taskFn BackgroundTask[T], params T, maxRetries int) {
	// Compile-time type check ensures params implement TaskParams.
	id := tm.CreateTask()
	if err := params.Validate(); err != nil {
		tm.UpdateTaskStatus(id, "failed", nil, err)
		return
	}

	tm.UpdateTaskStatus(id, "running", nil, nil)

	tm.EnqueueTask(func() {
		var result interface{}
		var err error
		for attempts := 0; attempts <= maxRetries; attempts++ {
			result, err = taskFn(params)
			if err == nil {
				tm.UpdateTaskStatus(id, "completed", result, nil)
				return
			}
			log.Printf("Task %s failed on attempt %d: %v", id, attempts+1, err)
			time.Sleep(1 * time.Second) // Optional: Add a delay between retries
		}
		tm.UpdateTaskStatus(id, "failed", nil, err)
	})

}

func Initialize() {
	// Example: Run UserNotificationTask with parameters
	tm = NewTaskManager()

	// How to push task in queue
	//taskID := tm.CreateTask()
	//ExecuteTask(tm, taskID, UserNotificationTask, params, 3)
	// Example: Parameters for a Task
	//type UserNotificationParams struct {
	//	BaseParams
	//	UserID     int
	//	Message    string
	//	Email      string
	//	RetryCount int
	//}

	// Validate uses the generic ValidateStruct for validation
	//func (p *UserNotificationParams) Validate() error {
	//	return ValidateStruct(p)
	//}

}

func ShutDown() {
	defer tm.Shutdown()
}
