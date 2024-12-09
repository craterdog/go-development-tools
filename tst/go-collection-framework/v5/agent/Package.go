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

/*
Package "agent" declares a set of agents that operate on values that have a
generic type.  They are used by the collection classes declared in this Go
module.

For detailed documentation on this package refer to the wiki:
  - https://github.com/craterdog/go-collection-framework/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-class-model/wiki

Additional concrete implementations of the classes declared by this package can
be developed and used seamlessly since the interface declarations only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package agent

import ()

// TYPE DECLARATIONS

/*
Rank is a constrained type representing the possible rankings for two values.
*/
type Rank uint8

const (
	LesserRank Rank = iota
	EqualRank
	GreaterRank
)

// FUNCTIONAL DECLARATIONS

/*
RankingFunction[V any] is a functional type that declares the signature for any
function that can determine the relative ranking of two values.
*/
type RankingFunction[V any] func(
	first V,
	second V,
) Rank

// CLASS DECLARATIONS

/*
CollatorClassLike[V any] is a class interface that declares the complete set
of class constructors, constants and functions that must be supported by each
concrete collator-like class.

An optional maximum depth may be specified that limits the depth of the
structures being collated in a way that avoids possible infinite recursion.
The default maximum depth is 16.
*/
type CollatorClassLike[V any] interface {
	// Constructor Methods
	Make() CollatorLike[V]
	MakeWithMaximumDepth(
		maximumDepth int,
	) CollatorLike[V]
}

/*
IteratorClassLike[V any] is a class interface that declares the complete set
of class constructors, constants and functions that must be supported by each
concrete iterator-like class.

An iterator-like class can be used to iterate over the specified array of
values.  The array is copied so that it is immutable during iteration.
*/
type IteratorClassLike[V any] interface {
	// Constructor Methods
	Make(
		values []V,
	) IteratorLike[V]
}

/*
SorterClassLike[V any] is a class interface that declares the complete set
of class constructors, constants and functions that must be supported by each
concrete sorter-like class.

An optional ranking function may be specified that is used to compare values.
The default ranking function uses the value's "natural" ordering based on its
type.
*/
type SorterClassLike[V any] interface {
	// Constructor Methods
	Make() SorterLike[V]
	MakeWithRanker(
		ranker RankingFunction[V],
	) SorterLike[V]
}

// INSTANCE DECLARATIONS

/*
CollatorLike[V any] is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete collator-like class.

A collator-like class is capable of recursively comparing and ranking two values
of any type.  The maximum depth attribute removes the possibility of infinite
recursion in the case of cycles.
*/
type CollatorLike[V any] interface {
	// Principal Methods
	GetClass() CollatorClassLike[V]
	CompareValues(
		first V,
		second V,
	) bool
	RankValues(
		first V,
		second V,
	) Rank

	// Attribute Methods
	GetMaximumDepth() int
}

/*
IteratorLike[V any] is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete iterator-like class.

An iterator-like class can be used to move forward and backward over the values
in an array.  It implements the Gang of Four (GoF) Iterator Design Pattern:
  - https://en.wikipedia.org/wiki/Iterator_pattern

An iterator agent locks into the slots that reside between each value in the
sequence:

	  . [value 1] . [value 2] . [value 3] ... [value N] .
	  ^           ^           ^                         ^
	slot 0      slot 1      slot 2                    slot N

It moves from slot to slot and has access to the values (if they exist) on each
side of the slot.  At each slot an iterator has access to the previous value
and next value in the array (assuming they exist). The slot at the start of
the array has no PREVIOUS value, and the slot at the end of the array has no
NEXT value.
*/
type IteratorLike[V any] interface {
	// Principal Methods
	GetClass() IteratorClassLike[V]
	IsEmpty() bool
	ToStart()
	ToSlot(
		slot int,
	)
	ToEnd()
	GetNext() V
	GetPrevious() V
	HasNext() bool
	HasPrevious() bool

	// Attribute Methods
	GetSize() int
	GetSlot() int
}

/*
SorterLike[V any] is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete sorter-like class.

A sorter-like class implements a specific sorting algorithm.  It uses a ranking
function to correlate the values.  If no ranking function is specified the
values are sorted into their "natural" ordering by type of value.
*/
type SorterLike[V any] interface {
	// Principal Methods
	GetClass() SorterClassLike[V]
	SortValues(
		values []V,
	)
	ReverseValues(
		values []V,
	)
	ShuffleValues(
		values []V,
	)

	// Attribute Methods
	GetRanker() RankingFunction[V]
}

// ASPECT DECLARATIONS
