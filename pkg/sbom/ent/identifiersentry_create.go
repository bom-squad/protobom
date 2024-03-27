// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/bom-squad/protobom/pkg/sbom/ent/identifiersentry"
)

// IdentifiersEntryCreate is the builder for creating a IdentifiersEntry entity.
type IdentifiersEntryCreate struct {
	config
	mutation *IdentifiersEntryMutation
	hooks    []Hook
}

// SetSoftwareIdentifierType sets the "software_identifier_type" field.
func (iec *IdentifiersEntryCreate) SetSoftwareIdentifierType(iit identifiersentry.SoftwareIdentifierType) *IdentifiersEntryCreate {
	iec.mutation.SetSoftwareIdentifierType(iit)
	return iec
}

// SetSoftwareIdentifierValue sets the "software_identifier_value" field.
func (iec *IdentifiersEntryCreate) SetSoftwareIdentifierValue(s string) *IdentifiersEntryCreate {
	iec.mutation.SetSoftwareIdentifierValue(s)
	return iec
}

// Mutation returns the IdentifiersEntryMutation object of the builder.
func (iec *IdentifiersEntryCreate) Mutation() *IdentifiersEntryMutation {
	return iec.mutation
}

// Save creates the IdentifiersEntry in the database.
func (iec *IdentifiersEntryCreate) Save(ctx context.Context) (*IdentifiersEntry, error) {
	return withHooks(ctx, iec.sqlSave, iec.mutation, iec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (iec *IdentifiersEntryCreate) SaveX(ctx context.Context) *IdentifiersEntry {
	v, err := iec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (iec *IdentifiersEntryCreate) Exec(ctx context.Context) error {
	_, err := iec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iec *IdentifiersEntryCreate) ExecX(ctx context.Context) {
	if err := iec.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iec *IdentifiersEntryCreate) check() error {
	if _, ok := iec.mutation.SoftwareIdentifierType(); !ok {
		return &ValidationError{Name: "software_identifier_type", err: errors.New(`ent: missing required field "IdentifiersEntry.software_identifier_type"`)}
	}
	if v, ok := iec.mutation.SoftwareIdentifierType(); ok {
		if err := identifiersentry.SoftwareIdentifierTypeValidator(v); err != nil {
			return &ValidationError{Name: "software_identifier_type", err: fmt.Errorf(`ent: validator failed for field "IdentifiersEntry.software_identifier_type": %w`, err)}
		}
	}
	if _, ok := iec.mutation.SoftwareIdentifierValue(); !ok {
		return &ValidationError{Name: "software_identifier_value", err: errors.New(`ent: missing required field "IdentifiersEntry.software_identifier_value"`)}
	}
	return nil
}

func (iec *IdentifiersEntryCreate) sqlSave(ctx context.Context) (*IdentifiersEntry, error) {
	if err := iec.check(); err != nil {
		return nil, err
	}
	_node, _spec := iec.createSpec()
	if err := sqlgraph.CreateNode(ctx, iec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	iec.mutation.id = &_node.ID
	iec.mutation.done = true
	return _node, nil
}

func (iec *IdentifiersEntryCreate) createSpec() (*IdentifiersEntry, *sqlgraph.CreateSpec) {
	var (
		_node = &IdentifiersEntry{config: iec.config}
		_spec = sqlgraph.NewCreateSpec(identifiersentry.Table, sqlgraph.NewFieldSpec(identifiersentry.FieldID, field.TypeInt))
	)
	if value, ok := iec.mutation.SoftwareIdentifierType(); ok {
		_spec.SetField(identifiersentry.FieldSoftwareIdentifierType, field.TypeEnum, value)
		_node.SoftwareIdentifierType = value
	}
	if value, ok := iec.mutation.SoftwareIdentifierValue(); ok {
		_spec.SetField(identifiersentry.FieldSoftwareIdentifierValue, field.TypeString, value)
		_node.SoftwareIdentifierValue = value
	}
	return _node, _spec
}

// IdentifiersEntryCreateBulk is the builder for creating many IdentifiersEntry entities in bulk.
type IdentifiersEntryCreateBulk struct {
	config
	err      error
	builders []*IdentifiersEntryCreate
}

// Save creates the IdentifiersEntry entities in the database.
func (iecb *IdentifiersEntryCreateBulk) Save(ctx context.Context) ([]*IdentifiersEntry, error) {
	if iecb.err != nil {
		return nil, iecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(iecb.builders))
	nodes := make([]*IdentifiersEntry, len(iecb.builders))
	mutators := make([]Mutator, len(iecb.builders))
	for i := range iecb.builders {
		func(i int, root context.Context) {
			builder := iecb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*IdentifiersEntryMutation)
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
					_, err = mutators[i+1].Mutate(root, iecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, iecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, iecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (iecb *IdentifiersEntryCreateBulk) SaveX(ctx context.Context) []*IdentifiersEntry {
	v, err := iecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (iecb *IdentifiersEntryCreateBulk) Exec(ctx context.Context) error {
	_, err := iecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iecb *IdentifiersEntryCreateBulk) ExecX(ctx context.Context) {
	if err := iecb.Exec(ctx); err != nil {
		panic(err)
	}
}
