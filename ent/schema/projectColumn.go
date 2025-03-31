package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ProjectColumn struct {
	ent.Schema
}

func (ProjectColumn) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Int("order"),
		field.String("color").NotEmpty(),
	}
}

func (ProjectColumn) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tasks", Task.Type),
		edge.From("project", Project.Type).Ref("columns").Unique(),
	}
}
