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

func AttributeMethod() AttributeMethodClassLike {
	return attributeMethodReference()
}

// Constructor Methods

func (c *attributeMethodClass_) Make(
	any_ any,
) AttributeMethodLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &attributeMethod_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *attributeMethod_) GetClass() AttributeMethodClassLike {
	return attributeMethodReference()
}

// Attribute Methods

func (v *attributeMethod_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type attributeMethod_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type attributeMethodClass_ struct {
	// Declare the class constants.
}

// Class Reference

func attributeMethodReference() *attributeMethodClass_ {
	return attributeMethodReference_
}

var attributeMethodReference_ = &attributeMethodClass_{
	// Initialize the class constants.
}
