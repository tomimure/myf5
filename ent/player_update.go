// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"myf5/ent/event"
	"myf5/ent/player"
	"myf5/ent/predicate"
	"myf5/ent/team"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PlayerUpdate is the builder for updating Player entities.
type PlayerUpdate struct {
	config
	hooks    []Hook
	mutation *PlayerMutation
}

// Where appends a list predicates to the PlayerUpdate builder.
func (pu *PlayerUpdate) Where(ps ...predicate.Player) *PlayerUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetGlobal sets the "global" field.
func (pu *PlayerUpdate) SetGlobal(i int) *PlayerUpdate {
	pu.mutation.ResetGlobal()
	pu.mutation.SetGlobal(i)
	return pu
}

// SetNillableGlobal sets the "global" field if the given value is not nil.
func (pu *PlayerUpdate) SetNillableGlobal(i *int) *PlayerUpdate {
	if i != nil {
		pu.SetGlobal(*i)
	}
	return pu
}

// AddGlobal adds i to the "global" field.
func (pu *PlayerUpdate) AddGlobal(i int) *PlayerUpdate {
	pu.mutation.AddGlobal(i)
	return pu
}

// SetName sets the "name" field.
func (pu *PlayerUpdate) SetName(s string) *PlayerUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *PlayerUpdate) SetNillableName(s *string) *PlayerUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// AddTeamIDs adds the "teams" edge to the Team entity by IDs.
func (pu *PlayerUpdate) AddTeamIDs(ids ...int) *PlayerUpdate {
	pu.mutation.AddTeamIDs(ids...)
	return pu
}

// AddTeams adds the "teams" edges to the Team entity.
func (pu *PlayerUpdate) AddTeams(t ...*Team) *PlayerUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pu.AddTeamIDs(ids...)
}

// AddEventIDs adds the "events" edge to the Event entity by IDs.
func (pu *PlayerUpdate) AddEventIDs(ids ...int) *PlayerUpdate {
	pu.mutation.AddEventIDs(ids...)
	return pu
}

// AddEvents adds the "events" edges to the Event entity.
func (pu *PlayerUpdate) AddEvents(e ...*Event) *PlayerUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return pu.AddEventIDs(ids...)
}

// Mutation returns the PlayerMutation object of the builder.
func (pu *PlayerUpdate) Mutation() *PlayerMutation {
	return pu.mutation
}

// ClearTeams clears all "teams" edges to the Team entity.
func (pu *PlayerUpdate) ClearTeams() *PlayerUpdate {
	pu.mutation.ClearTeams()
	return pu
}

// RemoveTeamIDs removes the "teams" edge to Team entities by IDs.
func (pu *PlayerUpdate) RemoveTeamIDs(ids ...int) *PlayerUpdate {
	pu.mutation.RemoveTeamIDs(ids...)
	return pu
}

// RemoveTeams removes "teams" edges to Team entities.
func (pu *PlayerUpdate) RemoveTeams(t ...*Team) *PlayerUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pu.RemoveTeamIDs(ids...)
}

// ClearEvents clears all "events" edges to the Event entity.
func (pu *PlayerUpdate) ClearEvents() *PlayerUpdate {
	pu.mutation.ClearEvents()
	return pu
}

// RemoveEventIDs removes the "events" edge to Event entities by IDs.
func (pu *PlayerUpdate) RemoveEventIDs(ids ...int) *PlayerUpdate {
	pu.mutation.RemoveEventIDs(ids...)
	return pu
}

