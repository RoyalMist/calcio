package schema

import (
	"context"
	"fmt"

	gen "calcio/ent"
	"calcio/ent/hook"
	"calcio/ent/privacy"
	"calcio/server/security"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Team holds the schema definition for the Team entity.
type Team struct {
	ent.Schema
}

// Fields of the Team.
func (Team) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.String("name").NotEmpty().Unique(),
	}
}

// Edges of the Team.
func (Team) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("players", User.Type).Ref("teams").Required(),
	}
}

func (Team) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(mutator ent.Mutator) ent.Mutator {
			return hook.TeamFunc(func(ctx context.Context, mutation *gen.TeamMutation) (gen.Value, error) {
				if name, exists := mutation.Name(); exists {
					fmt.Printf("Hello %s", name)
					return mutator.Mutate(ctx, mutation)
				}

				return mutator.Mutate(ctx, mutation)
			})
		}, ent.OpCreate|ent.OpUpdate),
	}
}

func (Team) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			security.DenyIfNotLoggedIn(),
			privacy.AlwaysDenyRule(),
		},
		Mutation: privacy.MutationPolicy{
			security.DenyIfNotLoggedIn(),
			security.AllowIfAdmin(),
			privacy.AlwaysDenyRule(),
		},
	}
}
