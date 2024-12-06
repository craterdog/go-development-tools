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

func AdditionalConstraintClass() AdditionalConstraintClassLike {
	return additionalConstraintClassReference()
}

// Constructor Methods

func (c *additionalConstraintClass_) Make(
	constraint ConstraintLike,
) AdditionalConstraintLike {
	if uti.IsUndefined(constraint) {
		panic("The \"constraint\" attribute is required by this class.")
	}
	var instance = &additionalConstraint_{
		// Initialize the instance attributes.
		constraint_: constraint,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *additionalConstraint_) GetClass() AdditionalConstraintClassLike {
	return additionalConstraintClassReference()
}

// Attribute Methods

func (v *additionalConstraint_) GetConstraint() ConstraintLike {
	return v.constraint_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type additionalConstraint_ struct {
	// Declare the instance attributes.
	constraint_ ConstraintLike
}

// Class Structure

type additionalConstraintClass_ struct {
	// Declare the class constants.
}

// Class Reference

func additionalConstraintClassReference() *additionalConstraintClass_ {
	return additionalConstraintClassReference_
}

var additionalConstraintClassReference_ = &additionalConstraintClass_{
	// Initialize the class constants.
}
