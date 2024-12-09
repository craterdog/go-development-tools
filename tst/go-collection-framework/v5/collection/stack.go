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

package collection

import (
	fmt "fmt"
	age "github.com/craterdog/go-collection-framework/v5/agent"
	uti "github.com/craterdog/go-missing-utilities/v2"
	syn "sync"
)

// CLASS INTERFACE

// Access Function

func StackClass[V any]() StackClassLike[V] {
	return stackClassReference[V]()
}

// Constructor Methods

func (c *stackClass_[V]) Make() StackLike[V] {
	var instance = &stack_[V]{
		// Initialize the instance attributes.
	}
	return instance
}

func (c *stackClass_[V]) MakeWithCapacity(
	capacity Size,
) StackLike[V] {
	if uti.IsUndefined(capacity) {
		panic("The \"capacity\" attribute is required by this class.")
	}
	var instance = &stack_[V]{
		// Initialize the instance attributes.
		capacity_: capacity,
	}
	return instance
}

func (c *stackClass_[V]) MakeFromArray(
	values []V,
) StackLike[V] {
	var instance StackLike[V]
	// TBD - Add the constructor implementation.
	return instance
}

func (c *stackClass_[V]) MakeFromSequence(
	values Sequential[V],
) StackLike[V] {
	var instance StackLike[V]
	// TBD - Add the constructor implementation.
	return instance
}

// Constant Methods

func (c *stackClass_[V]) DefaultCapacity() Size {
	return c.defaultCapacity_
}

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *stack_[V]) GetClass() StackClassLike[V] {
	return stackClassReference[V]()
}

func (v *stack_[V]) AddValue(
	value V,
) {
	// TBD - Add the method implementation.
}

func (v *stack_[V]) RemoveTop() V {
	var result_ V
	// TBD - Add the method implementation.
	return result_
}

func (v *stack_[V]) RemoveAll() {
	// TBD - Add the method implementation.
}

// Attribute Methods

func (v *stack_[V]) GetCapacity() Size {
	return v.capacity_
}

// Sequential[V] Methods

func (v *stack_[V]) IsEmpty() bool {
	var result_ bool
	// TBD - Add the method implementation.
	return result_
}

func (v *stack_[V]) GetSize() int {
	var result_ int
	// TBD - Add the method implementation.
	return result_
}

func (v *stack_[V]) AsArray() []V {
	var result_ []V
	// TBD - Add the method implementation.
	return result_
}

func (v *stack_[V]) GetIterator() age.IteratorLike[V] {
	var result_ age.IteratorLike[V]
	// TBD - Add the method implementation.
	return result_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type stack_[V any] struct {
	// Declare the instance attributes.
	capacity_ Size
}

// Class Structure

type stackClass_[V any] struct {
	// Declare the class constants.
	defaultCapacity_ Size
}

// Class Reference

var stackMap_ = map[string]any{}
var stackMutex_ syn.Mutex

func stackClassReference[V any]() *stackClass_[V] {
	// Generate the name of the bound class type.
	var class *stackClass_[V]
	var name = fmt.Sprintf("%T", class)

	// Check for an existing bound class type.
	stackMutex_.Lock()
	var value = stackMap_[name]
	switch actual := value.(type) {
	case *stackClass_[V]:
		// This bound class type already exists.
		class = actual
	default:
		// Add a new bound class type.
		class = &stackClass_[V]{
			// Initialize the class constants.
			// defaultCapacity_: constantValue,
		}
		stackMap_[name] = class
	}
	stackMutex_.Unlock()

	// Return a reference to the bound class type.
	return class
}
