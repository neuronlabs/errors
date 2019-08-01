package errors

// Major is the highest level subclassification.
// It is of maximum 8 bit size, which gives 2^8 - 256 combinations.
type Major uint8

// Valid checks if the major value isn't greater than the allowed size
// and if it's value is non-zero.
func (m Major) Valid() bool {
	return m != 0
}

// NewMajor creates new Major classification.
func NewMajor() (Major, error) {
	return container.newMajor()
}

// MustNewMajor creates new major error classification.
// Panics if reached maximum number of possible majors.
func MustNewMajor() Major {
	mjr, err := container.newMajor()
	if err != nil {
		panic(err)
	}
	return mjr
}

// Minor is mid level error subclassification.
// It is a 10 bit long value, which give 2^10 - 1024 - combinations
// for each major.
type Minor uint16

// Valid checks if the Minor is valid.
func (m Minor) Valid() bool {
	return m>>minorBitSize == 0 && m&maxMinorValue != 0
}

// NewMinor creates new Minor error classification
// for provided 'mjr' Major.
// Returns error if the 'mjr' is not valid.
func NewMinor(mjr Major) (Minor, error) {
	return container.newMinor(mjr)
}

// MustNewMinor creates new Minor error classification
// for provided 'mjr' Major.
// Panics if the 'mjr' is not valid.
func MustNewMinor(mjr Major) Minor {
	mnr, err := container.newMinor(mjr)
	if err != nil {
		panic(err)
	}
	return mnr
}

// Index is a 14 bit length lowest level error classification.
// It defines the most accurate class division.
// It's maximum size gives 2^14 - 16384 - index combinations for each minor.
type Index uint16

// Valid checks if the provided index is valid.
func (i Index) Valid() bool {
	return i>>indexBitSize == 0 && i&maxIndexValue != 0
}

// NewIndex creates new Index for the 'mjr' Major and 'mnr' Minor.
// Returns error if the 'mjr' or 'mnr' are not valid.
func NewIndex(mjr Major, mnr Minor) (Index, error) {
	return container.newIndex(mjr, mnr)
}

// MustNewIndex creates new Index for the 'mjr' Major and 'mnr' Minor.
// Panics if 'mjr' or 'mnr' are not valid.
func MustNewIndex(mjr Major, mnr Minor) Index {
	index, err := container.newIndex(mjr, mnr)
	if err != nil {
		panic(err)
	}
	return index
}
