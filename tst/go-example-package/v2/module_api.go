/*
................................................................................
.                  Copyright (c) 2025.  All Rights Reserved.                   .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

/*
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│         This "module_api.go" file was automatically generated using:         │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
│      Updates to any part of this file—other than the Module Description      │
│             and the Global Functions sections may be overwritten.            │
└──────────────────────────────────────────────────────────────────────────────┘

Package "module" declares type aliases for the commonly used types declared in
the packages contained in this module.  It also provides constructors for each
commonly used class that is exported by the module.  Each constructor delegates
the actual construction process to its corresponding concrete class declared in
the corresponding package contained within this module.

For detailed documentation on this entire module refer to the wiki:
  - https://github.com/craterdog/go-example-package/wiki
*/
package module

import (
	exa "github.com/craterdog/go-example-package/v2/example"
)

// TYPE ALIASES

// Example

type (
	Identifier = exa.Identifier
	Cardinal   = exa.Cardinal
	Ordinal    = exa.Ordinal
	Rank       = exa.Rank
	Slot       = exa.Slot
	Units      = exa.Units
	Regexp     = exa.Regexp
)

const (
	LesserRank  = exa.LesserRank
	EqualRank   = exa.EqualRank
	GreaterRank = exa.GreaterRank
	Degrees     = exa.Degrees
	Radians     = exa.Radians
	Gradians    = exa.Gradians
)

type (
	RankingFunction[V any] = exa.RankingFunction[V]
)

type (
	AngleClassLike                            = exa.AngleClassLike
	ArrayClassLike[V any]                     = exa.ArrayClassLike[V]
	AssociationClassLike[K comparable, V any] = exa.AssociationClassLike[K, V]
	CatalogClassLike[V any]                   = exa.CatalogClassLike[V]
	IteratorClassLike[V any]                  = exa.IteratorClassLike[V]
)

type (
	AngleLike                            = exa.AngleLike
	ArrayLike[V any]                     = exa.ArrayLike[V]
	AssociationLike[K comparable, V any] = exa.AssociationLike[K, V]
	CatalogLike[V any]                   = exa.CatalogLike[V]
	IteratorLike[V any]                  = exa.IteratorLike[V]
)

type (
	Accessible[V any]                = exa.Accessible[V]
	Angular                          = exa.Angular
	Associative[K comparable, V any] = exa.Associative[K, V]
	Continuous                       = exa.Continuous
	Sequential[V any]                = exa.Sequential[V]
	Updatable[V any]                 = exa.Updatable[V]
)

// CLASS ACCESSORS

// Example

func AngleClass() AngleClassLike {
	return exa.AngleClass()
}

func Angle(
	radians float64,
) AngleLike {
	return AngleClass().Angle(
		radians,
	)
}

func AngleFromSource(
	source string,
) AngleLike {
	return AngleClass().AngleFromSource(
		source,
	)
}

func ArrayClass[V any]() ArrayClassLike[V] {
	return exa.ArrayClass[V]()
}

func Array[V any](
	array []V,
) ArrayLike[V] {
	return ArrayClass[V]().Array(
		array,
	)
}

func ArrayWithSize[V any](
	size exa.Cardinal,
) ArrayLike[V] {
	return ArrayClass[V]().ArrayWithSize(
		size,
	)
}

func ArrayFromSequence[V any](
	values exa.Sequential[V],
) ArrayLike[V] {
	return ArrayClass[V]().ArrayFromSequence(
		values,
	)
}

func AssociationClass[K comparable, V any]() AssociationClassLike[K, V] {
	return exa.AssociationClass[K, V]()
}

func Association[K comparable, V any](
	key K,
	value V,
) AssociationLike[K, V] {
	return AssociationClass[K, V]().Association(
		key,
		value,
	)
}

func CatalogClass[V any]() CatalogClassLike[V] {
	return exa.CatalogClass[V]()
}

func Catalog[V any]() CatalogLike[V] {
	return CatalogClass[V]().Catalog()
}

func CatalogFromArray[V any](
	associations []exa.AssociationLike[exa.Identifier, V],
) CatalogLike[V] {
	return CatalogClass[V]().CatalogFromArray(
		associations,
	)
}

func CatalogFromMap[V any](
	associations map[Identifier]V,
) CatalogLike[V] {
	return CatalogClass[V]().CatalogFromMap(
		associations,
	)
}

func CatalogFromSequence[V any](
	associations exa.Sequential[exa.AssociationLike[exa.Identifier, V]],
) CatalogLike[V] {
	return CatalogClass[V]().CatalogFromSequence(
		associations,
	)
}

func IteratorClass[V any]() IteratorClassLike[V] {
	return exa.IteratorClass[V]()
}

func Iterator[V any](
	array []V,
) IteratorLike[V] {
	return IteratorClass[V]().Iterator(
		array,
	)
}

// GLOBAL FUNCTIONS
