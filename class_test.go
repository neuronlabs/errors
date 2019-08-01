package errors

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

// TestClass tests creation of the `Major` subclassification.
func TestClass(t *testing.T) {
	t.Run("Major", func(t *testing.T) {
		resetEmptyContainer()

		firstMjr, err := NewMajor()
		require.NoError(t, err)
		assert.Equal(t, 1, int(firstMjr))

		var secondMjr Major
		assert.NotPanics(t, func() { secondMjr = MustNewMajor() })
		assert.Equal(t, 2, int(secondMjr))

		firstClass, err := NewMajorClass(firstMjr)
		if assert.NoError(t, err) {
			assert.Equal(t, firstMjr, firstClass.Major())
		}

		var secondClass Class
		assert.NotPanics(t, func() { secondClass = MustNewMajorClass(secondMjr) })
		assert.Equal(t, secondMjr, secondClass.Major())

		var notDefinedMajor Major
		assert.Panics(t, func() { MustNewMajorClass(notDefinedMajor) })

		minorsLen := len(container.minors)
		indexesLen := len(container.indexes)
		topMjr := container.major
		for i := 0; i < 20; i++ {
			NewMajor()
		}
		assert.NotEqual(t, minorsLen, len(container.minors))
		assert.NotEqual(t, indexesLen, len(container.indexes))
		assert.Equal(t, topMjr+20, container.major)

		// on testing purpose change top major to max uint8 value.
		maxMajor := Major((2 << 7) - 1)
		container.major = maxMajor

		_, err = NewMajor()
		require.Error(t, err)

		assert.Panics(t, func() { MustNewMajor() })
	})

	t.Run("Minor", func(t *testing.T) {
		resetEmptyContainer()

		mjr, err := NewMajor()
		require.NoError(t, err)
		assert.Equal(t, 1, int(mjr))

		mnr1, err := NewMinor(mjr)
		if assert.NoError(t, err) {
			assert.Equal(t, 1, int(mnr1))
		}

		var mnr2 Minor
		getMnr2 := func() { mnr2 = MustNewMinor(mjr) }

		assert.NotPanics(t, getMnr2)
		assert.Equal(t, 2, int(mnr2))

		assert.Panics(t, func() { MustNewMinor(Major(0)) })

		firstClass, err := NewMinorClass(mjr, mnr1)
		if assert.NoError(t, err) {
			assert.Equal(t, mjr, firstClass.Major())
			assert.Equal(t, mnr1, firstClass.Minor())
			assert.False(t, firstClass.Index().Valid())
		}

		var secondClass Class
		assert.NotPanics(t, func() { secondClass = MustNewMinorClass(mjr, mnr2) })
		assert.Equal(t, mjr, secondClass.Major())
		assert.Equal(t, mnr2, secondClass.Minor())
		assert.False(t, secondClass.Index().Valid(), "%b", secondClass)

		var (
			invMjr Major
			invMnr Minor
		)
		assert.False(t, invMjr.Valid())
		assert.False(t, invMnr.Valid())

		_, err = NewMinorClass(invMjr, mnr1)
		assert.Error(t, err)

		assert.Panics(t, func() { MustNewMinorClass(mjr, invMnr) })

		initLen := len(container.indexes[mjr])
		for i := 0; i < 20; i++ {
			_, err = NewMinor(mjr)
			assert.NoError(t, err)
		}

		assert.NotEqual(t, initLen, len(container.indexes[mjr]))

		container.minors[mjr] = maxMinorValue
		_, err = NewMinor(mjr)
		require.Error(t, err)

		assert.Panics(t, func() { MustNewMinor(mjr) })
	})

	t.Run("Index", func(t *testing.T) {
		resetEmptyContainer()

		mjr, err := NewMajor()
		require.NoError(t, err)
		assert.True(t, mjr.Valid())

		mnr, err := NewMinor(mjr)
		require.NoError(t, err)

		inx1, err := NewIndex(mjr, mnr)
		if assert.NoError(t, err) {
			assert.Equal(t, 1, int(inx1))
			assert.True(t, inx1.Valid())
		}

		var inx2 Index
		assert.False(t, inx2.Valid())

		if assert.NotPanics(t, func() { inx2 = MustNewIndex(mjr, mnr) }) {
			assert.Equal(t, 2, int(inx2))
			assert.True(t, inx2.Valid())
		}

		var (
			invMjr Major
			invMnr Minor
		)
		require.False(t, invMjr.Valid())
		require.False(t, invMnr.Valid())

		_, err = NewIndex(invMjr, mnr)
		assert.Error(t, err)

		assert.Panics(t, func() { MustNewIndex(mjr, invMnr) })

		inx1Class, err := NewClass(mjr, mnr, inx1)
		if assert.NoError(t, err) {
			assert.Equal(t, mjr, inx1Class.Major())
			assert.Equal(t, mnr, inx1Class.Minor())
			assert.Equal(t, inx1, inx1Class.Index())
		}

		var inx2Class Class
		assert.False(t, inx2Class.Major().Valid())
		assert.False(t, inx2Class.Minor().Valid())
		assert.False(t, inx2Class.Index().Valid())

		if assert.NotPanics(t, func() { inx2Class = MustNewClass(mjr, mnr, inx2) }) {
			assert.Equal(t, mjr, inx2Class.Major())
			assert.Equal(t, mnr, inx2Class.Minor())
			assert.Equal(t, inx2, inx2Class.Index())
		}

		assert.NotEqual(t, inx2Class, inx1Class)

		var invIndex Index
		assert.Panics(t, func() { MustNewClass(mjr, invMnr, inx2) })
		assert.Panics(t, func() { MustNewClass(invMjr, mnr, inx2) })
		assert.Panics(t, func() { MustNewClass(mjr, mnr, invIndex) })

		c, err := NewClassWIndex(mjr, mnr)
		require.NoError(t, err)

		var c2 Class

		assert.NotPanics(t, func() { c2 = MustNewClassWIndex(mjr, mnr) })
		assert.NotEqual(t, c, c2)

		_, err = NewClassWIndex(invMjr, mnr)
		assert.Error(t, err)

		_, err = NewClassWIndex(mjr, invMnr)
		assert.Error(t, err)

		assert.Panics(t, func() { MustNewClassWIndex(invMjr, mnr) })

		container.indexes[mjr][mnr] = maxIndexValue
		_, err = NewIndex(mjr, mnr)

		require.Error(t, err)
		assert.Panics(t, func() { MustNewIndex(mjr, mnr) })
	})
}

func resetContainer() {
	container = &classContainer{}
	initClasses()
}

func resetEmptyContainer() {
	container = &classContainer{}
}
