// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/bom-squad/protobom/pkg/sbom/ent/edge"
)

// EdgeCreate is the builder for creating a Edge entity.
type EdgeCreate struct {
	config
	mutation *EdgeMutation
	hooks    []Hook
}

// SetType sets the "type" field.
func (ec *EdgeCreate) SetType(e edge.Type) *EdgeCreate {
	ec.mutation.SetType(e)
	return ec
}

// SetFrom sets the "from" field.
func (ec *EdgeCreate) SetFrom(s string) *EdgeCreate {
	ec.mutation.SetFrom(s)
	return ec
}

// SetTo sets the "to" field.
func (ec *EdgeCreate) SetTo(s string) *EdgeCreate {
	ec.mutation.SetTo(s)
	return ec
}

// Mutation returns the EdgeMutation object of the builder.
func (ec *EdgeCreate) Mutation() *EdgeMutation {
	return ec.mutation
}

// Save creates the Edge in the database.
func (ec *EdgeCreate) Save(ctx context.Context) (*Edge, error) {
	return withHooks(ctx, ec.sqlSave, ec.mutation, ec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EdgeCreate) SaveX(ctx context.Context) *Edge {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EdgeCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EdgeCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EdgeCreate) check() error {
	if _, ok := ec.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Edge.type"`)}
	}
	if v, ok := ec.mutation.GetType(); ok {
		if err := edge.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Edge.type": %w`, err)}
		}
	}
	if _, ok := ec.mutation.From(); !ok {
		return &ValidationError{Name: "from", err: errors.New(`ent: missing required field "Edge.from"`)}
	}
	if _, ok := ec.mutation.To(); !ok {
		return &ValidationError{Name: "to", err: errors.New(`ent: missing required field "Edge.to"`)}
	}
	return nil
}

func (ec *EdgeCreate) sqlSave(ctx context.Context) (*Edge, error) {
	if err := ec.check(); err != nil {
		return nil, err
	}
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ec.mutation.id = &_node.ID
	ec.mutation.done = true
	return _node, nil
}

func (ec *EdgeCreate) createSpec() (*Edge, *sqlgraph.CreateSpec) {
	var (
		_node = &Edge{config: ec.config}
		_spec = sqlgraph.NewCreateSpec(edge.Table, sqlgraph.NewFieldSpec(edge.FieldID, field.TypeInt))
	)
	if value, ok := ec.mutation.GetType(); ok {
		_spec.SetField(edge.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := ec.mutation.From(); ok {
		_spec.SetField(edge.FieldFrom, field.TypeString, value)
		_node.From = value
	}
	if value, ok := ec.mutation.To(); ok {
		_spec.SetField(edge.FieldTo, field.TypeString, value)
		_node.To = value
	}
	return _node, _spec
}

// EdgeCreateBulk is the builder for creating many Edge entities in bulk.
type EdgeCreateBulk struct {
	config
	err      error
	builders []*EdgeCreate
}

// Save creates the Edge entities in the database.
func (ecb *EdgeCreateBulk) Save(ctx context.Context) ([]*Edge, error) {
	if ecb.err != nil {
		return nil, ecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Edge, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EdgeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EdgeCreateBulk) SaveX(ctx context.Context) []*Edge {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EdgeCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EdgeCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}
