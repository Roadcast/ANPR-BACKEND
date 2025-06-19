package background_tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"go-ent-project/internal/ent"
	"go-ent-project/internal/ent/car"
	"go-ent-project/internal/ent/policestation"
	"go-ent-project/internal/ent/user"
)

func init() {
	_ = godotenv.Load() // Load env vars like ZeptoMail token
}

func CarHandler(client *ent.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		var body struct {
			PlateNumber string `json:"plate_number"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.Background()
		carObj, err := client.Car.Query().
			Where(car.RegistrationEQ(body.PlateNumber)).
			WithPoliceStation().
			Only(ctx)
		if err != nil {
			log.Printf("Car not found: %v", err)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("Plate not found or not registered as stolen."))
			return
		}

		station := carObj.Edges.PoliceStation
		if station == nil {
			log.Printf("Car found but has no linked police station")
			return
		}

		users, err := client.User.Query().
			Where(user.HasPoliceStationWith(policestation.ID(station.ID))).
			All(ctx)
		if err != nil {
			log.Printf("Error fetching users for station %s: %v", station.Name, err)
			return
		}

		log.Printf("🚨 Plate %s is stolen. Notifying users of station %s (%d user(s))", body.PlateNumber, station.Name, len(users))
		SendZeptoEmails(users, body.PlateNumber, station.Name)

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Alert processed for plate: " + body.PlateNumber))
	}
}

func SendZeptoEmails(users []*ent.User, plate string, station string) {
	apiKey := os.Getenv("ZEPTO_MAIL_ACCESS_TOKEN")
	domain := os.Getenv("ZEPTO_MAIL_FROM_DOMAIN")
	from := map[string]string{"address": domain}

	tos := []map[string]map[string]string{}
	for _, u := range users {
		if strings.TrimSpace(u.Email) == "" {
			continue
		}
		tos = append(tos, map[string]map[string]string{
			"email_address": {"address": u.Email, "name": u.Name},
		})
	}

	if len(tos) == 0 {
		log.Println("No users with valid emails to notify.")
		return
	}

	payload := map[string]interface{}{
		"from":     from,
		"to":       tos,
		"subject":  fmt.Sprintf("Stolen Car Alert: %s", plate),
		"htmlbody": fmt.Sprintf("<div><b>Plate Number: %s</b><br>Linked Station: %s</div>", plate, station),
	}

	jsonBody, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", "https://api.zeptomail.com/v1.1/email", strings.NewReader(string(jsonBody)))
	if err != nil {
		log.Printf("Failed to create email request: %v", err)
		return
	}
	req.Header.Set("Authorization", apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		return
	}
	defer resp.Body.Close()
	log.Printf("📧 Email alert sent for %s — Status: %s", plate, resp.Status)
}
