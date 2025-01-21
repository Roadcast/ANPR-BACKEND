package background_tasks

import (
	"fmt"
	"go-ent-project/utils/celery"
	"time"
)

type UserNotificationParams struct {
	celery.BaseParams
	UserID  int
	Message string
	Email   string
}

func (p *UserNotificationParams) Validate() error {
	return celery.ValidateStruct(p)
}

func SendEmailToUserHook(params *UserNotificationParams) {
	fmt.Printf("Sending email to user %d at %s: %s\n", params.UserID, params.Email, params.Message)
	time.Sleep(10 * time.Second)
	// Send email to user
}
