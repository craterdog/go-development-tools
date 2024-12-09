/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package agent

import (
	fmt "fmt"
	uti "github.com/craterdog/go-missing-utilities/v2"
	syn "sync"
)

// CLASS INTERFACE

// Access Function

func IteratorClass[V any]() IteratorClassLike[V] {
	return iteratorClassReference[V]()
}

// Constructor Methods

func (c *iteratorClass_[V]) Make(
	values []V,
) IteratorLike[V] {
	if uti.IsUndefined(values) {
		panic("The \"values\" attribute is required by this class.")
	}
	var instance = &iterator_[V]{
		// Initialize the instance attributes.
		values_: values,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *iterator_[V]) GetClass() IteratorClassLike[V] {
	return iteratorClassReference[V]()
}

func (v *iterator_[V]) IsEmpty() bool {
	var result_ bool
	// TBD - Add the method implementation.
	return result_
}

func (v *iterator_[V]) ToStart() {
	// TBD - Add the method implementation.
}

func (v *iterator_[V]) ToSlot(
	slot int,
) {
	// TBD - Add the method implementation.
}

func (v *iterator_[V]) ToEnd() {
	// TBD - Add the method implementation.
}

func (v *iterator_[V]) GetNext() V {
	var result_ V
	// TBD - Add the method implementation.
	return result_
}

func (v *iterator_[V]) GetPrevious() V {
	var result_ V
	// TBD - Add the method implementation.
	return result_
}

func (v *iterator_[V]) HasNext() bool {
	var result_ bool
	// TBD - Add the method implementation.
	return result_
}

func (v *iterator_[V]) HasPrevious() bool {
	var result_ bool
	// TBD - Add the method implementation.
	return result_
}

// Attribute Methods

func (v *iterator_[V]) GetSize() int {
	return v.size_
}

func (v *iterator_[V]) GetSlot() int {
	return v.slot_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type iterator_[V any] struct {
	// Declare the instance attributes.
	size_   int
	slot_   int
	values_ []V
}

// Class Structure

type iteratorClass_[V any] struct {
	// Declare the class constants.
}

// Class Reference

var iteratorMap_ = map[string]any{}
var iteratorMutex_ syn.Mutex

func iteratorClassReference[V any]() *iteratorClass_[V] {
	// Generate the name of the bound class type.
	var class *iteratorClass_[V]
	var name = fmt.Sprintf("%T", class)

	// Check for an existing bound class type.
	iteratorMutex_.Lock()
	var value = iteratorMap_[name]
	switch actual := value.(type) {
	case *iteratorClass_[V]:
		// This bound class type already exists.
		class = actual
	default:
		// Add a new bound class type.
		class = &iteratorClass_[V]{
			// Initialize the class constants.
		}
		iteratorMap_[name] = class
	}
	iteratorMutex_.Unlock()

	// Return a reference to the bound class type.
	return class
}
