package errors

const (
	majorBitSize = 8
	minorBitSize = 10
	indexBitSize = 32 - majorBitSize - minorBitSize

	maxIndexValue = (2 << (indexBitSize - 1)) - 1
	maxMinorValue = (2 << (minorBitSize - 1)) - 1
	maxMajorValue = (2 << (majorBitSize - 1)) - 1

	majorMinorMask = uint32((2<<(majorBitSize+minorBitSize-1) - 1) << indexBitSize)
)

func init() {
	internalMajor := container.newMajor()

	invalidMajor, _ := container.newMinor(internalMajor)
	ClInvalidMajor = MustNewMinorClass(internalMajor, invalidMajor)

	invalidMinor, _ := container.newMinor(internalMajor)
	ClInvalidMinor = MustNewMinorClass(internalMajor, invalidMinor)

	invalidIndex, _ := container.newMinor(internalMajor)
	ClInvalidIndex = MustNewMinorClass(internalMajor, invalidIndex)
}

var (
	// ClInvalidMajor defines the invalid major error classification.
	ClInvalidMajor Class
	// ClInvalidMinor defines the invalid minor error classification.
	ClInvalidMinor Class
	// ClInvalidIndex defines the invalid index error classification.
	ClInvalidIndex Class
)

// Class is the  error classification model.
// It is composed of the major, minor and index subclassifications.
// Each subclassifaction is a different length number, where
// major is composed of 8, minor 10 and index of 14 bits.
// Example:
//  44205263 in a binary form is:
//
//  00000010101000101000010011001111 which decomposes into:
//	00000010 - major (8 bit) - 2
//		    1010001010 - minor (10 bit) - 650
//					  00010011001111 - index (14 bit) - 1231
//
// Major should be a global scope division like 'Repository', 'Marshaler', 'Controller' etc.
// Minor should divide the 'major' into subclasses like the Repository filter builders, marshaler - invalid field etc.
// Index is the most precise classification - i.e. Repository - filter builder - unsupported operator.
type Class uint32

// Index is a four digit number unique within given minor and major.
func (c Class) Index() Index {
	return c.index()
}

// Major is a single digit major classification.
func (c Class) Major() Major {
	return c.major()
}

// Minor is a double digit minor classification unique within given major.
func (c Class) Minor() Minor {
	return c.minor()
}

// MjrMnrMasked returns the class value masked by the major and minor value only.
func (c Class) MjrMnrMasked() uint32 {
	return uint32(c) & majorMinorMask
}

func (c Class) index() Index {
	return Index(c & maxIndexValue)
}

func (c Class) major() Major {
	return Major(c >> (32 - majorBitSize))
}

func (c Class) minor() Minor {
	return Minor(c >> (indexBitSize) & maxMinorValue)
}

// NewClass gets new class from the provided 'minor' and 'index'.
// If any of the arguments is not valid or out of bands the function returns an error.
func NewClass(mjr Major, mnr Minor, index Index) (Class, error) {
	return newClass(mjr, mnr, index)
}

// MustNewClass gets new class from the provided 'minor' and 'index'.
// Panics if any of the arguments is not valid or out of bands.
func MustNewClass(mjr Major, mnr Minor, index Index) Class {
	c, err := NewClass(mjr, mnr, index)
	if err != nil {
		panic(err)
	}
	return c
}

// NewMinorClass gets the class from provided 'minor'.
// The function gets minor's major and gets the major/minor class.
func NewMinorClass(mjr Major, mnr Minor) (Class, error) {
	return newMinorClass(mjr, mnr)
}

// MustNewMinorClass creates a class from the provided 'mjr' Major and 'mnr' Minor.
// Panics when any of the provided arguments is invalid.
func MustNewMinorClass(mjr Major, mnr Minor) Class {
	c, err := newMinorClass(mjr, mnr)
	if err != nil {
		panic(err)
	}
	return c
}

// NewMajorClass creates Class from the provided 'mjr' Major.
// This class contains zero valued 'Minor' and 'Index'.
// Returns error if the 'mjr' is invalid.
func NewMajorClass(mjr Major) (Class, error) {
	return newMajorClass(mjr)
}

// MustNewMajorClass creates Class from the provided 'mjr' Major.
// This class contains zero valued 'Minor' and 'Index'.
// Panics if provided 'mjr' is invalid.
func MustNewMajorClass(mjr Major) Class {
	c, err := newMajorClass(mjr)
	if err != nil {
		panic(err)
	}
	return c
}

func newMajorClass(mjr Major) (Class, error) {
	if !mjr.Valid() {
		return Class(0), NewDet(ClInvalidMajor, "provided invalid major")
	}
	return Class(uint32(mjr) << (32 - majorBitSize)), nil
}

func newMinorClass(mjr Major, mnr Minor) (Class, error) {
	if !mjr.Valid() {
		return Class(0), NewDet(ClInvalidMajor, "provided invalid major")
	}

	if !mnr.Valid() {
		return Class(0), NewDet(ClInvalidMinor, "provided invalid minor")
	}
	return Class(uint32(mjr)<<(32-majorBitSize) | uint32(mnr)<<(32-minorBitSize-majorBitSize)), nil
}

func newClass(mjr Major, mnr Minor, index Index) (Class, error) {
	if !mjr.Valid() {
		return Class(0), NewDet(ClInvalidMajor, "provided invalid major")
	}

	if !mnr.Valid() {
		return Class(0), NewDet(ClInvalidMinor, "provided invalid minor")
	}

	if !index.Valid() {
		return Class(0), NewDet(ClInvalidIndex, "provided invalid index")
	}
	return Class(uint32(mjr)<<(32-majorBitSize) | uint32(mnr)<<(32-minorBitSize-majorBitSize) | uint32(index)), nil
}
