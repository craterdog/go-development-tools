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

func AdditionalValue() AdditionalValueClassLike {
	return additionalValueReference()
}

// Constructor Methods

func (c *additionalValueClass_) Make(
	name string,
) AdditionalValueLike {
	if uti.IsUndefined(name) {
		panic("The \"name\" attribute is required by this class.")
	}
	var instance = &additionalValue_{
		// Initialize the instance attributes.
		name_: name,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *additionalValue_) GetClass() AdditionalValueClassLike {
	return additionalValueReference()
}

// Attribute Methods

func (v *additionalValue_) GetName() string {
	return v.name_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type additionalValue_ struct {
	// Declare the instance attributes.
	name_ string
}

// Class Structure

type additionalValueClass_ struct {
	// Declare the class constants.
}

// Class Reference

func additionalValueReference() *additionalValueClass_ {
	return additionalValueReference_
}

var additionalValueReference_ = &additionalValueClass_{
	// Initialize the class constants.
}
