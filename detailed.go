package errors

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/google/uuid"
)

// compile time check for detailedError interfaces.
var (
	_ ClassError    = &detailedError{}
	_ Detailer      = &detailedError{}
	_ Operationer   = &detailedError{}
	_ Indexer       = &detailedError{}
	_ DetailedError = &detailedError{}
)

// detailedError is the class based error definition.
// Each instance has it's own trackable ID. It's chainable
// It contains also a Class variable that might be comparable in logic.
type detailedError struct {
	// ID is a unique error instance identification number.
	id uuid.UUID
	// class defines the error classification.
	class Class
	// details contains the detailed information.
	details string
	// message is a message used as a string for the
	// golang error interface implementation.
	message string
	// Opertaion is the operation name when the error occurred.
	operation string
}

// NewDet creates DetailedError with given 'class' and message 'message'.
func NewDet(c Class, message string) DetailedError {
	err := newDetailed(c)
	err.message = message
	return err
}

// NewDetf creates DetailedError instance with provided 'class' with formatted message.
// DetailedError implements ClassError interface.
func NewDetf(c Class, format string, args ...interface{}) DetailedError {
	err := newDetailed(c)
	err.message = fmt.Sprintf(format, args...)
	return err
}

// Class implements ClassError interface.
func (e *detailedError) Class() Class {
	return e.class
}

// Details implements DetailedError interface.
func (e *detailedError) Details() string {
	return e.details
}

// DetailedError implements error interface.
func (e *detailedError) Error() string {
	return e.message
}

// ID implements IndexedError interface.
func (e *detailedError) ID() uuid.UUID {
	return e.id
}

// Operation implements OperationError interface.
func (e *detailedError) Operation() string {
	return e.operation
}

// SetDetails sets the error 'detail' and returns itself.
func (e *detailedError) SetDetails(detail string) {
	e.details = detail
}

// SetDetailsf sets the error's formatted detail with provided and returns itself.
func (e *detailedError) SetDetailsf(format string, args ...interface{}) {
	e.details = fmt.Sprintf(format, args...)
}

// WrapDetail wraps the 'detail' for given error. Wrapping appends the new detail
// to the front of error detail message.
func (e *detailedError) WrapDetails(detail string) {
	e.wrapDetail(detail)
}

// WrapDetailf wraps the detail with provided formatting for given error.
// Wrapping appends the new detail to the front of error detail message.
func (e *detailedError) WrapDetailsf(format string, args ...interface{}) {
	e.wrapDetail(fmt.Sprintf(format, args...))
}

// AppendOperation wraps the 'operation' by concantinating 'e' Operation
// to its value. It create a chain of operation call.
func (e *detailedError) AppendOperation(operation string) {
	e.operation += "|" + operation
}

func (e *detailedError) wrapDetail(detail string) {
	if e.details == "" {
		e.details = detail
	} else {
		e.details = detail + " " + e.details
	}
}

func newDetailed(c Class) *detailedError {
	err := &detailedError{
		id:    uuid.New(),
		class: c,
	}
	pc, _, _, ok := runtime.Caller(2)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		file, line := details.FileLine(pc)
		_, singleFile := filepath.Split(file)
		err.operation = details.Name() + "#" + singleFile + ":" + strconv.Itoa(line)
	}
	return err
}
