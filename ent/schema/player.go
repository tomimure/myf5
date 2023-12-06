package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Player holds the schema definition for the Player entity.
type Player struct {
	ent.Schema
}

// Fields of the Player.
func (Player) Fields() []ent.Field {
	return []ent.Field{
		field.Int("global"),
		field.String("name").Unique(),
	}
}

// Edges of the Player.
func (Player) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("teams", Team.Type).Ref("players"),
		edge.To("events", Event.Type),
	}
}

// Indexes of the Player.
func (Player) Indexes() []ent.Index {
	return nil
}
