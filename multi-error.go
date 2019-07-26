package errors

import (
	"strings"
)

// MultiError is the slice of errors parsable into a single error.
type MultiError []ClassError

// Error implements error interface.
func (m MultiError) Error() string {
	sb := &strings.Builder{}

	for i, e := range m {
		sb.WriteString(e.Error())
		if i != len(m)-1 {
			sb.WriteString(",")
		}
	}
	return sb.String()
}

// HasMajor checks if provided 'mjr' occurs in given multi error slice.
func (m MultiError) HasMajor(mjr Major) bool {
	for _, err := range m {
		if err.Class().Major() == mjr {
			return true
		}
	}
	return false
}
