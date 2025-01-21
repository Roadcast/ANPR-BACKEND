package hook

import (
	"context"
	"entgo.io/ent"
	"errors"
	"fmt"
	ent2 "go-ent-project/internal/ent"
	"go-ent-project/internal/ent/hook"
	"golang.org/x/crypto/bcrypt"
)

func HashPasswordHook() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		fmt.Printf("HashPasswordHook\n")
		return hook.UserFunc(func(ctx context.Context, mutation *ent2.UserMutation) (ent.Value, error) {
			// Placeholder for mutation logic
			if mutation == nil {
				fmt.Printf("HashPasswordHook mutation is nil\n")
				return nil, errors.New("mutation is nil")
			}
			if password, exists := mutation.Password(); exists {
				fmt.Printf("HashPasswordHook password exists\n")
				hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
				if err != nil {
					return nil, err
				}
				mutation.SetPassword(string(hashedPassword))
			}
			return next.Mutate(ctx, mutation)
		})
	}
}

//func SendEmailToUserHook() ent.Hook {
//	return func(next ent.Mutator) ent.Mutator {
//
//		return hook.UserFunc(func(ctx context.Context, mutation *ent2.UserMutation) (ent.Value, error) {
//			// Placeholder for mutation logic
//			userID, _ := mutation.ID()
//			email, _ := mutation.Email()
//			params := &background_tasks.UserNotificationParams{
//				UserID:  userID,
//				Message: fmt.Sprintf("Welcome to our platform!"),
//				Email:   email,
//			}
//			celery.ExecuteTask(background_tasks.SendEmailToUserHook, params, 3)
//			return next.Mutate(ctx, mutation)
//		})
//	}
//}

func MakeHooks() []ent.Hook {

	return []ent.Hook{
		hook.On(HashPasswordHook(), ent.OpCreate|ent.OpUpdate),
		//hook.On(SendEmailToUserHook(), ent.OpCreate),
	}
}
