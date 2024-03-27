// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/bom-squad/protobom/pkg/sbom/ent/tool"
)

// ToolCreate is the builder for creating a Tool entity.
type ToolCreate struct {
	config
	mutation *ToolMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (tc *ToolCreate) SetName(s string) *ToolCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetVersion sets the "version" field.
func (tc *ToolCreate) SetVersion(s string) *ToolCreate {
	tc.mutation.SetVersion(s)
	return tc
}

// SetVendor sets the "vendor" field.
func (tc *ToolCreate) SetVendor(s string) *ToolCreate {
	tc.mutation.SetVendor(s)
	return tc
}

// Mutation returns the ToolMutation object of the builder.
func (tc *ToolCreate) Mutation() *ToolMutation {
	return tc.mutation
}

// Save creates the Tool in the database.
func (tc *ToolCreate) Save(ctx context.Context) (*Tool, error) {
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *ToolCreate) SaveX(ctx context.Context) *Tool {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *ToolCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *ToolCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *ToolCreate) check() error {
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Tool.name"`)}
	}
	if _, ok := tc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "Tool.version"`)}
	}
	if _, ok := tc.mutation.Vendor(); !ok {
		return &ValidationError{Name: "vendor", err: errors.New(`ent: missing required field "Tool.vendor"`)}
	}
	return nil
}

func (tc *ToolCreate) sqlSave(ctx context.Context) (*Tool, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *ToolCreate) createSpec() (*Tool, *sqlgraph.CreateSpec) {
	var (
		_node = &Tool{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(tool.Table, sqlgraph.NewFieldSpec(tool.FieldID, field.TypeInt))
	)
	if value, ok := tc.mutation.Name(); ok {
		_spec.SetField(tool.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := tc.mutation.Version(); ok {
		_spec.SetField(tool.FieldVersion, field.TypeString, value)
		_node.Version = value
	}
	if value, ok := tc.mutation.Vendor(); ok {
		_spec.SetField(tool.FieldVendor, field.TypeString, value)
		_node.Vendor = value
	}
	return _node, _spec
}

// ToolCreateBulk is the builder for creating many Tool entities in bulk.
type ToolCreateBulk struct {
	config
	err      error
	builders []*ToolCreate
}

// Save creates the Tool entities in the database.
func (tcb *ToolCreateBulk) Save(ctx context.Context) ([]*Tool, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Tool, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ToolMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *ToolCreateBulk) SaveX(ctx context.Context) []*Tool {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *ToolCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *ToolCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
