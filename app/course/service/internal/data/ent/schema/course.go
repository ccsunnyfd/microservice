package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Course holds the schema definition for the Course entity.
type Course struct {
	ent.Schema
}

// Fields of the Course.
func (Course) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("title").Unique(),
		field.String("description").Optional(),
		field.Int64("teacher_id").Optional(),
	}
}

// Edges of the Course.
func (Course) Edges() []ent.Edge {
	return nil
}
