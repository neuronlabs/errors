// Package errors provides leightweight error handling and classification primitives.
//
// The package defines blazingly fast classification system.
// A class is composed of the major, minor and index subclassifications.
// Each subclassifaction has different bitwise length with total of 32 bits.
// Thus a Class is a wrapper over uint32.
// A major is composed of 8, minor 10 and index of 14 bits.
//
// Example:
// Class with decimal value of 44205263, in a binary form equals to
//  00000010101000101000010011001111 which decomposes into:
//	00000010 - major (8 bit)
//		    1010001010 - minor (10 bit)
//					  00010011001111 - index (14 bit)
//
// The class concept was inspired by the need of multiple errors
// with the same logic but different messages.
//
// The package provides simple error handling interfaces and functions.
// It allows to create simple and detailed classified errors.
package errors
