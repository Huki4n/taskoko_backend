// Code generated by ent, DO NOT EDIT.

package ent

import (
	"awesomeProject/ent/subtask"
	"awesomeProject/ent/task"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Subtask is the model entity for the Subtask schema.
type Subtask struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SubtaskQuery when eager-loading is set.
	Edges         SubtaskEdges `json:"edges"`
	task_subtasks *int
	selectValues  sql.SelectValues
}

// SubtaskEdges holds the relations/edges for other nodes in the graph.
type SubtaskEdges struct {
	// Task holds the value of the task edge.
	Task *Task `json:"task,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TaskOrErr returns the Task value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SubtaskEdges) TaskOrErr() (*Task, error) {
	if e.Task != nil {
		return e.Task, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: task.Label}
	}
	return nil, &NotLoadedError{edge: "task"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Subtask) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case subtask.FieldID:
			values[i] = new(sql.NullInt64)
		case subtask.FieldTitle, subtask.FieldDescription:
			values[i] = new(sql.NullString)
		case subtask.ForeignKeys[0]: // task_subtasks
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Subtask fields.
func (s *Subtask) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case subtask.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case subtask.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				s.Title = value.String
			}
		case subtask.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				s.Description = value.String
			}
		case subtask.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field task_subtasks", value)
			} else if value.Valid {
				s.task_subtasks = new(int)
				*s.task_subtasks = int(value.Int64)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Subtask.
// This includes values selected through modifiers, order, etc.
func (s *Subtask) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryTask queries the "task" edge of the Subtask entity.
func (s *Subtask) QueryTask() *TaskQuery {
	return NewSubtaskClient(s.config).QueryTask(s)
}

// Update returns a builder for updating this Subtask.
// Note that you need to call Subtask.Unwrap() before calling this method if this Subtask
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Subtask) Update() *SubtaskUpdateOne {
	return NewSubtaskClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Subtask entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Subtask) Unwrap() *Subtask {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Subtask is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Subtask) String() string {
	var builder strings.Builder
	builder.WriteString("Subtask(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("title=")
	builder.WriteString(s.Title)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(s.Description)
	builder.WriteByte(')')
	return builder.String()
}

// Subtasks is a parsable slice of Subtask.
type Subtasks []*Subtask
