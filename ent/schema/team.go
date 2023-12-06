package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Team holds the schema definition for the Team entity.
type Team struct {
	ent.Schema
}

// Fields of the Team.
func (Team) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("team").Values("1", "2"),
		field.Int("match_id"),
	}
}

// Edges of the Team.
func (Team) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("players", Player.Type).
			Required(),
		edge.From("match", Match.Type).
			Ref("teams").
			Unique().
			Required().
			Field("match_id"),
	}
}
