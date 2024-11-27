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
│                 This class file was automatically generated.                 │
│   Updates to any section other than the Private Methods may be overwritten.  │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func Map() MapClassLike {
	return mapReference()
}

// Constructor Methods

func (c *mapClass_) Make(
	name string,
) MapLike {
	if uti.IsUndefined(name) {
		panic("The \"name\" attribute is required by this class.")
	}
	var instance = &map_{
		// Initialize the instance attributes.
		name_: name,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *map_) GetClass() MapClassLike {
	return mapReference()
}

// Attribute Methods

func (v *map_) GetName() string {
	return v.name_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type map_ struct {
	// Declare the instance attributes.
	name_ string
}

// Class Structure

type mapClass_ struct {
	// Declare the class constants.
}

// Class Reference

func mapReference() *mapClass_ {
	return mapReference_
}

var mapReference_ = &mapClass_{
	// Initialize the class constants.
}
