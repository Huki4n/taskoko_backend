package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type Task struct {
	ent.Schema
}

func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Int("order"),
		field.String("type"),
		field.String("name").NotEmpty(),
		field.Text("description"),
		field.Strings("tags"),
		field.Time("timer").Default(time.Now),
		field.String("image").Optional(),
		field.Bool("isDone").Default(false),
	}
}

func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("subtasks", Subtask.Type),
		edge.To("comments", Comment.Type),
		edge.To("attachments", Attachment.Type),
		edge.From("column", ProjectColumn.Type).Ref("tasks").Unique(),
		edge.From("users", User.Type).Ref("tasks"),
	}
}
