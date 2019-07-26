package errors

import (
	"github.com/google/uuid"
)

// ClassError is the interface used for all errors
// that uses classification system.
type ClassError interface {
	error
	// Class gets current error classification.
	Class() Class
}

// Detailer is the interface that defines methods used for setting details (errors).
type Detailer interface {
	// Details are human readable information about the error.
	Details() string
	// SetDetails sets the detail for given error.
	SetDetails(message string)
	// SetDetailsf sets formatted detail function.
	SetDetailsf(format string, args ...interface{})
	// WrapDetails if Detailer contains any 'Details' then it would
	// be combined with the 'message'. Otherwise works as SetDetails.
	WrapDetails(message string)
	// WrapDetailsf if Detailer contains any 'Details' then it would
	// be combined with the formatted message. Otherwise works as SetDetailsf.
	WrapDetailsf(format string, args ...interface{})
}

// DetailedError is the error that implements
// ClassError, Detailer, Indexer, Operationer interfaces.
type DetailedError interface {
	ClassError
	Indexer
	Detailer
	Operationer
}

// Indexer is the an enhanced error interface.
type Indexer interface {
	// ID gets a unique error instance identification number.
	ID() uuid.UUID
}

// Operationer is enhanced error interface.
// It contains and allows to get the run time operation name.
type Operationer interface {
	// Operation gets the runtime information about the file:line and function
	// where the error was created.
	Operation() string
	// AppendOperation wraps the operation creating a chain of operations.
	AppendOperation(operation string)
}
