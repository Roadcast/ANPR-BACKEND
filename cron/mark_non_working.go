package cron

import (
	"context"
	"go-ent-project/internal/ent"
	"go-ent-project/internal/ent/camera"
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
			if cam.LastPingTime == nil {
				_, err := client.Camera.UpdateOneID(cam.ID).
					SetIsWorking(false).
					Save(ctx)
				if err != nil {
					log.Printf("[%s] ❌ Failed to mark camera %s (IMEI: %s) as not working: %v", workerID, cam.Name, cam.Imei, err)
				} else {
					log.Printf("[%s] 🛠️ Camera marked not working: %s (IMEI: %s)", workerID, cam.Name, cam.Imei)
				}
			}

			if now.Sub(*cam.LastPingTime) > cutoffAge {
				_, err := client.Camera.UpdateOneID(cam.ID).
					SetIsWorking(false).
					Save(ctx)
				if err != nil {
					log.Printf("[%s] ❌ Failed to mark camera %s (IMEI: %s) as not working: %v", workerID, cam.Name, cam.Imei, err)
				} else {
					log.Printf("[%s] 🛠️ Camera marked not working: %s (IMEI: %s)", workerID, cam.Name, cam.Imei)
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
