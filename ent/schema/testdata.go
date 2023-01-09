package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// TestData holds the schema definition for the TestData entity.
type TestData struct {
	ent.Schema
}

// Fields of the TestData.
func (TestData) Fields() []ent.Field {
	return []ent.Field{
		field.String("text"),
		field.Time("created_at"),
		field.Time("updated_at"),
	}
}

// Edges of the TestData.
func (TestData) Edges() []ent.Edge {
	return nil
}
