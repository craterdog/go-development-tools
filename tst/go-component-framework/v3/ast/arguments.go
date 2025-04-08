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
	col "github.com/craterdog/go-collection-framework/v5/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func ArgumentsClass() ArgumentsClassLike {
	return argumentsClass()
}

// Constructor Methods

func (c *argumentsClass_) Arguments(
	argument ArgumentLike,
	additionalArguments col.Sequential[AdditionalArgumentLike],
) ArgumentsLike {
	if uti.IsUndefined(argument) {
		panic("The \"argument\" attribute is required by this class.")
	}
	if uti.IsUndefined(additionalArguments) {
		panic("The \"additionalArguments\" attribute is required by this class.")
	}
	var instance = &arguments_{
		// Initialize the instance attributes.
		argument_:            argument,
		additionalArguments_: additionalArguments,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *arguments_) GetClass() ArgumentsClassLike {
	return argumentsClass()
}

// Attribute Methods

func (v *arguments_) GetArgument() ArgumentLike {
	return v.argument_
}

func (v *arguments_) GetAdditionalArguments() col.Sequential[AdditionalArgumentLike] {
	return v.additionalArguments_
}

// PROTECTED INTERFACE

// Instance Structure

type arguments_ struct {
	// Declare the instance attributes.
	argument_            ArgumentLike
	additionalArguments_ col.Sequential[AdditionalArgumentLike]
}

// Class Structure

type argumentsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func argumentsClass() *argumentsClass_ {
	return argumentsClassReference_
}

var argumentsClassReference_ = &argumentsClass_{
	// Initialize the class constants.
}
