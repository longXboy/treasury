// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/longxboy/treasure/ent/request"
	"github.com/longxboy/treasure/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	requestFields := schema.Request{}.Fields()
	_ = requestFields
	// requestDescStatus is the schema descriptor for status field.
	requestDescStatus := requestFields[1].Descriptor()
	// request.DefaultStatus holds the default value on creation for the status field.
	request.DefaultStatus = requestDescStatus.Default.(string)
	// requestDescAmount is the schema descriptor for amount field.
	requestDescAmount := requestFields[2].Descriptor()
	// request.DefaultAmount holds the default value on creation for the amount field.
	request.DefaultAmount = requestDescAmount.Default.(int64)
	// requestDescExecuted is the schema descriptor for executed field.
	requestDescExecuted := requestFields[6].Descriptor()
	// request.DefaultExecuted holds the default value on creation for the executed field.
	request.DefaultExecuted = requestDescExecuted.Default.(bool)
}
