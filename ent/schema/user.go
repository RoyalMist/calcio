package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique().Immutable(),
		field.String("password").NotEmpty().Sensitive(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

/*func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(mutator ent.Mutator) ent.Mutator {
			return hook.UserFunc(func(ctx context.Context, mutation *gen.UserMutation) (gen.Value, error) {
				return nil, nil
			})
		}, ent.OpCreate|ent.OpUpdate),
	}
}*/
