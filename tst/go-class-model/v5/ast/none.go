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

func None() NoneClassLike {
	return noneReference()
}

// Constructor Methods

func (c *noneClass_) Make(
	newline string,
) NoneLike {
	if uti.IsUndefined(newline) {
		panic("The \"newline\" attribute is required by this class.")
	}
	var instance = &none_{
		// Initialize the instance attributes.
		newline_: newline,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *none_) GetClass() NoneClassLike {
	return noneReference()
}

// Attribute Methods

func (v *none_) GetNewline() string {
	return v.newline_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type none_ struct {
	// Declare the instance attributes.
	newline_ string
}

// Class Structure

type noneClass_ struct {
	// Declare the class constants.
}

// Class Reference

func noneReference() *noneClass_ {
	return noneReference_
}

var noneReference_ = &noneClass_{
	// Initialize the class constants.
}
