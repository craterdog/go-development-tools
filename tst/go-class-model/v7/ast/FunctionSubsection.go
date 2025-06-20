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
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func FunctionSubsectionClass() FunctionSubsectionClassLike {
	return functionSubsectionClass()
}

// Constructor Methods

func (c *functionSubsectionClass_) FunctionSubsection(
	delimiter string,
	functionMethods fra.ListLike[FunctionMethodLike],
) FunctionSubsectionLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(functionMethods) {
		panic("The \"functionMethods\" attribute is required by this class.")
	}
	var instance = &functionSubsection_{
		// Initialize the instance attributes.
		delimiter_:       delimiter,
		functionMethods_: functionMethods,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *functionSubsection_) GetClass() FunctionSubsectionClassLike {
	return functionSubsectionClass()
}

// Attribute Methods

func (v *functionSubsection_) GetDelimiter() string {
	return v.delimiter_
}

func (v *functionSubsection_) GetFunctionMethods() fra.ListLike[FunctionMethodLike] {
	return v.functionMethods_
}

// PROTECTED INTERFACE

// Instance Structure

type functionSubsection_ struct {
	// Declare the instance attributes.
	delimiter_       string
	functionMethods_ fra.ListLike[FunctionMethodLike]
}

// Class Structure

type functionSubsectionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func functionSubsectionClass() *functionSubsectionClass_ {
	return functionSubsectionClassReference_
}

var functionSubsectionClassReference_ = &functionSubsectionClass_{
	// Initialize the class constants.
}
