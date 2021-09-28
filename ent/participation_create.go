// Code generated by entc, DO NOT EDIT.

package ent

import (
	"calcio/ent/game"
	"calcio/ent/participation"
	"calcio/ent/team"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ParticipationCreate is the builder for creating a Participation entity.
type ParticipationCreate struct {
	config
	mutation *ParticipationMutation
	hooks    []Hook
}

// SetGoals sets the "goals" field.
func (pc *ParticipationCreate) SetGoals(i int) *ParticipationCreate {
	pc.mutation.SetGoals(i)
	return pc
}

// SetNillableGoals sets the "goals" field if the given value is not nil.
func (pc *ParticipationCreate) SetNillableGoals(i *int) *ParticipationCreate {
	if i != nil {
		pc.SetGoals(*i)
	}
	return pc
}

// SetGameID sets the "game" edge to the Game entity by ID.
func (pc *ParticipationCreate) SetGameID(id uuid.UUID) *ParticipationCreate {
	pc.mutation.SetGameID(id)
	return pc
}

// SetGame sets the "game" edge to the Game entity.
func (pc *ParticipationCreate) SetGame(g *Game) *ParticipationCreate {
	return pc.SetGameID(g.ID)
}

// SetTeamID sets the "team" edge to the Team entity by ID.
func (pc *ParticipationCreate) SetTeamID(id uuid.UUID) *ParticipationCreate {
	pc.mutation.SetTeamID(id)
	return pc
}

// SetTeam sets the "team" edge to the Team entity.
func (pc *ParticipationCreate) SetTeam(t *Team) *ParticipationCreate {
	return pc.SetTeamID(t.ID)
}

// Mutation returns the ParticipationMutation object of the builder.
func (pc *ParticipationCreate) Mutation() *ParticipationMutation {
	return pc.mutation
}

// Save creates the Participation in the database.
func (pc *ParticipationCreate) Save(ctx context.Context) (*Participation, error) {
	var (
		err  error
		node *Participation
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ParticipationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ParticipationCreate) SaveX(ctx context.Context) *Participation {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ParticipationCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ParticipationCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ParticipationCreate) defaults() {
	if _, ok := pc.mutation.Goals(); !ok {
		v := participation.DefaultGoals
		pc.mutation.SetGoals(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ParticipationCreate) check() error {
	if _, ok := pc.mutation.Goals(); !ok {
		return &ValidationError{Name: "goals", err: errors.New(`ent: missing required field "goals"`)}
	}
	if v, ok := pc.mutation.Goals(); ok {
		if err := participation.GoalsValidator(v); err != nil {
			return &ValidationError{Name: "goals", err: fmt.Errorf(`ent: validator failed for field "goals": %w`, err)}
		}
	}
	if _, ok := pc.mutation.GameID(); !ok {
		return &ValidationError{Name: "game", err: errors.New("ent: missing required edge \"game\"")}
	}
	if _, ok := pc.mutation.TeamID(); !ok {
		return &ValidationError{Name: "team", err: errors.New("ent: missing required edge \"team\"")}
	}
	return nil
}

func (pc *ParticipationCreate) sqlSave(ctx context.Context) (*Participation, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pc *ParticipationCreate) createSpec() (*Participation, *sqlgraph.CreateSpec) {
	var (
		_node = &Participation{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: participation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: participation.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.Goals(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: participation.FieldGoals,
		})
		_node.Goals = value
	}
	if nodes := pc.mutation.GameIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   participation.GameTable,
			Columns: []string{participation.GameColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: game.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.participation_game = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.TeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   participation.TeamTable,
			Columns: []string{participation.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: team.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.participation_team = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ParticipationCreateBulk is the builder for creating many Participation entities in bulk.
type ParticipationCreateBulk struct {
	config
	builders []*ParticipationCreate
}

// Save creates the Participation entities in the database.
func (pcb *ParticipationCreateBulk) Save(ctx context.Context) ([]*Participation, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Participation, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ParticipationMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ParticipationCreateBulk) SaveX(ctx context.Context) []*Participation {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ParticipationCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ParticipationCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}