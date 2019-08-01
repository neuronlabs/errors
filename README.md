![Neuron Logo](logo.svg)

# Errors [![Go Report Card](https://goreportcard.com/badge/github.com/neuronlabs/errors)](https://goreportcard.com/report/github.com/neuronlabs/errors) [![GoDoc](https://godoc.org/github.com/neuronlabs/errors?status.svg)](https://godoc.org/github.com/neuronlabs/errors) [![Build Status](https://travis-ci.com/neuronlabs/errors.svg?branch=master)](https://travis-ci.com/neuronlabs/errors) [![Coverage Status](https://coveralls.io/repos/github/neuronlabs/errors/badge.svg?branch=master)](https://coveralls.io/github/neuronlabs/errors?branch=master) ![License](https://img.shields.io/github/license/neuronlabs/errors.svg)

Package errors provides simple golang error and classification primitives.

* [Class](#class)
* [Interfaces](#interfaces)
* [Error Handling](#error-handling)
* [Example](#example)
* [Links](#links)

## Class

The package defines blazingly fast classification system.
A `Class` is an uint32 wrapper, composed of the `Major`, `Minor` and `Index` subclassifications.
Each subclassifaction has different bitwise length.
A major is composed of 8, minor 10 and index of 14 bits - total 32bits.

Example:

```Class with decimal value of 44205263, in a binary form equals to
 00000010101000101000010011001111 which decomposes into:
 00000010 - major (8 bit)
         1010001010 - minor (10 bit)
                   00010011001111 - index (14 bit)
```

The class concept was inspired by the need of multiple errors
with the same logic but different messages.

A class might be composed in three different ways:

* Major only - the class is Major singleton.
* Major, Minor only - classes that don't need triple subclassification divison.
* Major, Minor, Index - classes that decomposes 

Use `NewMajor` or `MustNewMajor` functions to create `Major`, `NewMinor` or `MustNewMinor` for new `Minor` and `NewIndex` or `MustNewIndex` for new `Index`.



## Interfaces

The package provides simple error handling interfaces and functions.
It allows to create simple and detailed classified errors.

## ClassError

A `ClassError` is the interface that provides error classification with the`Class` method.

## DetailedError

`DetailedError` is the interface used for errors that stores and handles human readable details, contains it's instance id and runtime call operation.
Implements `ClassError`, `Detailer`, `Operationer`, `Indexer`, `error` interfaces.

### Detailer

`Detailer` interface allows to set and get the human readable details - full sentences.

### Operationer

`OperationError` is the interface used to get the runtime operation information.

### Indexer

`Indexer` is the interface used to obtain 'ID' for each error instance.


## Error handling

This package contains two error structure implementations: 

* [Simple Error Structure](#simple-error)
* [Detailed Error Structure](#detailed-error)

### Simple Error

A simple error implements `ClassError` interface. It is lightweight error that contains only a message and it's class.

Created by the `New` and `Newf` functions.

**Example:**
```go
import "github.com/neuronlabs/errors"
// let's assume we have some ClassInvalidRequest already defined.
var ClassInvalidInput errors.Class

func createValue(input int) error {
    if input < 0 {
        return errors.New(ClassInvalidInput, "provided input lower than zero")
    }

    if input > 50 {
        return errors.Newf(ClassInvalidInput, "provided input value: '%d' is not valid", input) 
    }
    // do the logic here
    return nil
}
```

### Detailed Error

The detailed error struct (`detailedError`) implements `DetailedError`.

It contains a lot of information about given error instance:

* Human readable `Details`
* Runtime function call `Operations`
* Unique error instance `ID` 

In order to create detailed error use the `NewDet` or `NewDetf` functions.

### Example

```go
import (
    "fmt"
    "os"

    "github.com/neuronlabs/errors"
)

var (
    ClInputInvalidValue errors.Class
    ClInputNotProvided  errors.Class
)

func init() {
    initClasses()
}

func initClasses() {
    inputMjr := errors.MustNewMajor()
    invalidValueMnr := errors.MustNewMinor(inputMjr)
    ClInputInvalidValue = errors.MustNewMinorClass(inputMjr, invalidValueMnr)
    
    inputNotProvidedMnr := errors.MustNewMinor(inputMjr)
    ClInputNotProvided = errors.MustNewMinorClass(inputMjr, inputNotProvidedMnr)
}


func main() {
    input, err := getInput()
    if err == nil {
        // Everything is fine.
        os.Exit(0)
    }

    if classed, ok := err.(errors.ClassError); ok {
        if classed.Class() == ClInputNotProvided {
            fmt.Println("No required integer arguments provided.")
            os.Exit(1)
        }
    }

    var details string
    detailed, ok := err.(errors.DetailedError)
    if ok {
        details = detailed.Details()
    } else {
        details = err.Error()
    }
    fmt.Printf("Invalid input value provided: '%s'\n", details)
    os.Exit(1)    
}


func checkInput(input int) error {
    if input < 0 {
        err := errors.NewDet(ClassInputInvalidValue, "provided input lower than zero")        
        err.SetDetailsf("The input value provided to the function is invalid. The value must be greater than zero.")
        return err
    }

    if input > 50 {
        err := errors.NewDetf(ClassInvalidInput, "provided input value: '%d' is not valid", input) 
        err.SetDetailsf("The input value: '%d' provided to the function is invalid. The value can't be greater than '50'.", input)
        return err
    }
    // do the logic here
    return nil
}



func getInput() (int, error) {
    if len(os.Args) == 0 {
        return errors.New(ClInputNotProvided, "no input provided")
    }

    input, err := strconv.Atoi(os.Args[0])
    if err != nil {
        err := errors.NewDetf(ClInputInvalidValue, "provided input is not an integer")        
        err.SetDetail(err.Error())
        return 0, err
    }

    if err = checkInput(input); err != nil {
        return 0, err
    }
    return input, nil
}
```

## Links

* [Neuron-Core](https://github.com/neuronlabs/neuron-core)
* [Docs](https://docs.neuronlabs.io/errors)
