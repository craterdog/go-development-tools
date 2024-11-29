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

func Parameter() ParameterClassLike {
	return parameterReference()
}

// Constructor Methods

func (c *parameterClass_) Make(
	name string,
	abstraction AbstractionLike,
) ParameterLike {
	if uti.IsUndefined(name) {
		panic("The \"name\" attribute is required by this class.")
	}
	if uti.IsUndefined(abstraction) {
		panic("The \"abstraction\" attribute is required by this class.")
	}
	var instance = &parameter_{
		// Initialize the instance attributes.
		name_:        name,
		abstraction_: abstraction,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *parameter_) GetClass() ParameterClassLike {
	return parameterReference()
}

// Attribute Methods

func (v *parameter_) GetName() string {
	return v.name_
}

func (v *parameter_) GetAbstraction() AbstractionLike {
	return v.abstraction_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type parameter_ struct {
	// Declare the instance attributes.
	name_        string
	abstraction_ AbstractionLike
}

// Class Structure

type parameterClass_ struct {
	// Declare the class constants.
}

// Class Reference

func parameterReference() *parameterClass_ {
	return parameterReference_
}

var parameterReference_ = &parameterClass_{
	// Initialize the class constants.
}
