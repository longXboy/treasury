package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Confirm holds the schema definition for the Confirm entity.
type Confirm struct {
	ent.Schema
}

// Fields of the Confirm.
func (Confirm) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable(),
		field.Int64("request_id").Comment("request id"),
		field.Int("manager_id").Comment("manager id"),
		field.Bool("approved").Comment("approve or reject").Immutable(),
	}
}

// Edges of the Confirm.
func (Confirm) Edges() []ent.Edge {
	return nil
}

// Edges of the Confirm.
func (Confirm) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("request_id", "manager_id").Unique(),
	}
}
