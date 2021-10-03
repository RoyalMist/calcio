package schema

import (
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

func (Team) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			security.DenyIfNotLoggedIn(),
			privacy.AlwaysAllowRule(),
		},
		Mutation: privacy.MutationPolicy{
			security.DenyIfNotLoggedIn(),
			privacy.AlwaysAllowRule(),
		},
	}
}
