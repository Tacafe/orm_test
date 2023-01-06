package schema

import "entgo.io/ent"

// Data holds the schema definition for the Data entity.
type Data struct {
	ent.Schema
}

// Fields of the Data.
func (Data) Fields() []ent.Field {
	return nil
}

// Edges of the Data.
func (Data) Edges() []ent.Edge {
	return nil
}
