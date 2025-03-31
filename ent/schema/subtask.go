package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Subtask struct {
	ent.Schema
}

func (Subtask) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.Text("description").NotEmpty(),
	}
}

func (Subtask) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("task", Task.Type).Ref("subtasks").Unique(),
	}
}
