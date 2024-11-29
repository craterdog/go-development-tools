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
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func Suffix() SuffixClassLike {
	return suffixReference()
}

// Constructor Methods

func (c *suffixClass_) Make(
	name string,
) SuffixLike {
	if uti.IsUndefined(name) {
		panic("The \"name\" attribute is required by this class.")
	}
	var instance = &suffix_{
		// Initialize the instance attributes.
		name_: name,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *suffix_) GetClass() SuffixClassLike {
	return suffixReference()
}

// Attribute Methods

func (v *suffix_) GetName() string {
	return v.name_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type suffix_ struct {
	// Declare the instance attributes.
	name_ string
}

// Class Structure

type suffixClass_ struct {
	// Declare the class constants.
}

// Class Reference

func suffixReference() *suffixClass_ {
	return suffixReference_
}

var suffixReference_ = &suffixClass_{
	// Initialize the class constants.
}
