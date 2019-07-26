package errors

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

// TestDetailedError tests detailed error functions.
func TestDetailedError(t *testing.T) {
	resetContainer()

	message := "some testing message"
	first := NewDet(ClInvalidIndex, message)
	second := NewDetf(ClInvalidIndex, "formatted: '%d'", 2)

	assert.Equal(t, "some testing message", first.Error())
	assert.Equal(t, "formatted: '2'", second.Error())

	// check operations
	firstOperation := "github.com/neuronlabs/errors.TestDetailedError#detailed_test.go:14"
	secondOperation := "github.com/neuronlabs/errors.TestDetailedError#detailed_test.go:15"
	assert.Equal(t, firstOperation, first.Operation())
	assert.Equal(t, secondOperation, second.Operation())

	second.AppendOperation(first.Operation())
	assert.Equal(t, secondOperation+"|"+firstOperation, second.Operation())

	assert.NotEqual(t, first.ID(), second.ID())

	assert.Equal(t, ClInvalidIndex, first.Class())

	detail := "This is detail."
	first.SetDetails(detail)

	assert.Equal(t, detail, first.Details())
	first.WrapDetails("Wrapped.")
	assert.Equal(t, "Wrapped. This is detail.", first.Details())

	second.SetDetailsf("This is %dnd detail.", 2)
	assert.Equal(t, "This is 2nd detail.", second.Details())

	second.WrapDetailsf("Wrapped %dnd.", 2)
	assert.Equal(t, "Wrapped 2nd. This is 2nd detail.", second.Details())

	sd, ok := second.(*detailedError)
	require.True(t, ok)

	sd.details = ""
	second.WrapDetails("Should be stored.")

	assert.Equal(t, "Should be stored.", sd.Details())
}
