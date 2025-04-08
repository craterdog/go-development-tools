/*
................................................................................
.    Copyright (c) 2009-2025 Crater Dog Technologies.  All Rights Reserved.    .
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
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func KeyClass() KeyClassLike {
	return keyClass()
}

// Constructor Methods

func (c *keyClass_) Key(
	primitive PrimitiveLike,
) KeyLike {
	if uti.IsUndefined(primitive) {
		panic("The \"primitive\" attribute is required by this class.")
	}
	var instance = &key_{
		// Initialize the instance attributes.
		primitive_: primitive,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *key_) GetClass() KeyClassLike {
	return keyClass()
}

// Attribute Methods

func (v *key_) GetPrimitive() PrimitiveLike {
	return v.primitive_
}

// PROTECTED INTERFACE

// Instance Structure

type key_ struct {
	// Declare the instance attributes.
	primitive_ PrimitiveLike
}

// Class Structure

type keyClass_ struct {
	// Declare the class constants.
}

// Class Reference

func keyClass() *keyClass_ {
	return keyClassReference_
}

var keyClassReference_ = &keyClass_{
	// Initialize the class constants.
}
