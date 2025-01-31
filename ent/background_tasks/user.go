package background_tasks

import (
	"fmt"
	"github.com/google/uuid"
	"go-ent-project/utils/celery"
	"time"
)

type UserNotificationParams struct {
	celery.BaseParams
	UserID  uuid.UUID
	Message string
	Email   string
}

func (p *UserNotificationParams) Validate() error {
	return celery.ValidateStruct(p)
}

func SendEmailToUserHook(params *UserNotificationParams) (interface{}, error) {
	time.Sleep(10 * time.Second)
	fmt.Printf("Sending email to user %d at %s: %s\n", params.UserID, params.Email, params.Message)

	// Send email to user
	return nil, nil
}
