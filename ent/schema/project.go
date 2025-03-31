package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Project struct {
	ent.Schema
}

func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
	}
}

func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("columns", ProjectColumn.Type),
		edge.From("users", User.Type).Ref("projects"),
	}
}
