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

func AdditionalValueClass() AdditionalValueClassLike {
	return additionalValueClass()
}

// Constructor Methods

func (c *additionalValueClass_) AdditionalValue(
	value ValueLike,
) AdditionalValueLike {
	if uti.IsUndefined(value) {
		panic("The \"value\" attribute is required by this class.")
	}
	var instance = &additionalValue_{
		// Initialize the instance attributes.
		value_: value,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *additionalValue_) GetClass() AdditionalValueClassLike {
	return additionalValueClass()
}

// Attribute Methods

func (v *additionalValue_) GetValue() ValueLike {
	return v.value_
}

// PROTECTED INTERFACE

// Instance Structure

type additionalValue_ struct {
	// Declare the instance attributes.
	value_ ValueLike
}

// Class Structure

type additionalValueClass_ struct {
	// Declare the class constants.
}

// Class Reference

func additionalValueClass() *additionalValueClass_ {
	return additionalValueClassReference_
}

var additionalValueClassReference_ = &additionalValueClass_{
	// Initialize the class constants.
}
