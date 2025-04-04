// Code generated by ent, DO NOT EDIT.

package ent

import (
	"awesomeProject/ent/predicate"
	"awesomeProject/ent/project"
	"awesomeProject/ent/projectcolumn"
	"awesomeProject/ent/user"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProjectUpdate is the builder for updating Project entities.
type ProjectUpdate struct {
	config
	hooks    []Hook
	mutation *ProjectMutation
}

// Where appends a list predicates to the ProjectUpdate builder.
func (pu *ProjectUpdate) Where(ps ...predicate.Project) *ProjectUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *ProjectUpdate) SetName(s string) *ProjectUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableName(s *string) *ProjectUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// AddColumnIDs adds the "columns" edge to the ProjectColumn entity by IDs.
func (pu *ProjectUpdate) AddColumnIDs(ids ...int) *ProjectUpdate {
	pu.mutation.AddColumnIDs(ids...)
	return pu
}

// AddColumns adds the "columns" edges to the ProjectColumn entity.
func (pu *ProjectUpdate) AddColumns(p ...*ProjectColumn) *ProjectUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddColumnIDs(ids...)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (pu *ProjectUpdate) AddUserIDs(ids ...int) *ProjectUpdate {
	pu.mutation.AddUserIDs(ids...)
	return pu
}

// AddUsers adds the "users" edges to the User entity.
func (pu *ProjectUpdate) AddUsers(u ...*User) *ProjectUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pu.AddUserIDs(ids...)
}

// Mutation returns the ProjectMutation object of the builder.
func (pu *ProjectUpdate) Mutation() *ProjectMutation {
	return pu.mutation
}

// ClearColumns clears all "columns" edges to the ProjectColumn entity.
func (pu *ProjectUpdate) ClearColumns() *ProjectUpdate {
	pu.mutation.ClearColumns()
	return pu
}

// RemoveColumnIDs removes the "columns" edge to ProjectColumn entities by IDs.
func (pu *ProjectUpdate) RemoveColumnIDs(ids ...int) *ProjectUpdate {
	pu.mutation.RemoveColumnIDs(ids...)
	return pu
}

// RemoveColumns removes "columns" edges to ProjectColumn entities.
func (pu *ProjectUpdate) RemoveColumns(p ...*ProjectColumn) *ProjectUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemoveColumnIDs(ids...)
}

// ClearUsers clears all "users" edges to the User entity.
func (pu *ProjectUpdate) ClearUsers() *ProjectUpdate {
	pu.mutation.ClearUsers()
	return pu
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (pu *ProjectUpdate) RemoveUserIDs(ids ...int) *ProjectUpdate {
	pu.mutation.RemoveUserIDs(ids...)
	return pu
}

// RemoveUsers removes "users" edges to User entities.
func (pu *ProjectUpdate) RemoveUsers(u ...*User) *ProjectUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pu.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProjectUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProjectUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProjectUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProjectUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *ProjectUpdate) check() error {
	if v, ok := pu.mutation.Name(); ok {
		if err := project.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Project.name": %w`, err)}
		}
	}
	return nil
}

