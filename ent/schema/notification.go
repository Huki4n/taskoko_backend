package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Notification struct {
	ent.Schema
}

func (Notification) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.String("text").NotEmpty(),
		field.String("type").NotEmpty(),
	}
}
