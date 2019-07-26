package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSimpleError tests simpleError structure.
func TestSimpleError(t *testing.T) {
	resetContainer()

	message := "testing message"
	simple1 := New(ClInvalidMajor, message)

	assert.Equal(t, ClInvalidMajor, simple1.Class())
	assert.Equal(t, message, simple1.Error())

	format := "two: %d"
	simple2 := Newf(ClInvalidMinor, format, 2)

	assert.Equal(t, ClInvalidMinor, simple2.Class())
	assert.Equal(t, "two: 2", simple2.Error())

	multi := MultiError([]ClassError{simple1, simple2})

	assert.Equal(t, "testing message,two: 2", multi.Error())
	assert.True(t, multi.HasMajor(ClInvalidMajor.Major()))

	var invalidMajor Major
	assert.False(t, multi.HasMajor(invalidMajor))
}
