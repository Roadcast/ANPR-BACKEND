package hook

import (
	"context"
	"entgo.io/ent"
	ent2 "go-ent-project/internal/ent"
	"go-ent-project/internal/ent/hook"
	"go-ent-project/utils/redis"
	"log"
)

// Hook to add car plate to Redis on create
func AddCarToRedisHook() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return hook.CarFunc(func(ctx context.Context, mutation *ent2.CarMutation) (ent.Value, error) {
			// Get registration plate from mutation
			plate, exists := mutation.Registration()
			if exists {
				err := AddCarPlate(ctx, plate)
				if err != nil {
					log.Printf("Failed to add car plate to Redis: %v", err)
				}
			}
			return next.Mutate(ctx, mutation)
		})
	}
}

// Hook to remove car plate from Redis on delete
func RemoveCarFromRedisHook() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return hook.CarFunc(func(ctx context.Context, mutation *ent2.CarMutation) (ent.Value, error) {
			id, exists := mutation.ID()
			if !exists {
				return next.Mutate(ctx, mutation)
			}
			// Fetch car from DB
			client := ent2.FromContext(ctx)
			car, err := client.Car.Get(ctx, id)
			if err != nil {
				log.Printf("Failed to fetch car before delete: %v", err)
				return next.Mutate(ctx, mutation)
			}
			// Remove plate from Redis
			err = RemoveCarPlate(ctx, car.Registration)
			if err != nil {
				log.Printf("Failed to remove car plate from Redis: %v", err)
			}
			return next.Mutate(ctx, mutation)
		})
	}
}

func AddCarPlate(ctx context.Context, plate string) error {
	return redis.Client.SAdd(ctx, "car:plates", plate).Err()
}

func RemoveCarPlate(ctx context.Context, plate string) error {
	return redis.Client.SRem(ctx, "car:plates", plate).Err()
}

func CarHooks() []ent.Hook {

	return []ent.Hook{
		hook.On(AddCarToRedisHook(), ent.OpCreate),
		hook.On(RemoveCarFromRedisHook(), ent.OpDelete),
	}
}
