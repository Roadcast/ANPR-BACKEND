package cron

import (
	"context"
	"go-ent-project/ent/background_tasks"
	"go-ent-project/internal/ent"
	"go-ent-project/internal/ent/camera"
	"go-ent-project/internal/ent/policestation"
	"go-ent-project/internal/ent/user"
	"log"
	"time"
)

const (
	checkInterval = 5 * time.Minute
	cutoffAge     = 15 * time.Minute
)

func StartCameraHealthCheck(client *ent.Client, workerID string) {
	ticker := time.NewTicker(checkInterval)
	defer ticker.Stop()

	for {
		<-ticker.C
		ctx := context.Background()
		log.Printf("[%s] ⏰ Starting camera health check...", workerID)

		// Fetch all active cameras
		cameras, err := client.Camera.Query().
			Where(camera.Active(true)).
			All(ctx)
		if err != nil {
			log.Printf("[%s] ❌ Error fetching cameras: %v", workerID, err)
			continue
		}

		now := time.Now()
		for _, cam := range cameras {
			prevIsWorking := cam.IsWorking
			if cam.LastPingTime == nil {
				_, err := client.Camera.UpdateOneID(cam.ID).
					SetIsWorking(false).
					Save(ctx)
				if err != nil {
					log.Printf("[%s] ❌ Failed to mark camera %s (IMEI: %s) as not working: %v", workerID, cam.Name, cam.Imei, err)
				} else {
					log.Printf("[%s] 🛠️ Camera marked not working: %s (IMEI: %s)", workerID, cam.Name, cam.Imei)
					if prevIsWorking {
						go notifyCameraNonWorkingUsers(client, ctx, cam)
					}
				}
			} else {
				if now.Sub(*cam.LastPingTime) > cutoffAge {
					_, err := client.Camera.UpdateOneID(cam.ID).
						SetIsWorking(false).
						Save(ctx)
					if err != nil {
						log.Printf("[%s] ❌ Failed to mark camera %s (IMEI: %s) as not working: %v", workerID, cam.Name, cam.Imei, err)
					} else {
						log.Printf("[%s] 🛠️ Camera marked not working: %s (IMEI: %s)", workerID, cam.Name, cam.Imei)
						if prevIsWorking {
							go notifyCameraNonWorkingUsers(client, ctx, cam)
						}
					}
				} else {
					_, _ = client.Camera.UpdateOneID(cam.ID).
						SetIsWorking(true).
						Save(ctx)
					log.Printf("[%s] ✅ Camera is working: %s (IMEI: %s)", workerID, cam.Name, cam.Imei)
				}
			}
		}
	}
}

// notifyCameraNonWorkingUsers sends an email to users if a camera goes from working to non-working
func notifyCameraNonWorkingUsers(client *ent.Client, ctx context.Context, cam *ent.Camera) {
	// Fetch police station
	if cam.PoliceStationID == nil {
		log.Printf("Camera %s has no linked police station, cannot notify users.", cam.Name)
		return
	}
	station, err := client.PoliceStation.Get(ctx, *cam.PoliceStationID)
	if err != nil {
		log.Printf("Failed to fetch police station for camera %s: %v", cam.Name, err)
		return
	}
	// Fetch users for the station
	users, err := client.User.Query().Where(user.HasPoliceStationWith(policestation.IDEQ(station.ID))).All(ctx)
	if err != nil {
		log.Printf("Failed to fetch users for station %s: %v", station.Name, err)
		return
	}
	if len(users) == 0 {
		log.Printf("No users to notify for station %s", station.Name)
		return
	}
	// Send email (reuse sendZeptoEmails logic, or call a similar function)
	background_tasks.SendZeptoEmails(users, cam.Name, station.Name)
}
