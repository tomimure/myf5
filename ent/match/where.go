// Code generated by ent, DO NOT EDIT.

package match

import (
	"myf5/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Match {
	return predicate.Match(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Match {
	return predicate.Match(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Match {
	return predicate.Match(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Match {
	return predicate.Match(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Match {
	return predicate.Match(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Match {
	return predicate.Match(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Match {
	return predicate.Match(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Match {
	return predicate.Match(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Match {
	return predicate.Match(sql.FieldLTE(FieldID, id))
}

// Date applies equality check predicate on the "date" field. It's identical to DateEQ.
func Date(v time.Time) predicate.Match {
	return predicate.Match(sql.FieldEQ(FieldDate, v))
}

// GoalsTeam1 applies equality check predicate on the "goalsTeam1" field. It's identical to GoalsTeam1EQ.
func GoalsTeam1(v int) predicate.Match {
	return predicate.Match(sql.FieldEQ(FieldGoalsTeam1, v))
}

// GoalsTeam2 applies equality check predicate on the "goalsTeam2" field. It's identical to GoalsTeam2EQ.
func GoalsTeam2(v int) predicate.Match {
	return predicate.Match(sql.FieldEQ(FieldGoalsTeam2, v))
}

// DateEQ applies the EQ predicate on the "date" field.
func DateEQ(v time.Time) predicate.Match {
	return predicate.Match(sql.FieldEQ(FieldDate, v))
}

// DateNEQ applies the NEQ predicate on the "date" field.
func DateNEQ(v time.Time) predicate.Match {
	return predicate.Match(sql.FieldNEQ(FieldDate, v))
}

// DateIn applies the In predicate on the "date" field.
func DateIn(vs ...time.Time) predicate.Match {
	return predicate.Match(sql.FieldIn(FieldDate, vs...))
}

// DateNotIn applies the NotIn predicate on the "date" field.
func DateNotIn(vs ...time.Time) predicate.Match {
	return predicate.Match(sql.FieldNotIn(FieldDate, vs...))
}

// DateGT applies the GT predicate on the "date" field.
func DateGT(v time.Time) predicate.Match {
	return predicate.Match(sql.FieldGT(FieldDate, v))
}

// DateGTE applies the GTE predicate on the "date" field.
func DateGTE(v time.Time) predicate.Match {
	return predicate.Match(sql.FieldGTE(FieldDate, v))
}

// DateLT applies the LT predicate on the "date" field.
func DateLT(v time.Time) predicate.Match {
	return predicate.Match(sql.FieldLT(FieldDate, v))
}

// DateLTE applies the LTE predicate on the "date" field.
func DateLTE(v time.Time) predicate.Match {
	return predicate.Match(sql.FieldLTE(FieldDate, v))
}

// GoalsTeam1EQ applies the EQ predicate on the "goalsTeam1" field.
func GoalsTeam1EQ(v int) predicate.Match {
	return predicate.Match(sql.FieldEQ(FieldGoalsTeam1, v))
}

// GoalsTeam1NEQ applies the NEQ predicate on the "goalsTeam1" field.
func GoalsTeam1NEQ(v int) predicate.Match {
	return predicate.Match(sql.FieldNEQ(FieldGoalsTeam1, v))
}

// GoalsTeam1In applies the In predicate on the "goalsTeam1" field.
func GoalsTeam1In(vs ...int) predicate.Match {
	return predicate.Match(sql.FieldIn(FieldGoalsTeam1, vs...))
}

// GoalsTeam1NotIn applies the NotIn predicate on the "goalsTeam1" field.
func GoalsTeam1NotIn(vs ...int) predicate.Match {
	return predicate.Match(sql.FieldNotIn(FieldGoalsTeam1, vs...))
}

// GoalsTeam1GT applies the GT predicate on the "goalsTeam1" field.
func GoalsTeam1GT(v int) predicate.Match {
	return predicate.Match(sql.FieldGT(FieldGoalsTeam1, v))
}

// GoalsTeam1GTE applies the GTE predicate on the "goalsTeam1" field.
func GoalsTeam1GTE(v int) predicate.Match {
	return predicate.Match(sql.FieldGTE(FieldGoalsTeam1, v))
}

// GoalsTeam1LT applies the LT predicate on the "goalsTeam1" field.
func GoalsTeam1LT(v int) predicate.Match {
	return predicate.Match(sql.FieldLT(FieldGoalsTeam1, v))
}

// GoalsTeam1LTE applies the LTE predicate on the "goalsTeam1" field.
func GoalsTeam1LTE(v int) predicate.Match {
	return predicate.Match(sql.FieldLTE(FieldGoalsTeam1, v))
}

// GoalsTeam2EQ applies the EQ predicate on the "goalsTeam2" field.
func GoalsTeam2EQ(v int) predicate.Match {
	return predicate.Match(sql.FieldEQ(FieldGoalsTeam2, v))
}

// GoalsTeam2NEQ applies the NEQ predicate on the "goalsTeam2" field.
func GoalsTeam2NEQ(v int) predicate.Match {
	return predicate.Match(sql.FieldNEQ(FieldGoalsTeam2, v))
}

// GoalsTeam2In applies the In predicate on the "goalsTeam2" field.
func GoalsTeam2In(vs ...int) predicate.Match {
	return predicate.Match(sql.FieldIn(FieldGoalsTeam2, vs...))
}

// GoalsTeam2NotIn applies the NotIn predicate on the "goalsTeam2" field.
func GoalsTeam2NotIn(vs ...int) predicate.Match {
	return predicate.Match(sql.FieldNotIn(FieldGoalsTeam2, vs...))
}

// GoalsTeam2GT applies the GT predicate on the "goalsTeam2" field.
func GoalsTeam2GT(v int) predicate.Match {
	return predicate.Match(sql.FieldGT(FieldGoalsTeam2, v))
}

// GoalsTeam2GTE applies the GTE predicate on the "goalsTeam2" field.
func GoalsTeam2GTE(v int) predicate.Match {
	return predicate.Match(sql.FieldGTE(FieldGoalsTeam2, v))
}

// GoalsTeam2LT applies the LT predicate on the "goalsTeam2" field.
func GoalsTeam2LT(v int) predicate.Match {
	return predicate.Match(sql.FieldLT(FieldGoalsTeam2, v))
}

// GoalsTeam2LTE applies the LTE predicate on the "goalsTeam2" field.
func GoalsTeam2LTE(v int) predicate.Match {
	return predicate.Match(sql.FieldLTE(FieldGoalsTeam2, v))
}

// HasEvents applies the HasEdge predicate on the "events" edge.
func HasEvents() predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EventsTable, EventsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEventsWith applies the HasEdge predicate on the "events" edge with a given conditions (other predicates).
func HasEventsWith(preds ...predicate.Event) predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		step := newEventsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTeams applies the HasEdge predicate on the "teams" edge.
func HasTeams() predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TeamsTable, TeamsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTeamsWith applies the HasEdge predicate on the "teams" edge with a given conditions (other predicates).
func HasTeamsWith(preds ...predicate.Team) predicate.Match {
	return predicate.Match(func(s *sql.Selector) {
		step := newTeamsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Match) predicate.Match {
	return predicate.Match(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Match) predicate.Match {
	return predicate.Match(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Match) predicate.Match {
	return predicate.Match(sql.NotPredicates(p))
}
