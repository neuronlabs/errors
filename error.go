package errors

import (
	"fmt"
)

var _ ClassError = &simpleError{}

type simpleError struct {
	class Class
	msg   string
}

// New creates simple ClassError for provided 'c' Class and 'msg' message.
func New(c Class, msg string) ClassError {
	return &simpleError{c, msg}
}

// Newf creates simple formatted ClassError for provided 'c' Class, 'format' and arguments 'args'.
func Newf(c Class, format string, args ...interface{}) ClassError {
	return &simpleError{c, fmt.Sprintf(format, args...)}
}

// Error implements error interface.
func (s *simpleError) Error() string {
	return s.msg
}

// Class implements ClassError interface.
func (s *simpleError) Class() Class {
	return s.class
}
