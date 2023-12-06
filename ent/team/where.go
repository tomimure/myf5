// Code generated by ent, DO NOT EDIT.

package team

import (
	"myf5/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Team {
	return predicate.Team(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Team {
	return predicate.Team(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Team {
	return predicate.Team(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Team {
	return predicate.Team(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Team {
	return predicate.Team(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Team {
	return predicate.Team(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Team {
	return predicate.Team(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Team {
	return predicate.Team(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Team {
	return predicate.Team(sql.FieldLTE(FieldID, id))
}

// MatchID applies equality check predicate on the "match_id" field. It's identical to MatchIDEQ.
func MatchID(v int) predicate.Team {
	return predicate.Team(sql.FieldEQ(FieldMatchID, v))
}

// TeamEQ applies the EQ predicate on the "team" field.
func TeamEQ(v Team) predicate.Team {
	return predicate.Team(sql.FieldEQ(FieldTeam, v))
}

// TeamNEQ applies the NEQ predicate on the "team" field.
func TeamNEQ(v Team) predicate.Team {
	return predicate.Team(sql.FieldNEQ(FieldTeam, v))
}

// TeamIn applies the In predicate on the "team" field.
func TeamIn(vs ...Team) predicate.Team {
	return predicate.Team(sql.FieldIn(FieldTeam, vs...))
}

// TeamNotIn applies the NotIn predicate on the "team" field.
func TeamNotIn(vs ...Team) predicate.Team {
	return predicate.Team(sql.FieldNotIn(FieldTeam, vs...))
}

// MatchIDEQ applies the EQ predicate on the "match_id" field.
func MatchIDEQ(v int) predicate.Team {
	return predicate.Team(sql.FieldEQ(FieldMatchID, v))
}

// MatchIDNEQ applies the NEQ predicate on the "match_id" field.
func MatchIDNEQ(v int) predicate.Team {
	return predicate.Team(sql.FieldNEQ(FieldMatchID, v))
}

// MatchIDIn applies the In predicate on the "match_id" field.
func MatchIDIn(vs ...int) predicate.Team {
	return predicate.Team(sql.FieldIn(FieldMatchID, vs...))
}

// MatchIDNotIn applies the NotIn predicate on the "match_id" field.
func MatchIDNotIn(vs ...int) predicate.Team {
	return predicate.Team(sql.FieldNotIn(FieldMatchID, vs...))
}

// HasPlayers applies the HasEdge predicate on the "players" edge.
func HasPlayers() predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, PlayersTable, PlayersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlayersWith applies the HasEdge predicate on the "players" edge with a given conditions (other predicates).
func HasPlayersWith(preds ...predicate.Player) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		step := newPlayersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMatch applies the HasEdge predicate on the "match" edge.
func HasMatch() predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MatchTable, MatchColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMatchWith applies the HasEdge predicate on the "match" edge with a given conditions (other predicates).
func HasMatchWith(preds ...predicate.Match) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		step := newMatchStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Team) predicate.Team {
	return predicate.Team(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Team) predicate.Team {
	return predicate.Team(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Team) predicate.Team {
	return predicate.Team(sql.NotPredicates(p))
}