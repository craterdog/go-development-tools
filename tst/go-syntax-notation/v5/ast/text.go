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

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func TextClass() TextClassLike {
	return textClassReference()
}

// Constructor Methods

func (c *textClass_) Text(
	any_ any,
) TextLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &text_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *text_) GetClass() TextClassLike {
	return textClassReference()
}

// Attribute Methods

func (v *text_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type text_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type textClass_ struct {
	// Declare the class constants.
}

// Class Reference

func textClassReference() *textClass_ {
	return textClassReference_
}

var textClassReference_ = &textClass_{
	// Initialize the class constants.
}
