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
		field.Int64("status").Comment("pending, approved, rejected"),
		field.Int64("amount").Comment("amount of token requested"),
		field.String("recipient").Comment("recipient's account address"),
		field.String("tx_hash").Comment("transaction hash of token transfer"),
		field.Int64("nonce").Comment("nonce of token transfer"),
		field.Bool("executed").Comment("whether the request has been executed successfully"),
	}
}

// Edges of the Request.
func (Request) Edges() []ent.Edge {
	return nil
}
