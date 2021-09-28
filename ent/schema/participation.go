package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Participation holds the schema definition for the Participation entity.
type Participation struct {
	ent.Schema
}

// Fields of the Participation.
func (Participation) Fields() []ent.Field {
	return []ent.Field{
		field.Int("goals").Default(0).NonNegative(),
	}
}

// Edges of the Participation.
func (Participation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("game", Game.Type).Unique().Required(),
		edge.To("team", Team.Type).Unique().Required(),
	}
}