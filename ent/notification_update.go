// Code generated by ent, DO NOT EDIT.

package ent

import (
	"awesomeProject/ent/notification"
	"awesomeProject/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NotificationUpdate is the builder for updating Notification entities.
type NotificationUpdate struct {
	config
	hooks    []Hook
	mutation *NotificationMutation
}

// Where appends a list predicates to the NotificationUpdate builder.
func (nu *NotificationUpdate) Where(ps ...predicate.Notification) *NotificationUpdate {
	nu.mutation.Where(ps...)
	return nu
}

// SetTitle sets the "title" field.
func (nu *NotificationUpdate) SetTitle(s string) *NotificationUpdate {
	nu.mutation.SetTitle(s)
	return nu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (nu *NotificationUpdate) SetNillableTitle(s *string) *NotificationUpdate {
	if s != nil {
		nu.SetTitle(*s)
	}
	return nu
}

// SetText sets the "text" field.
func (nu *NotificationUpdate) SetText(s string) *NotificationUpdate {
	nu.mutation.SetText(s)
	return nu
}

// SetNillableText sets the "text" field if the given value is not nil.
func (nu *NotificationUpdate) SetNillableText(s *string) *NotificationUpdate {
	if s != nil {
		nu.SetText(*s)
	}
	return nu
}

// SetType sets the "type" field.
func (nu *NotificationUpdate) SetType(s string) *NotificationUpdate {
	nu.mutation.SetType(s)
	return nu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (nu *NotificationUpdate) SetNillableType(s *string) *NotificationUpdate {
	if s != nil {
		nu.SetType(*s)
	}
	return nu
}

// Mutation returns the NotificationMutation object of the builder.
func (nu *NotificationUpdate) Mutation() *NotificationMutation {
	return nu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NotificationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, nu.sqlSave, nu.mutation, nu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NotificationUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NotificationUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NotificationUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nu *NotificationUpdate) check() error {
	if v, ok := nu.mutation.Title(); ok {
		if err := notification.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Notification.title": %w`, err)}
		}
	}
	if v, ok := nu.mutation.Text(); ok {
		if err := notification.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "Notification.text": %w`, err)}
		}
	}
	if v, ok := nu.mutation.GetType(); ok {
		if err := notification.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Notification.type": %w`, err)}
		}
	}
	return nil
}

func (nu *NotificationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := nu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(notification.Table, notification.Columns, sqlgraph.NewFieldSpec(notification.FieldID, field.TypeInt))
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.Title(); ok {
		_spec.SetField(notification.FieldTitle, field.TypeString, value)
	}
	if value, ok := nu.mutation.Text(); ok {
		_spec.SetField(notification.FieldText, field.TypeString, value)
	}
	if value, ok := nu.mutation.GetType(); ok {
		_spec.SetField(notification.FieldType, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notification.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nu.mutation.done = true
	return n, nil
}

// NotificationUpdateOne is the builder for updating a single Notification entity.
type NotificationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NotificationMutation
}

// SetTitle sets the "title" field.
func (nuo *NotificationUpdateOne) SetTitle(s string) *NotificationUpdateOne {
	nuo.mutation.SetTitle(s)
	return nuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (nuo *NotificationUpdateOne) SetNillableTitle(s *string) *NotificationUpdateOne {
	if s != nil {
		nuo.SetTitle(*s)
	}
	return nuo
}

// SetText sets the "text" field.
func (nuo *NotificationUpdateOne) SetText(s string) *NotificationUpdateOne {
	nuo.mutation.SetText(s)
	return nuo
}

// SetNillableText sets the "text" field if the given value is not nil.
func (nuo *NotificationUpdateOne) SetNillableText(s *string) *NotificationUpdateOne {
	if s != nil {
		nuo.SetText(*s)
	}
	return nuo
}

// SetType sets the "type" field.
func (nuo *NotificationUpdateOne) SetType(s string) *NotificationUpdateOne {
	nuo.mutation.SetType(s)
	return nuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (nuo *NotificationUpdateOne) SetNillableType(s *string) *NotificationUpdateOne {
	if s != nil {
		nuo.SetType(*s)
	}
	return nuo
}

// Mutation returns the NotificationMutation object of the builder.
func (nuo *NotificationUpdateOne) Mutation() *NotificationMutation {
	return nuo.mutation
}

// Where appends a list predicates to the NotificationUpdate builder.
func (nuo *NotificationUpdateOne) Where(ps ...predicate.Notification) *NotificationUpdateOne {
	nuo.mutation.Where(ps...)
	return nuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NotificationUpdateOne) Select(field string, fields ...string) *NotificationUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated Notification entity.
func (nuo *NotificationUpdateOne) Save(ctx context.Context) (*Notification, error) {
	return withHooks(ctx, nuo.sqlSave, nuo.mutation, nuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NotificationUpdateOne) SaveX(ctx context.Context) *Notification {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NotificationUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NotificationUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nuo *NotificationUpdateOne) check() error {
	if v, ok := nuo.mutation.Title(); ok {
		if err := notification.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Notification.title": %w`, err)}
		}
	}
	if v, ok := nuo.mutation.Text(); ok {
		if err := notification.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "Notification.text": %w`, err)}
		}
	}
	if v, ok := nuo.mutation.GetType(); ok {
		if err := notification.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Notification.type": %w`, err)}
		}
	}
	return nil
}

func (nuo *NotificationUpdateOne) sqlSave(ctx context.Context) (_node *Notification, err error) {
	if err := nuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(notification.Table, notification.Columns, sqlgraph.NewFieldSpec(notification.FieldID, field.TypeInt))
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Notification.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notification.FieldID)
		for _, f := range fields {
			if !notification.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != notification.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nuo.mutation.Title(); ok {
		_spec.SetField(notification.FieldTitle, field.TypeString, value)
	}
	if value, ok := nuo.mutation.Text(); ok {
		_spec.SetField(notification.FieldText, field.TypeString, value)
	}
	if value, ok := nuo.mutation.GetType(); ok {
		_spec.SetField(notification.FieldType, field.TypeString, value)
	}
	_node = &Notification{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notification.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nuo.mutation.done = true
	return _node, nil
}
