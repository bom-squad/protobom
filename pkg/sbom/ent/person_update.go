// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/bom-squad/protobom/pkg/sbom/ent/person"
	"github.com/bom-squad/protobom/pkg/sbom/ent/predicate"
)

// PersonUpdate is the builder for updating Person entities.
type PersonUpdate struct {
	config
	hooks    []Hook
	mutation *PersonMutation
}

// Where appends a list predicates to the PersonUpdate builder.
func (pu *PersonUpdate) Where(ps ...predicate.Person) *PersonUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *PersonUpdate) SetName(s string) *PersonUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *PersonUpdate) SetNillableName(s *string) *PersonUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetIsOrg sets the "is_org" field.
func (pu *PersonUpdate) SetIsOrg(b bool) *PersonUpdate {
	pu.mutation.SetIsOrg(b)
	return pu
}

// SetNillableIsOrg sets the "is_org" field if the given value is not nil.
func (pu *PersonUpdate) SetNillableIsOrg(b *bool) *PersonUpdate {
	if b != nil {
		pu.SetIsOrg(*b)
	}
	return pu
}

// SetEmail sets the "email" field.
func (pu *PersonUpdate) SetEmail(s string) *PersonUpdate {
	pu.mutation.SetEmail(s)
	return pu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (pu *PersonUpdate) SetNillableEmail(s *string) *PersonUpdate {
	if s != nil {
		pu.SetEmail(*s)
	}
	return pu
}

// SetURL sets the "url" field.
func (pu *PersonUpdate) SetURL(s string) *PersonUpdate {
	pu.mutation.SetURL(s)
	return pu
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (pu *PersonUpdate) SetNillableURL(s *string) *PersonUpdate {
	if s != nil {
		pu.SetURL(*s)
	}
	return pu
}

// SetPhone sets the "phone" field.
func (pu *PersonUpdate) SetPhone(s string) *PersonUpdate {
	pu.mutation.SetPhone(s)
	return pu
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (pu *PersonUpdate) SetNillablePhone(s *string) *PersonUpdate {
	if s != nil {
		pu.SetPhone(*s)
	}
	return pu
}

// AddContactIDs adds the "contacts" edge to the Person entity by IDs.
func (pu *PersonUpdate) AddContactIDs(ids ...int) *PersonUpdate {
	pu.mutation.AddContactIDs(ids...)
	return pu
}

// AddContacts adds the "contacts" edges to the Person entity.
func (pu *PersonUpdate) AddContacts(p ...*Person) *PersonUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddContactIDs(ids...)
}

// Mutation returns the PersonMutation object of the builder.
func (pu *PersonUpdate) Mutation() *PersonMutation {
	return pu.mutation
}

// ClearContacts clears all "contacts" edges to the Person entity.
func (pu *PersonUpdate) ClearContacts() *PersonUpdate {
	pu.mutation.ClearContacts()
	return pu
}

// RemoveContactIDs removes the "contacts" edge to Person entities by IDs.
func (pu *PersonUpdate) RemoveContactIDs(ids ...int) *PersonUpdate {
	pu.mutation.RemoveContactIDs(ids...)
	return pu
}

// RemoveContacts removes "contacts" edges to Person entities.
func (pu *PersonUpdate) RemoveContacts(p ...*Person) *PersonUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemoveContactIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PersonUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PersonUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PersonUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PersonUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PersonUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(person.Table, person.Columns, sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(person.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.IsOrg(); ok {
		_spec.SetField(person.FieldIsOrg, field.TypeBool, value)
	}
	if value, ok := pu.mutation.Email(); ok {
		_spec.SetField(person.FieldEmail, field.TypeString, value)
	}
	if value, ok := pu.mutation.URL(); ok {
		_spec.SetField(person.FieldURL, field.TypeString, value)
	}
	if value, ok := pu.mutation.Phone(); ok {
		_spec.SetField(person.FieldPhone, field.TypeString, value)
	}
	if pu.mutation.ContactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   person.ContactsTable,
			Columns: person.ContactsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedContactsIDs(); len(nodes) > 0 && !pu.mutation.ContactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   person.ContactsTable,
			Columns: person.ContactsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ContactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   person.ContactsTable,
			Columns: person.ContactsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{person.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PersonUpdateOne is the builder for updating a single Person entity.
type PersonUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PersonMutation
}

// SetName sets the "name" field.
func (puo *PersonUpdateOne) SetName(s string) *PersonUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *PersonUpdateOne) SetNillableName(s *string) *PersonUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetIsOrg sets the "is_org" field.
func (puo *PersonUpdateOne) SetIsOrg(b bool) *PersonUpdateOne {
	puo.mutation.SetIsOrg(b)
	return puo
}

// SetNillableIsOrg sets the "is_org" field if the given value is not nil.
func (puo *PersonUpdateOne) SetNillableIsOrg(b *bool) *PersonUpdateOne {
	if b != nil {
		puo.SetIsOrg(*b)
	}
	return puo
}

// SetEmail sets the "email" field.
func (puo *PersonUpdateOne) SetEmail(s string) *PersonUpdateOne {
	puo.mutation.SetEmail(s)
	return puo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (puo *PersonUpdateOne) SetNillableEmail(s *string) *PersonUpdateOne {
	if s != nil {
		puo.SetEmail(*s)
	}
	return puo
}

// SetURL sets the "url" field.
func (puo *PersonUpdateOne) SetURL(s string) *PersonUpdateOne {
	puo.mutation.SetURL(s)
	return puo
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (puo *PersonUpdateOne) SetNillableURL(s *string) *PersonUpdateOne {
	if s != nil {
		puo.SetURL(*s)
	}
	return puo
}

// SetPhone sets the "phone" field.
func (puo *PersonUpdateOne) SetPhone(s string) *PersonUpdateOne {
	puo.mutation.SetPhone(s)
	return puo
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (puo *PersonUpdateOne) SetNillablePhone(s *string) *PersonUpdateOne {
	if s != nil {
		puo.SetPhone(*s)
	}
	return puo
}

// AddContactIDs adds the "contacts" edge to the Person entity by IDs.
func (puo *PersonUpdateOne) AddContactIDs(ids ...int) *PersonUpdateOne {
	puo.mutation.AddContactIDs(ids...)
	return puo
}

// AddContacts adds the "contacts" edges to the Person entity.
func (puo *PersonUpdateOne) AddContacts(p ...*Person) *PersonUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddContactIDs(ids...)
}

// Mutation returns the PersonMutation object of the builder.
func (puo *PersonUpdateOne) Mutation() *PersonMutation {
	return puo.mutation
}

// ClearContacts clears all "contacts" edges to the Person entity.
func (puo *PersonUpdateOne) ClearContacts() *PersonUpdateOne {
	puo.mutation.ClearContacts()
	return puo
}

// RemoveContactIDs removes the "contacts" edge to Person entities by IDs.
func (puo *PersonUpdateOne) RemoveContactIDs(ids ...int) *PersonUpdateOne {
	puo.mutation.RemoveContactIDs(ids...)
	return puo
}

// RemoveContacts removes "contacts" edges to Person entities.
func (puo *PersonUpdateOne) RemoveContacts(p ...*Person) *PersonUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemoveContactIDs(ids...)
}

// Where appends a list predicates to the PersonUpdate builder.
func (puo *PersonUpdateOne) Where(ps ...predicate.Person) *PersonUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PersonUpdateOne) Select(field string, fields ...string) *PersonUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Person entity.
func (puo *PersonUpdateOne) Save(ctx context.Context) (*Person, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PersonUpdateOne) SaveX(ctx context.Context) *Person {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PersonUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PersonUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PersonUpdateOne) sqlSave(ctx context.Context) (_node *Person, err error) {
	_spec := sqlgraph.NewUpdateSpec(person.Table, person.Columns, sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Person.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, person.FieldID)
		for _, f := range fields {
			if !person.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != person.FieldID {
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
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(person.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.IsOrg(); ok {
		_spec.SetField(person.FieldIsOrg, field.TypeBool, value)
	}
	if value, ok := puo.mutation.Email(); ok {
		_spec.SetField(person.FieldEmail, field.TypeString, value)
	}
	if value, ok := puo.mutation.URL(); ok {
		_spec.SetField(person.FieldURL, field.TypeString, value)
	}
	if value, ok := puo.mutation.Phone(); ok {
		_spec.SetField(person.FieldPhone, field.TypeString, value)
	}
	if puo.mutation.ContactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   person.ContactsTable,
			Columns: person.ContactsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedContactsIDs(); len(nodes) > 0 && !puo.mutation.ContactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   person.ContactsTable,
			Columns: person.ContactsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ContactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   person.ContactsTable,
			Columns: person.ContactsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Person{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{person.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
