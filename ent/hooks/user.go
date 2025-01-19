package hook

import (
	"context"
	"entgo.io/ent"
	"errors"
	ent2 "go-ent-project/internal/ent"
	"go-ent-project/internal/ent/hook"
	"golang.org/x/crypto/bcrypt"
)

func HashPasswordHook() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return hook.UserFunc(func(ctx context.Context, mutation *ent2.UserMutation) (ent.Value, error) {
			// Placeholder for mutation logic
			if mutation == nil {
				return nil, errors.New("mutation is nil")
			}

			if password, exists := mutation.Password(); exists {
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

func MakeHooks() []ent.Hook {

	return []ent.Hook{
		hook.On(HashPasswordHook(), ent.OpCreate|ent.OpUpdate),
	}
}
