package privacy

import (
	"context"
	ent2 "go-ent-project/internal/ent"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// AllowIfAdmin ensures only admin users can perform certain actions.
func AllowIfAdmin() privacy.MutationRule {
	return privacy.MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {
		user, ok := ctx.Value("user").(ent2.User)
		if !ok || user.RoleID != 1 {
			return privacy.Denyf("only admins can perform this action")
		}
		return privacy.Allow
	})
}

// AllowIfAdminOrManager ensures admins and managers can perform certain actions.
func AllowIfAdminOrManager() privacy.MutationRule {
	return privacy.MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {
		user, ok := ctx.Value("user").(ent2.User)
		if !ok || (user.RoleID != 1 && user.RoleID != 2) {
			return privacy.Denyf("only admins or managers can perform this action")
		}
		return privacy.Allow
	})
}

// DenyRoleChange ensures that only admins can update the `role` field.
func DenyRoleChange() privacy.MutationRule {
	return privacy.MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {
		user, ok := ctx.Value("user").(ent2.User)
		if !ok || user.RoleID != 1 {
			return privacy.Denyf("only admins can update the role field")
		}

		// Check if the `role` field is being updated
		if _, exists := m.Field("role_id"); exists {
			return privacy.Allow
		}

		return privacy.Skip
	})
}

// Policy defines the privacy rules for the User entity.
var UserPolicy = privacy.Policy{
	Mutation: privacy.MutationPolicy{
		// Apply AllowIfSelf to update mutations
		privacy.OnMutationOperation(AllowIfAdminOrManager(), ent.OpUpdate|ent.OpUpdateOne),

		// Apply DenyRoleChange to prevent unauthorized role changes
		privacy.OnMutationOperation(DenyRoleChange(), ent.OpUpdate|ent.OpUpdateOne),
	},
	Query: privacy.QueryPolicy{

		// Allow all users to read the `name` field
	},
}