// RemoveEvents removes "events" edges to Event entities.
func (pu *PlayerUpdate) RemoveEvents(e ...*Event) *PlayerUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return pu.RemoveEventIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PlayerUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PlayerUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PlayerUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PlayerUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PlayerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(player.Table, player.Columns, sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Global(); ok {
		_spec.SetField(player.FieldGlobal, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedGlobal(); ok {
		_spec.AddField(player.FieldGlobal, field.TypeInt, value)
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(player.FieldName, field.TypeString, value)
	}
	if pu.mutation.TeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   player.TeamsTable,
			Columns: player.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedTeamsIDs(); len(nodes) > 0 && !pu.mutation.TeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   player.TeamsTable,
			Columns: player.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.TeamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   player.TeamsTable,
			Columns: player.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.EventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   player.EventsTable,
			Columns: []string{player.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedEventsIDs(); len(nodes) > 0 && !pu.mutation.EventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   player.EventsTable,
			Columns: []string{player.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   player.EventsTable,
			Columns: []string{player.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{player.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PlayerUpdateOne is the builder for updating a single Player entity.
type PlayerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PlayerMutation
}

// SetGlobal sets the "global" field.
func (puo *PlayerUpdateOne) SetGlobal(i int) *PlayerUpdateOne {
	puo.mutation.ResetGlobal()
	puo.mutation.SetGlobal(i)
	return puo
}

// SetNillableGlobal sets the "global" field if the given value is not nil.
func (puo *PlayerUpdateOne) SetNillableGlobal(i *int) *PlayerUpdateOne {
	if i != nil {
		puo.SetGlobal(*i)
	}
	return puo
}

// AddGlobal adds i to the "global" field.
func (puo *PlayerUpdateOne) AddGlobal(i int) *PlayerUpdateOne {
	puo.mutation.AddGlobal(i)
	return puo
}

// SetName sets the "name" field.
func (puo *PlayerUpdateOne) SetName(s string) *PlayerUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *PlayerUpdateOne) SetNillableName(s *string) *PlayerUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// AddTeamIDs adds the "teams" edge to the Team entity by IDs.
func (puo *PlayerUpdateOne) AddTeamIDs(ids ...int) *PlayerUpdateOne {
	puo.mutation.AddTeamIDs(ids...)
	return puo
}

// AddTeams adds the "teams" edges to the Team entity.
func (puo *PlayerUpdateOne) AddTeams(t ...*Team) *PlayerUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return puo.AddTeamIDs(ids...)
}

// AddEventIDs adds the "events" edge to the Event entity by IDs.
func (puo *PlayerUpdateOne) AddEventIDs(ids ...int) *PlayerUpdateOne {
	puo.mutation.AddEventIDs(ids...)
	return puo
}

// AddEvents adds the "events" edges to the Event entity.
func (puo *PlayerUpdateOne) AddEvents(e ...*Event) *PlayerUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return puo.AddEventIDs(ids...)
}

// Mutation returns the PlayerMutation object of the builder.
func (puo *PlayerUpdateOne) Mutation() *PlayerMutation {
	return puo.mutation
}

// ClearTeams clears all "teams" edges to the Team entity.
func (puo *PlayerUpdateOne) ClearTeams() *PlayerUpdateOne {
	puo.mutation.ClearTeams()
	return puo
}

// RemoveTeamIDs removes the "teams" edge to Team entities by IDs.
func (puo *PlayerUpdateOne) RemoveTeamIDs(ids ...int) *PlayerUpdateOne {
	puo.mutation.RemoveTeamIDs(ids...)
	return puo
}

// RemoveTeams removes "teams" edges to Team entities.
func (puo *PlayerUpdateOne) RemoveTeams(t ...*Team) *PlayerUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return puo.RemoveTeamIDs(ids...)
}

// ClearEvents clears all "events" edges to the Event entity.
func (puo *PlayerUpdateOne) ClearEvents() *PlayerUpdateOne {
	puo.mutation.ClearEvents()
	return puo
}

// RemoveEventIDs removes the "events" edge to Event entities by IDs.
func (puo *PlayerUpdateOne) RemoveEventIDs(ids ...int) *PlayerUpdateOne {
	puo.mutation.RemoveEventIDs(ids...)
	return puo
}

// RemoveEvents removes "events" edges to Event entities.
func (puo *PlayerUpdateOne) RemoveEvents(e ...*Event) *PlayerUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return puo.RemoveEventIDs(ids...)
}

// Where appends a list predicates to the PlayerUpdate builder.
func (puo *PlayerUpdateOne) Where(ps ...predicate.Player) *PlayerUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PlayerUpdateOne) Select(field string, fields ...string) *PlayerUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Player entity.
func (puo *PlayerUpdateOne) Save(ctx context.Context) (*Player, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PlayerUpdateOne) SaveX(ctx context.Context) *Player {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PlayerUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PlayerUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PlayerUpdateOne) sqlSave(ctx context.Context) (_node *Player, err error) {
	_spec := sqlgraph.NewUpdateSpec(player.Table, player.Columns, sqlgraph.NewFieldSpec(player.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Player.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, player.FieldID)
		for _, f := range fields {
			if !player.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != player.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Global(); ok {
		_spec.SetField(player.FieldGlobal, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedGlobal(); ok {
		_spec.AddField(player.FieldGlobal, field.TypeInt, value)
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(player.FieldName, field.TypeString, value)
	}
	if puo.mutation.TeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   player.TeamsTable,
			Columns: player.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedTeamsIDs(); len(nodes) > 0 && !puo.mutation.TeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   player.TeamsTable,
			Columns: player.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.TeamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   player.TeamsTable,
			Columns: player.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.EventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   player.EventsTable,
			Columns: []string{player.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedEventsIDs(); len(nodes) > 0 && !puo.mutation.EventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   player.EventsTable,
			Columns: []string{player.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   player.EventsTable,
			Columns: []string{player.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Player{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{player.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
