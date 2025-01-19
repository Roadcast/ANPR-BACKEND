package privacy

import (
	"context"
	"entgo.io/ent/entql"
	"fmt"
	ent2 "go-ent-project/internal/ent"
	"go-ent-project/internal/ent/privacy"
	"go-ent-project/utils/constant"

	"entgo.io/ent"
)

// AllowIfBypass checks if privacy should be bypassed.
func AllowIfBypass() privacy.QueryRule {
	return privacy.QueryRuleFunc(func(ctx context.Context, q ent.Query) error {
		fmt.Printf("AllowIfBypass %T\n", q)
		if bypass, ok := ctx.Value(constant.BypassPrivacyKey).(bool); ok && bypass {
			fmt.Printf("AllowIfBypass %T\n", bypass)
			return privacy.Allow
		}
		fmt.Printf("AllowIfBypass NOT")
		return privacy.Skip
	})
}

// FilterTenantRule is a query/mutation rule that filters out entities that are not in the tenant.
func FilterTenantRule() privacy.QueryMutationRule {
	// TenantsFilter is an interface to wrap WhereHasTenantWith()
	// predicate that is used by both `Group` and `User` schemas.
	type UserFiler interface {
		WhereUserID(entql.IntP)
	}
	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
		fmt.Printf("FilterTenantRule %T\n", f)
		fmt.Printf("FilterTenantRule %T\n", ctx)
		user, ok := ctx.Value("user").(ent2.User)
		tid := user.ID
		if !ok {
			fmt.Printf("FilterTenantRule %T\n", user)
			return privacy.Denyf("missing tenant information in viewer")
		}
		tf, ok := f.(UserFiler)
		if !ok {
			return privacy.Denyf("unexpected filter type %T", f)
		}
		// Make sure that a tenant reads only entities that have an edge to it.
		tf.WhereUserID(entql.IntEQ(tid))
		// Skip to the next privacy rule (equivalent to return nil).
		return privacy.Skip
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

// DenyRoleChange ensures that only admins can update the `role` field.
func AllowMutationByPass() privacy.MutationRule {
	return privacy.MutationRuleFunc(func(ctx context.Context, q ent.Mutation) error {
		if bypass, ok := ctx.Value(constant.BypassPrivacyKey).(bool); ok && bypass {
			return privacy.Allow
		}
		return privacy.Skip
	})
}

// Policy defines the privacy rules for the User entity.
var UserPolicy = privacy.Policy{
	Mutation: privacy.MutationPolicy{
		AllowMutationByPass(),
		// Apply AllowIfSelf to update mutations
		privacy.OnMutationOperation(AllowIfAdminOrManager(), ent.OpUpdate|ent.OpUpdateOne),

		// Apply DenyRoleChange to prevent unauthorized role changes
		privacy.OnMutationOperation(DenyRoleChange(), ent.OpUpdate|ent.OpUpdateOne),
	},
	Query: privacy.QueryPolicy{
		AllowIfBypass(),
		FilterTenantRule(),
		// Allow all users to read the `name` field
	},
}
