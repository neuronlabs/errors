package errors

import (
	"github.com/google/uuid"
)

// ClassError is the error interface used for all errors
// that uses classification system.
type ClassError interface {
	error
	// Class gets current error classification.
	Class() Class
}

// DetailedError is an enhanced ClassError interface.
// It defines methods used for setting error details.
type DetailedError interface {
	ClassError
	// Details are human readable information about the error.
	Details() string
	// SetDetails sets the detail for given error. Callbacks itself on return.
	SetDetails(message string) DetailedError
	// SetDetailsf sets formatted detail function. Callbacks itself on return.
	SetDetailsf(format string, args ...interface{}) DetailedError
}

// IndexedError is the an enhanced ClassError interface.
type IndexedError interface {
	ClassError
	// ID gets a unique error instance identification number.
	ID() uuid.UUID
}

// OperationError is enhanced ClassError interface.
// It contains and allows to get the run time operation name.
type OperationError interface {
	ClassError
	// Operation gets the runtime information about the file:line and function
	// where the error was created.
	Operation() string
	// WrapOperation wraps the operation creating a chain of operations.
	WrapOperation(operation string)
}
