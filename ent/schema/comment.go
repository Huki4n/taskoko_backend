package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type Comment struct {
	ent.Schema
}

func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.Text("text").NotEmpty(),
		field.Time("created_at").Default(time.Now),
	}
}

func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("task", Task.Type).Ref("comments").Unique(),
		edge.To("author", User.Type).Unique().Required(),
	}
}
