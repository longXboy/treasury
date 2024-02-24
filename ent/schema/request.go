package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Request holds the schema definition for the Request entity.
type Request struct {
	ent.Schema
}

// Fields of the Request.
func (Request) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable(),
		field.String("status").Comment("pending, approved, rejected").Default("pending"),
		field.Int64("amount").Comment("amount of token requested").Default(0),
		field.String("recipient").Comment("recipient's account address"),
		field.String("tx_hash").Comment("transaction hash of token transfer").Default(""),
		field.Int64("nonce").Comment("nonce of token transfer").Default(-1),
		field.Bool("executed").Comment("whether the request has been executed successfully").Default(false),
	}
}

// Edges of the Request.
func (Request) Edges() []ent.Edge {
	return nil
}
