package errors

import (
	"sync"
)

var container = &classContainer{}

// classContainer is the container for the subclass boundaries and definitions.
type classContainer struct {
	sync.Mutex

	major   Major
	minors  []Minor
	indexes [][]Index
}

func (c *classContainer) newMajor() (Major, error) {
	c.Lock()
	defer c.Unlock()

	if c.major+1 == 0 {
		return 0, New(ClInvalidMajor, "reached maximum number of 'Major' classes")
	}
	c.major++

	c.resizeMinors(c.major)
	c.resizeIndexesMajor(c.major)
	return c.major, nil
}

func (c *classContainer) newMinor(mjr Major) (Minor, error) {
	c.Lock()
	defer c.Unlock()

	if !mjr.Valid() {
		return 0, New(ClInvalidMajor, "provided invalid Major")
	}

	c.resizeMinors(mjr)

	if c.minors[mjr] == maxMinorValue {
		return 0, Newf(ClInvalidMinor, "created maximum number of minors for major: '%d'", mjr)
	}
	c.minors[mjr]++

	c.resizeIndexesMinors(mjr, c.minors[mjr])
	return c.minors[mjr], nil
}

func (c *classContainer) newIndex(mjr Major, mnr Minor) (Index, error) {
	c.Lock()
	defer c.Unlock()

	if !mjr.Valid() {
		return 0, New(ClInvalidMajor, "provided invalid Major")
	}

	if !mnr.Valid() {
		return 0, New(ClInvalidMinor, "provided invalid Minor")
	}

	c.resizeIndexesMajor(mjr)
	c.resizeIndexesMinors(mjr, mnr)

	if c.indexes[mjr][mnr] == maxIndexValue {
		return 0, Newf(ClInvalidIndex, "reached maximum index subclass number for mjr: '%d', mnr: '%d'", mjr, mnr)
	}
	c.indexes[mjr][mnr]++
	return c.indexes[mjr][mnr], nil
}

func (c *classContainer) resizeMinors(mjr Major) {
	if int(mjr) < len(c.minors)-2 {
		return
	}

	size := len(c.minors) - 2
	if size <= 0 {
		size = 4
	}

	for size <= int(mjr) {
		size *= 2
	}

	temp := make([]Minor, size)
	copy(temp, c.minors)
	c.minors = temp
}

func (c *classContainer) resizeIndexesMajor(mjr Major) {
	if int(mjr) <= len(c.indexes)-2 {
		return
	}

	size := len(c.indexes) - 2
	if size <= 0 {
		size = 4
	}

	for size <= int(mjr) {
		size *= 2
	}
	temp := make([][]Index, size)
	copy(temp, c.indexes)
	c.indexes = temp
}

func (c *classContainer) resizeIndexesMinors(mjr Major, mnr Minor) {
	if int(mnr) <= len(c.indexes[mjr])-2 {
		return
	}

	size := len(c.indexes[mjr]) - 2
	if size <= 0 {
		size = 4
	}

	for size <= int(mnr) {
		size *= 2
	}
	temp := make([]Index, size)
	copy(temp, c.indexes[mjr])
	c.indexes[mjr] = temp
}
