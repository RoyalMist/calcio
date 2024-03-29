package schema

import (
	"context"
	"fmt"

	gen "calcio/ent"
	"calcio/ent/hook"
	"calcio/ent/privacy"
	"calcio/server/security"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

const PasswordMinLength = 8

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.String("name").NotEmpty().Unique().Immutable(),
		field.String("password").NotEmpty().Sensitive(),
		field.Bool("admin").Default(false).Immutable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("teams", Team.Type).Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}),
	}
}

func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(mutator ent.Mutator) ent.Mutator {
			return hook.UserFunc(func(ctx context.Context, mutation *gen.UserMutation) (gen.Value, error) {
				if password, exists := mutation.Password(); exists {
					if len(password) < PasswordMinLength {
						return nil, fmt.Errorf("password too short, minimum length of %d", PasswordMinLength)
					}

					hash, err := security.HashPassword(password)
					if err != nil {
						return nil, fmt.Errorf("unable to hash password")
					}

					mutation.SetPassword(hash)
					return mutator.Mutate(ctx, mutation)
				}

				return mutator.Mutate(ctx, mutation)
			})
		}, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne),
	}
}

func (User) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			security.DenyIfNotLoggedIn(),
			privacy.AlwaysAllowRule(),
		},
		Mutation: privacy.MutationPolicy{
			security.DenyIfNotLoggedIn(),
			security.AllowIfAdmin(),
			privacy.AlwaysDenyRule(),
		},
	}
}