func (pu *ProjectUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(project.Table, project.Columns, sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(project.FieldName, field.TypeString, value)
	}
	if pu.mutation.ColumnsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.ColumnsTable,
			Columns: []string{project.ColumnsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(projectcolumn.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedColumnsIDs(); len(nodes) > 0 && !pu.mutation.ColumnsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.ColumnsTable,
			Columns: []string{project.ColumnsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(projectcolumn.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ColumnsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.ColumnsTable,
			Columns: []string{project.ColumnsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(projectcolumn.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   project.UsersTable,
			Columns: project.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedUsersIDs(); len(nodes) > 0 && !pu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   project.UsersTable,
			Columns: project.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   project.UsersTable,
			Columns: project.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{project.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProjectUpdateOne is the builder for updating a single Project entity.
type ProjectUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProjectMutation
}

// SetName sets the "name" field.
func (puo *ProjectUpdateOne) SetName(s string) *ProjectUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableName(s *string) *ProjectUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// AddColumnIDs adds the "columns" edge to the ProjectColumn entity by IDs.
func (puo *ProjectUpdateOne) AddColumnIDs(ids ...int) *ProjectUpdateOne {
	puo.mutation.AddColumnIDs(ids...)
	return puo
}

// AddColumns adds the "columns" edges to the ProjectColumn entity.
func (puo *ProjectUpdateOne) AddColumns(p ...*ProjectColumn) *ProjectUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddColumnIDs(ids...)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (puo *ProjectUpdateOne) AddUserIDs(ids ...int) *ProjectUpdateOne {
	puo.mutation.AddUserIDs(ids...)
	return puo
}

// AddUsers adds the "users" edges to the User entity.
func (puo *ProjectUpdateOne) AddUsers(u ...*User) *ProjectUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return puo.AddUserIDs(ids...)
}

// Mutation returns the ProjectMutation object of the builder.
func (puo *ProjectUpdateOne) Mutation() *ProjectMutation {
	return puo.mutation
}

// ClearColumns clears all "columns" edges to the ProjectColumn entity.
func (puo *ProjectUpdateOne) ClearColumns() *ProjectUpdateOne {
	puo.mutation.ClearColumns()
	return puo
}

// RemoveColumnIDs removes the "columns" edge to ProjectColumn entities by IDs.
func (puo *ProjectUpdateOne) RemoveColumnIDs(ids ...int) *ProjectUpdateOne {
	puo.mutation.RemoveColumnIDs(ids...)
	return puo
}

// RemoveColumns removes "columns" edges to ProjectColumn entities.
func (puo *ProjectUpdateOne) RemoveColumns(p ...*ProjectColumn) *ProjectUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemoveColumnIDs(ids...)
}

// ClearUsers clears all "users" edges to the User entity.
func (puo *ProjectUpdateOne) ClearUsers() *ProjectUpdateOne {
	puo.mutation.ClearUsers()
	return puo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (puo *ProjectUpdateOne) RemoveUserIDs(ids ...int) *ProjectUpdateOne {
	puo.mutation.RemoveUserIDs(ids...)
	return puo
}

// RemoveUsers removes "users" edges to User entities.
func (puo *ProjectUpdateOne) RemoveUsers(u ...*User) *ProjectUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return puo.RemoveUserIDs(ids...)
}

// Where appends a list predicates to the ProjectUpdate builder.
func (puo *ProjectUpdateOne) Where(ps ...predicate.Project) *ProjectUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProjectUpdateOne) Select(field string, fields ...string) *ProjectUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Project entity.
func (puo *ProjectUpdateOne) Save(ctx context.Context) (*Project, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProjectUpdateOne) SaveX(ctx context.Context) *Project {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProjectUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProjectUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *ProjectUpdateOne) check() error {
	if v, ok := puo.mutation.Name(); ok {
		if err := project.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Project.name": %w`, err)}
		}
	}
	return nil
}

func (puo *ProjectUpdateOne) sqlSave(ctx context.Context) (_node *Project, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(project.Table, project.Columns, sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Project.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, project.FieldID)
		for _, f := range fields {
			if !project.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != project.FieldID {
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
		_spec.SetField(project.FieldName, field.TypeString, value)
	}
	if puo.mutation.ColumnsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.ColumnsTable,
			Columns: []string{project.ColumnsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(projectcolumn.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedColumnsIDs(); len(nodes) > 0 && !puo.mutation.ColumnsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.ColumnsTable,
			Columns: []string{project.ColumnsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(projectcolumn.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ColumnsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.ColumnsTable,
			Columns: []string{project.ColumnsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(projectcolumn.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   project.UsersTable,
			Columns: project.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !puo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   project.UsersTable,
			Columns: project.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   project.UsersTable,
			Columns: project.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Project{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{project.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
