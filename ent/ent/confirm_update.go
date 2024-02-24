// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"treasure/ent/ent/confirm"
	"treasure/ent/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ConfirmUpdate is the builder for updating Confirm entities.
type ConfirmUpdate struct {
	config
	hooks    []Hook
	mutation *ConfirmMutation
}

// Where appends a list predicates to the ConfirmUpdate builder.
func (cu *ConfirmUpdate) Where(ps ...predicate.Confirm) *ConfirmUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetRequestID sets the "request_id" field.
func (cu *ConfirmUpdate) SetRequestID(i int64) *ConfirmUpdate {
	cu.mutation.ResetRequestID()
	cu.mutation.SetRequestID(i)
	return cu
}

// SetNillableRequestID sets the "request_id" field if the given value is not nil.
func (cu *ConfirmUpdate) SetNillableRequestID(i *int64) *ConfirmUpdate {
	if i != nil {
		cu.SetRequestID(*i)
	}
	return cu
}

// AddRequestID adds i to the "request_id" field.
func (cu *ConfirmUpdate) AddRequestID(i int64) *ConfirmUpdate {
	cu.mutation.AddRequestID(i)
	return cu
}

// SetManagerID sets the "manager_id" field.
func (cu *ConfirmUpdate) SetManagerID(i int) *ConfirmUpdate {
	cu.mutation.ResetManagerID()
	cu.mutation.SetManagerID(i)
	return cu
}

// SetNillableManagerID sets the "manager_id" field if the given value is not nil.
func (cu *ConfirmUpdate) SetNillableManagerID(i *int) *ConfirmUpdate {
	if i != nil {
		cu.SetManagerID(*i)
	}
	return cu
}

// AddManagerID adds i to the "manager_id" field.
func (cu *ConfirmUpdate) AddManagerID(i int) *ConfirmUpdate {
	cu.mutation.AddManagerID(i)
	return cu
}

// Mutation returns the ConfirmMutation object of the builder.
func (cu *ConfirmUpdate) Mutation() *ConfirmMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ConfirmUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ConfirmUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ConfirmUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ConfirmUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *ConfirmUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(confirm.Table, confirm.Columns, sqlgraph.NewFieldSpec(confirm.FieldID, field.TypeInt64))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.RequestID(); ok {
		_spec.SetField(confirm.FieldRequestID, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.AddedRequestID(); ok {
		_spec.AddField(confirm.FieldRequestID, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.ManagerID(); ok {
		_spec.SetField(confirm.FieldManagerID, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedManagerID(); ok {
		_spec.AddField(confirm.FieldManagerID, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{confirm.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ConfirmUpdateOne is the builder for updating a single Confirm entity.
type ConfirmUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ConfirmMutation
}

// SetRequestID sets the "request_id" field.
func (cuo *ConfirmUpdateOne) SetRequestID(i int64) *ConfirmUpdateOne {
	cuo.mutation.ResetRequestID()
	cuo.mutation.SetRequestID(i)
	return cuo
}

// SetNillableRequestID sets the "request_id" field if the given value is not nil.
func (cuo *ConfirmUpdateOne) SetNillableRequestID(i *int64) *ConfirmUpdateOne {
	if i != nil {
		cuo.SetRequestID(*i)
	}
	return cuo
}

// AddRequestID adds i to the "request_id" field.
func (cuo *ConfirmUpdateOne) AddRequestID(i int64) *ConfirmUpdateOne {
	cuo.mutation.AddRequestID(i)
	return cuo
}

// SetManagerID sets the "manager_id" field.
func (cuo *ConfirmUpdateOne) SetManagerID(i int) *ConfirmUpdateOne {
	cuo.mutation.ResetManagerID()
	cuo.mutation.SetManagerID(i)
	return cuo
}

// SetNillableManagerID sets the "manager_id" field if the given value is not nil.
func (cuo *ConfirmUpdateOne) SetNillableManagerID(i *int) *ConfirmUpdateOne {
	if i != nil {
		cuo.SetManagerID(*i)
	}
	return cuo
}

// AddManagerID adds i to the "manager_id" field.
func (cuo *ConfirmUpdateOne) AddManagerID(i int) *ConfirmUpdateOne {
	cuo.mutation.AddManagerID(i)
	return cuo
}

// Mutation returns the ConfirmMutation object of the builder.
func (cuo *ConfirmUpdateOne) Mutation() *ConfirmMutation {
	return cuo.mutation
}

// Where appends a list predicates to the ConfirmUpdate builder.
func (cuo *ConfirmUpdateOne) Where(ps ...predicate.Confirm) *ConfirmUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ConfirmUpdateOne) Select(field string, fields ...string) *ConfirmUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Confirm entity.
func (cuo *ConfirmUpdateOne) Save(ctx context.Context) (*Confirm, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ConfirmUpdateOne) SaveX(ctx context.Context) *Confirm {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ConfirmUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ConfirmUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *ConfirmUpdateOne) sqlSave(ctx context.Context) (_node *Confirm, err error) {
	_spec := sqlgraph.NewUpdateSpec(confirm.Table, confirm.Columns, sqlgraph.NewFieldSpec(confirm.FieldID, field.TypeInt64))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Confirm.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, confirm.FieldID)
		for _, f := range fields {
			if !confirm.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != confirm.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.RequestID(); ok {
		_spec.SetField(confirm.FieldRequestID, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.AddedRequestID(); ok {
		_spec.AddField(confirm.FieldRequestID, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.ManagerID(); ok {
		_spec.SetField(confirm.FieldManagerID, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedManagerID(); ok {
		_spec.AddField(confirm.FieldManagerID, field.TypeInt, value)
	}
	_node = &Confirm{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{confirm.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}