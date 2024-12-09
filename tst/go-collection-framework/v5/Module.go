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
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│              This "Module.go" file was automatically generated.              │
│      Updates to any part of this file—other than the Module Description      │
│             and the Global Functions sections may be overwritten.            │
└──────────────────────────────────────────────────────────────────────────────┘

Package "module" declares type aliases for the commonly used types declared in
the packages contained in this module.  It also provides a default constructor
for each commonly used class that is exported by the module.  Each constructor
delegates the actual construction process to its corresponding concrete class
declared in the corresponding package contained within this module.

For detailed documentation on this entire module refer to the wiki:
  - https://github.com/craterdog/go-collection-framework//wiki
*/
package module

import (
	age "github.com/craterdog/go-collection-framework/v5/agent"
	col "github.com/craterdog/go-collection-framework/v5/collection"
)

// TYPE ALIASES

// Agent

type (
	Rank = age.Rank
)

const (
	LesserRank  = age.LesserRank
	EqualRank   = age.EqualRank
	GreaterRank = age.GreaterRank
)

// Collection

type (
	Index = col.Index
	Size  = col.Size
	Slot  = col.Slot
)

type (
	Synchronized = col.Synchronized
)

// DEFAULT CONSTRUCTORS

// Agent

func Collator[V any]() age.CollatorLike[V] {
	return age.CollatorClass[V]().Make()
}

func Iterator[V any](
	values []V,
) age.IteratorLike[V] {
	return age.IteratorClass[V]().Make(
		values,
	)
}

func Sorter[V any]() age.SorterLike[V] {
	return age.SorterClass[V]().Make()
}

// Collection

func Array[V any](
	size Size,
) col.ArrayLike[V] {
	return col.ArrayClass[V]().Make(
		size,
	)
}

func Association[K comparable, V any](
	key K,
	value V,
) col.AssociationLike[K, V] {
	return col.AssociationClass[K, V]().Make(
		key,
		value,
	)
}

func Catalog[K comparable, V any]() col.CatalogLike[K, V] {
	return col.CatalogClass[K, V]().Make()
}

func List[V any]() col.ListLike[V] {
	return col.ListClass[V]().Make()
}

func Map[K comparable, V any]() col.MapLike[K, V] {
	return col.MapClass[K, V]().Make()
}

func Queue[V any]() col.QueueLike[V] {
	return col.QueueClass[V]().Make()
}

func Set[V any]() col.SetLike[V] {
	return col.SetClass[V]().Make()
}

func Stack[V any]() col.StackLike[V] {
	return col.StackClass[V]().Make()
}

// GLOBAL FUNCTIONS
