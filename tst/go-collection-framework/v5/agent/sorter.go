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

func SorterClass[V any]() SorterClassLike[V] {
	return sorterClass[V]()
}

// Constructor Methods

func (c *sorterClass_[V]) Sorter() SorterLike[V] {
	var instance = &sorter_[V]{
		// Initialize the instance attributes.
	}
	return instance
}

func (c *sorterClass_[V]) SorterWithRanker(
	ranker RankingFunction[V],
) SorterLike[V] {
	if uti.IsUndefined(ranker) {
		panic("The \"ranker\" attribute is required by this class.")
	}
	var instance = &sorter_[V]{
		// Initialize the instance attributes.
		ranker_: ranker,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *sorter_[V]) GetClass() SorterClassLike[V] {
	return sorterClass[V]()
}

func (v *sorter_[V]) SortValues(
	values []V,
) {
	// TBD - Add the method implementation.
}

func (v *sorter_[V]) ReverseValues(
	values []V,
) {
	// TBD - Add the method implementation.
}

func (v *sorter_[V]) ShuffleValues(
	values []V,
) {
	// TBD - Add the method implementation.
}

// Attribute Methods

func (v *sorter_[V]) GetRanker() RankingFunction[V] {
	return v.ranker_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type sorter_[V any] struct {
	// Declare the instance attributes.
	ranker_ RankingFunction[V]
}

// Class Structure

type sorterClass_[V any] struct {
	// Declare the class constants.
}

// Class Reference

var sorterMap_ = map[string]any{}
var sorterMutex_ syn.Mutex

func sorterClass[V any]() *sorterClass_[V] {
	// Generate the name of the bound class type.
	var class *sorterClass_[V]
	var name = fmt.Sprintf("%T", class)

	// Check for an existing bound class type.
	sorterMutex_.Lock()
	var value = sorterMap_[name]
	switch actual := value.(type) {
	case *sorterClass_[V]:
		// This bound class type already exists.
		class = actual
	default:
		// Add a new bound class type.
		class = &sorterClass_[V]{
			// Initialize the class constants.
		}
		sorterMap_[name] = class
	}
	sorterMutex_.Unlock()

	// Return a reference to the bound class type.
	return class
}
