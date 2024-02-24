// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ConfirmsColumns holds the columns for the "confirms" table.
	ConfirmsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "request_id", Type: field.TypeInt64},
		{Name: "manager_id", Type: field.TypeInt64},
		{Name: "approved", Type: field.TypeBool},
	}
	// ConfirmsTable holds the schema information for the "confirms" table.
	ConfirmsTable = &schema.Table{
		Name:       "confirms",
		Columns:    ConfirmsColumns,
		PrimaryKey: []*schema.Column{ConfirmsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "confirm_request_id_manager_id",
				Unique:  true,
				Columns: []*schema.Column{ConfirmsColumns[1], ConfirmsColumns[2]},
			},
		},
	}
	// RequestsColumns holds the columns for the "requests" table.
	RequestsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "status", Type: field.TypeString, Default: "pending"},
		{Name: "amount", Type: field.TypeInt64, Default: 0},
		{Name: "recipient", Type: field.TypeString},
		{Name: "tx_hash", Type: field.TypeString, Default: ""},
		{Name: "nonce", Type: field.TypeInt64, Default: -1},
		{Name: "executed", Type: field.TypeBool, Default: false},
	}
	// RequestsTable holds the schema information for the "requests" table.
	RequestsTable = &schema.Table{
		Name:       "requests",
		Columns:    RequestsColumns,
		PrimaryKey: []*schema.Column{RequestsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ConfirmsTable,
		RequestsTable,
	}
)

func init() {
}
