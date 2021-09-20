package schema

import (
	"context"
	"fmt"

	"calcio/api/services/auth"
	gen "calcio/ent"
	"calcio/ent/hook"
	"entgo.io/ent"
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
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").NotEmpty().Unique().Immutable(),
		field.String("password").NotEmpty().Sensitive(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(mutator ent.Mutator) ent.Mutator {
			return hook.UserFunc(func(ctx context.Context, mutation *gen.UserMutation) (gen.Value, error) {
				if password, exists := mutation.Password(); exists {
					if len(password) < PasswordMinLength {
						return nil, fmt.Errorf("password too short, minimum length of %d", PasswordMinLength)
					}

					hash, err := auth.HashPassword(password)
					if err != nil {
						return nil, fmt.Errorf("unable to hash password")
					}

					mutation.SetPassword(hash)
					return mutator.Mutate(ctx, mutation)
				}

				return nil, fmt.Errorf("password was not set")
			})
		}, ent.OpCreate|ent.OpUpdate),
	}
}
