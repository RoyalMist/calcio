package security

import (
	"context"

	"calcio/ent/privacy"
)

// DenyIfNotLoggedIn is a rule that returns Allow decision if there is a connected user
func DenyIfNotLoggedIn() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		claims := FromContext(ctx)
		if len(claims.UserId) < 36 {
			return privacy.Denyf("user is missing from context")
		}

		return privacy.Skip
	})
}

// AllowIfAdmin is a rule that returns Allow decision if the user is admin.
func AllowIfAdmin() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		claims := FromContext(ctx)
		if claims.IsAdmin {
			return privacy.Allow
		}

		return privacy.Skip
	})
}
