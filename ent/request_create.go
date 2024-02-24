// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/longxboy/treasure/ent/request"
)

// RequestCreate is the builder for creating a Request entity.
type RequestCreate struct {
	config
	mutation *RequestMutation
	hooks    []Hook
}

// SetStatus sets the "status" field.
func (rc *RequestCreate) SetStatus(s string) *RequestCreate {
	rc.mutation.SetStatus(s)
	return rc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (rc *RequestCreate) SetNillableStatus(s *string) *RequestCreate {
	if s != nil {
		rc.SetStatus(*s)
	}
	return rc
}

// SetAmount sets the "amount" field.
func (rc *RequestCreate) SetAmount(i int64) *RequestCreate {
	rc.mutation.SetAmount(i)
	return rc
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (rc *RequestCreate) SetNillableAmount(i *int64) *RequestCreate {
	if i != nil {
		rc.SetAmount(*i)
	}
	return rc
}

// SetRecipient sets the "recipient" field.
func (rc *RequestCreate) SetRecipient(s string) *RequestCreate {
	rc.mutation.SetRecipient(s)
	return rc
}

// SetTxHash sets the "tx_hash" field.
func (rc *RequestCreate) SetTxHash(s string) *RequestCreate {
	rc.mutation.SetTxHash(s)
	return rc
}

// SetNillableTxHash sets the "tx_hash" field if the given value is not nil.
func (rc *RequestCreate) SetNillableTxHash(s *string) *RequestCreate {
	if s != nil {
		rc.SetTxHash(*s)
	}
	return rc
}

// SetNonce sets the "nonce" field.
func (rc *RequestCreate) SetNonce(i int64) *RequestCreate {
	rc.mutation.SetNonce(i)
	return rc
}

// SetNillableNonce sets the "nonce" field if the given value is not nil.
func (rc *RequestCreate) SetNillableNonce(i *int64) *RequestCreate {
	if i != nil {
		rc.SetNonce(*i)
	}
	return rc
}

// SetExecuted sets the "executed" field.
func (rc *RequestCreate) SetExecuted(b bool) *RequestCreate {
	rc.mutation.SetExecuted(b)
	return rc
}

// SetNillableExecuted sets the "executed" field if the given value is not nil.
func (rc *RequestCreate) SetNillableExecuted(b *bool) *RequestCreate {
	if b != nil {
		rc.SetExecuted(*b)
	}
	return rc
}

// SetID sets the "id" field.
func (rc *RequestCreate) SetID(i int64) *RequestCreate {
	rc.mutation.SetID(i)
	return rc
}

// Mutation returns the RequestMutation object of the builder.
func (rc *RequestCreate) Mutation() *RequestMutation {
	return rc.mutation
}

// Save creates the Request in the database.
func (rc *RequestCreate) Save(ctx context.Context) (*Request, error) {
	rc.defaults()
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RequestCreate) SaveX(ctx context.Context) *Request {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RequestCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RequestCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RequestCreate) defaults() {
	if _, ok := rc.mutation.Status(); !ok {
		v := request.DefaultStatus
		rc.mutation.SetStatus(v)
	}
	if _, ok := rc.mutation.Amount(); !ok {
		v := request.DefaultAmount
		rc.mutation.SetAmount(v)
	}
	if _, ok := rc.mutation.TxHash(); !ok {
		v := request.DefaultTxHash
		rc.mutation.SetTxHash(v)
	}
	if _, ok := rc.mutation.Nonce(); !ok {
		v := request.DefaultNonce
		rc.mutation.SetNonce(v)
	}
	if _, ok := rc.mutation.Executed(); !ok {
		v := request.DefaultExecuted
		rc.mutation.SetExecuted(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RequestCreate) check() error {
	if _, ok := rc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Request.status"`)}
	}
	if _, ok := rc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "Request.amount"`)}
	}
	if _, ok := rc.mutation.Recipient(); !ok {
		return &ValidationError{Name: "recipient", err: errors.New(`ent: missing required field "Request.recipient"`)}
	}
	if _, ok := rc.mutation.TxHash(); !ok {
		return &ValidationError{Name: "tx_hash", err: errors.New(`ent: missing required field "Request.tx_hash"`)}
	}
	if _, ok := rc.mutation.Nonce(); !ok {
		return &ValidationError{Name: "nonce", err: errors.New(`ent: missing required field "Request.nonce"`)}
	}
	if _, ok := rc.mutation.Executed(); !ok {
		return &ValidationError{Name: "executed", err: errors.New(`ent: missing required field "Request.executed"`)}
	}
	return nil
}

func (rc *RequestCreate) sqlSave(ctx context.Context) (*Request, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *RequestCreate) createSpec() (*Request, *sqlgraph.CreateSpec) {
	var (
		_node = &Request{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(request.Table, sqlgraph.NewFieldSpec(request.FieldID, field.TypeInt64))
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rc.mutation.Status(); ok {
		_spec.SetField(request.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := rc.mutation.Amount(); ok {
		_spec.SetField(request.FieldAmount, field.TypeInt64, value)
		_node.Amount = value
	}
	if value, ok := rc.mutation.Recipient(); ok {
		_spec.SetField(request.FieldRecipient, field.TypeString, value)
		_node.Recipient = value
	}
	if value, ok := rc.mutation.TxHash(); ok {
		_spec.SetField(request.FieldTxHash, field.TypeString, value)
		_node.TxHash = value
	}
	if value, ok := rc.mutation.Nonce(); ok {
		_spec.SetField(request.FieldNonce, field.TypeInt64, value)
		_node.Nonce = value
	}
	if value, ok := rc.mutation.Executed(); ok {
		_spec.SetField(request.FieldExecuted, field.TypeBool, value)
		_node.Executed = value
	}
	return _node, _spec
}

// RequestCreateBulk is the builder for creating many Request entities in bulk.
type RequestCreateBulk struct {
	config
	err      error
	builders []*RequestCreate
}

// Save creates the Request entities in the database.
func (rcb *RequestCreateBulk) Save(ctx context.Context) ([]*Request, error) {
	if rcb.err != nil {
		return nil, rcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Request, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RequestMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RequestCreateBulk) SaveX(ctx context.Context) []*Request {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RequestCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RequestCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
