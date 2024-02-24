// Code generated by ent, DO NOT EDIT.

package request

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the request type in the database.
	Label = "request"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldRecipient holds the string denoting the recipient field in the database.
	FieldRecipient = "recipient"
	// FieldTxHash holds the string denoting the tx_hash field in the database.
	FieldTxHash = "tx_hash"
	// Table holds the table name of the request in the database.
	Table = "requests"
)

// Columns holds all SQL columns for request fields.
var Columns = []string{
	FieldID,
	FieldStatus,
	FieldAmount,
	FieldRecipient,
	FieldTxHash,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Request queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByAmount orders the results by the amount field.
func ByAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmount, opts...).ToFunc()
}

// ByRecipient orders the results by the recipient field.
func ByRecipient(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRecipient, opts...).ToFunc()
}

// ByTxHash orders the results by the tx_hash field.
func ByTxHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTxHash, opts...).ToFunc()
}