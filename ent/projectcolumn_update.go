// Code generated by ent, DO NOT EDIT.

package ent

import (
	"awesomeProject/ent/predicate"
	"awesomeProject/ent/project"
	"awesomeProject/ent/projectcolumn"
	"awesomeProject/ent/task"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProjectColumnUpdate is the builder for updating ProjectColumn entities.
type ProjectColumnUpdate struct {
	config
	hooks    []Hook
	mutation *ProjectColumnMutation
}

// Where appends a list predicates to the ProjectColumnUpdate builder.
func (pcu *ProjectColumnUpdate) Where(ps ...predicate.ProjectColumn) *ProjectColumnUpdate {
	pcu.mutation.Where(ps...)
	return pcu
}

// SetName sets the "name" field.
func (pcu *ProjectColumnUpdate) SetName(s string) *ProjectColumnUpdate {
	pcu.mutation.SetName(s)
	return pcu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pcu *ProjectColumnUpdate) SetNillableName(s *string) *ProjectColumnUpdate {
	if s != nil {
		pcu.SetName(*s)
	}
	return pcu
}

// SetOrder sets the "order" field.
func (pcu *ProjectColumnUpdate) SetOrder(i int) *ProjectColumnUpdate {
	pcu.mutation.ResetOrder()
	pcu.mutation.SetOrder(i)
	return pcu
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (pcu *ProjectColumnUpdate) SetNillableOrder(i *int) *ProjectColumnUpdate {
	if i != nil {
		pcu.SetOrder(*i)
	}
	return pcu
}

// AddOrder adds i to the "order" field.
func (pcu *ProjectColumnUpdate) AddOrder(i int) *ProjectColumnUpdate {
	pcu.mutation.AddOrder(i)
	return pcu
}

// SetColor sets the "color" field.
func (pcu *ProjectColumnUpdate) SetColor(s string) *ProjectColumnUpdate {
	pcu.mutation.SetColor(s)
	return pcu
}

// SetNillableColor sets the "color" field if the given value is not nil.
func (pcu *ProjectColumnUpdate) SetNillableColor(s *string) *ProjectColumnUpdate {
	if s != nil {
		pcu.SetColor(*s)
	}
	return pcu
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (pcu *ProjectColumnUpdate) AddTaskIDs(ids ...int) *ProjectColumnUpdate {
	pcu.mutation.AddTaskIDs(ids...)
	return pcu
}

// AddTasks adds the "tasks" edges to the Task entity.
func (pcu *ProjectColumnUpdate) AddTasks(t ...*Task) *ProjectColumnUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pcu.AddTaskIDs(ids...)
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (pcu *ProjectColumnUpdate) SetProjectID(id int) *ProjectColumnUpdate {
	pcu.mutation.SetProjectID(id)
	return pcu
}

// SetNillableProjectID sets the "project" edge to the Project entity by ID if the given value is not nil.
func (pcu *ProjectColumnUpdate) SetNillableProjectID(id *int) *ProjectColumnUpdate {
	if id != nil {
		pcu = pcu.SetProjectID(*id)
	}
	return pcu
}

// SetProject sets the "project" edge to the Project entity.
func (pcu *ProjectColumnUpdate) SetProject(p *Project) *ProjectColumnUpdate {
	return pcu.SetProjectID(p.ID)
}

// Mutation returns the ProjectColumnMutation object of the builder.
func (pcu *ProjectColumnUpdate) Mutation() *ProjectColumnMutation {
	return pcu.mutation
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (pcu *ProjectColumnUpdate) ClearTasks() *ProjectColumnUpdate {
	pcu.mutation.ClearTasks()
	return pcu
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (pcu *ProjectColumnUpdate) RemoveTaskIDs(ids ...int) *ProjectColumnUpdate {
	pcu.mutation.RemoveTaskIDs(ids...)
	return pcu
}

// RemoveTasks removes "tasks" edges to Task entities.
func (pcu *ProjectColumnUpdate) RemoveTasks(t ...*Task) *ProjectColumnUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pcu.RemoveTaskIDs(ids...)
}

// ClearProject clears the "project" edge to the Project entity.
func (pcu *ProjectColumnUpdate) ClearProject() *ProjectColumnUpdate {
	pcu.mutation.ClearProject()
	return pcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pcu *ProjectColumnUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pcu.sqlSave, pcu.mutation, pcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pcu *ProjectColumnUpdate) SaveX(ctx context.Context) int {
	affected, err := pcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pcu *ProjectColumnUpdate) Exec(ctx context.Context) error {
	_, err := pcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcu *ProjectColumnUpdate) ExecX(ctx context.Context) {
	if err := pcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pcu *ProjectColumnUpdate) check() error {
	if v, ok := pcu.mutation.Name(); ok {
		if err := projectcolumn.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "ProjectColumn.name": %w`, err)}
		}
	}
	if v, ok := pcu.mutation.Color(); ok {
		if err := projectcolumn.ColorValidator(v); err != nil {
			return &ValidationError{Name: "color", err: fmt.Errorf(`ent: validator failed for field "ProjectColumn.color": %w`, err)}
		}
	}
	return nil
}

func (pcu *ProjectColumnUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pcu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(projectcolumn.Table, projectcolumn.Columns, sqlgraph.NewFieldSpec(projectcolumn.FieldID, field.TypeInt))
	if ps := pcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pcu.mutation.Name(); ok {
		_spec.SetField(projectcolumn.FieldName, field.TypeString, value)
	}
	if value, ok := pcu.mutation.Order(); ok {
		_spec.SetField(projectcolumn.FieldOrder, field.TypeInt, value)
	}
	if value, ok := pcu.mutation.AddedOrder(); ok {
		_spec.AddField(projectcolumn.FieldOrder, field.TypeInt, value)
	}
	if value, ok := pcu.mutation.Color(); ok {
		_spec.SetField(projectcolumn.FieldColor, field.TypeString, value)
	}
	if pcu.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projectcolumn.TasksTable,
			Columns: []string{projectcolumn.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcu.mutation.RemovedTasksIDs(); len(nodes) > 0 && !pcu.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projectcolumn.TasksTable,
			Columns: []string{projectcolumn.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcu.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projectcolumn.TasksTable,
			Columns: []string{projectcolumn.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pcu.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projectcolumn.ProjectTable,
			Columns: []string{projectcolumn.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcu.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projectcolumn.ProjectTable,
			Columns: []string{projectcolumn.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{projectcolumn.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pcu.mutation.done = true
	return n, nil
}

// ProjectColumnUpdateOne is the builder for updating a single ProjectColumn entity.
type ProjectColumnUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProjectColumnMutation
}

// SetName sets the "name" field.
func (pcuo *ProjectColumnUpdateOne) SetName(s string) *ProjectColumnUpdateOne {
	pcuo.mutation.SetName(s)
	return pcuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pcuo *ProjectColumnUpdateOne) SetNillableName(s *string) *ProjectColumnUpdateOne {
	if s != nil {
		pcuo.SetName(*s)
	}
	return pcuo
}

// SetOrder sets the "order" field.
func (pcuo *ProjectColumnUpdateOne) SetOrder(i int) *ProjectColumnUpdateOne {
	pcuo.mutation.ResetOrder()
	pcuo.mutation.SetOrder(i)
	return pcuo
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (pcuo *ProjectColumnUpdateOne) SetNillableOrder(i *int) *ProjectColumnUpdateOne {
	if i != nil {
		pcuo.SetOrder(*i)
	}
	return pcuo
}

// AddOrder adds i to the "order" field.
func (pcuo *ProjectColumnUpdateOne) AddOrder(i int) *ProjectColumnUpdateOne {
	pcuo.mutation.AddOrder(i)
	return pcuo
}

// SetColor sets the "color" field.
func (pcuo *ProjectColumnUpdateOne) SetColor(s string) *ProjectColumnUpdateOne {
	pcuo.mutation.SetColor(s)
	return pcuo
}

// SetNillableColor sets the "color" field if the given value is not nil.
func (pcuo *ProjectColumnUpdateOne) SetNillableColor(s *string) *ProjectColumnUpdateOne {
	if s != nil {
		pcuo.SetColor(*s)
	}
	return pcuo
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (pcuo *ProjectColumnUpdateOne) AddTaskIDs(ids ...int) *ProjectColumnUpdateOne {
	pcuo.mutation.AddTaskIDs(ids...)
	return pcuo
}

// AddTasks adds the "tasks" edges to the Task entity.
func (pcuo *ProjectColumnUpdateOne) AddTasks(t ...*Task) *ProjectColumnUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pcuo.AddTaskIDs(ids...)
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (pcuo *ProjectColumnUpdateOne) SetProjectID(id int) *ProjectColumnUpdateOne {
	pcuo.mutation.SetProjectID(id)
	return pcuo
}

// SetNillableProjectID sets the "project" edge to the Project entity by ID if the given value is not nil.
func (pcuo *ProjectColumnUpdateOne) SetNillableProjectID(id *int) *ProjectColumnUpdateOne {
	if id != nil {
		pcuo = pcuo.SetProjectID(*id)
	}
	return pcuo
}

// SetProject sets the "project" edge to the Project entity.
func (pcuo *ProjectColumnUpdateOne) SetProject(p *Project) *ProjectColumnUpdateOne {
	return pcuo.SetProjectID(p.ID)
}

// Mutation returns the ProjectColumnMutation object of the builder.
func (pcuo *ProjectColumnUpdateOne) Mutation() *ProjectColumnMutation {
	return pcuo.mutation
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (pcuo *ProjectColumnUpdateOne) ClearTasks() *ProjectColumnUpdateOne {
	pcuo.mutation.ClearTasks()
	return pcuo
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (pcuo *ProjectColumnUpdateOne) RemoveTaskIDs(ids ...int) *ProjectColumnUpdateOne {
	pcuo.mutation.RemoveTaskIDs(ids...)
	return pcuo
}

// RemoveTasks removes "tasks" edges to Task entities.
func (pcuo *ProjectColumnUpdateOne) RemoveTasks(t ...*Task) *ProjectColumnUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pcuo.RemoveTaskIDs(ids...)
}

// ClearProject clears the "project" edge to the Project entity.
func (pcuo *ProjectColumnUpdateOne) ClearProject() *ProjectColumnUpdateOne {
	pcuo.mutation.ClearProject()
	return pcuo
}

// Where appends a list predicates to the ProjectColumnUpdate builder.
func (pcuo *ProjectColumnUpdateOne) Where(ps ...predicate.ProjectColumn) *ProjectColumnUpdateOne {
	pcuo.mutation.Where(ps...)
	return pcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pcuo *ProjectColumnUpdateOne) Select(field string, fields ...string) *ProjectColumnUpdateOne {
	pcuo.fields = append([]string{field}, fields...)
	return pcuo
}

// Save executes the query and returns the updated ProjectColumn entity.
func (pcuo *ProjectColumnUpdateOne) Save(ctx context.Context) (*ProjectColumn, error) {
	return withHooks(ctx, pcuo.sqlSave, pcuo.mutation, pcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pcuo *ProjectColumnUpdateOne) SaveX(ctx context.Context) *ProjectColumn {
	node, err := pcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pcuo *ProjectColumnUpdateOne) Exec(ctx context.Context) error {
	_, err := pcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcuo *ProjectColumnUpdateOne) ExecX(ctx context.Context) {
	if err := pcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pcuo *ProjectColumnUpdateOne) check() error {
	if v, ok := pcuo.mutation.Name(); ok {
		if err := projectcolumn.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "ProjectColumn.name": %w`, err)}
		}
	}
	if v, ok := pcuo.mutation.Color(); ok {
		if err := projectcolumn.ColorValidator(v); err != nil {
			return &ValidationError{Name: "color", err: fmt.Errorf(`ent: validator failed for field "ProjectColumn.color": %w`, err)}
		}
	}
	return nil
}

func (pcuo *ProjectColumnUpdateOne) sqlSave(ctx context.Context) (_node *ProjectColumn, err error) {
	if err := pcuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(projectcolumn.Table, projectcolumn.Columns, sqlgraph.NewFieldSpec(projectcolumn.FieldID, field.TypeInt))
	id, ok := pcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProjectColumn.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, projectcolumn.FieldID)
		for _, f := range fields {
			if !projectcolumn.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != projectcolumn.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pcuo.mutation.Name(); ok {
		_spec.SetField(projectcolumn.FieldName, field.TypeString, value)
	}
	if value, ok := pcuo.mutation.Order(); ok {
		_spec.SetField(projectcolumn.FieldOrder, field.TypeInt, value)
	}
	if value, ok := pcuo.mutation.AddedOrder(); ok {
		_spec.AddField(projectcolumn.FieldOrder, field.TypeInt, value)
	}
	if value, ok := pcuo.mutation.Color(); ok {
		_spec.SetField(projectcolumn.FieldColor, field.TypeString, value)
	}
	if pcuo.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projectcolumn.TasksTable,
			Columns: []string{projectcolumn.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcuo.mutation.RemovedTasksIDs(); len(nodes) > 0 && !pcuo.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projectcolumn.TasksTable,
			Columns: []string{projectcolumn.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcuo.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   projectcolumn.TasksTable,
			Columns: []string{projectcolumn.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pcuo.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projectcolumn.ProjectTable,
			Columns: []string{projectcolumn.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcuo.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   projectcolumn.ProjectTable,
			Columns: []string{projectcolumn.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ProjectColumn{config: pcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{projectcolumn.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pcuo.mutation.done = true
	return _node, nil
}
