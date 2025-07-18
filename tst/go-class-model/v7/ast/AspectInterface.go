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
│              This class file was automatically generated using:              │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func AspectInterfaceClass() AspectInterfaceClassLike {
	return aspectInterfaceClass()
}

// Constructor Methods

func (c *aspectInterfaceClass_) AspectInterface(
	abstraction AbstractionLike,
) AspectInterfaceLike {
	if uti.IsUndefined(abstraction) {
		panic("The \"abstraction\" attribute is required by this class.")
	}
	var instance = &aspectInterface_{
		// Initialize the instance attributes.
		abstraction_: abstraction,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *aspectInterface_) GetClass() AspectInterfaceClassLike {
	return aspectInterfaceClass()
}

// Attribute Methods

func (v *aspectInterface_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

// PROTECTED INTERFACE

// Instance Structure

type aspectInterface_ struct {
	// Declare the instance attributes.
	abstraction_ AbstractionLike
}

// Class Structure

type aspectInterfaceClass_ struct {
	// Declare the class constants.
}

// Class Reference

func aspectInterfaceClass() *aspectInterfaceClass_ {
	return aspectInterfaceClassReference_
}

var aspectInterfaceClassReference_ = &aspectInterfaceClass_{
	// Initialize the class constants.
}
