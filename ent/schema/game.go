package schema

import (
	"time"

	"calcio/ent/privacy"
	"calcio/server/security"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Game holds the schema definition for the Game entity.
type Game struct {
	ent.Schema
}

// Fields of the Game.
func (Game) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.Time("date").Default(time.Now).Immutable(),
	}
}

// Edges of the Game.
func (Game) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("participations", Participation.Type).Ref("game"),
	}
}

func (Game) Policy() ent.Policy {
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
