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

func ThrowClauseClass() ThrowClauseClassLike {
	return throwClauseClass()
}

// Constructor Methods

func (c *throwClauseClass_) ThrowClause(
	exception ExceptionLike,
) ThrowClauseLike {
	if uti.IsUndefined(exception) {
		panic("The \"exception\" attribute is required by this class.")
	}
	var instance = &throwClause_{
		// Initialize the instance attributes.
		exception_: exception,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *throwClause_) GetClass() ThrowClauseClassLike {
	return throwClauseClass()
}

// Attribute Methods

func (v *throwClause_) GetException() ExceptionLike {
	return v.exception_
}

// PROTECTED INTERFACE

// Instance Structure

type throwClause_ struct {
	// Declare the instance attributes.
	exception_ ExceptionLike
}

// Class Structure

type throwClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func throwClauseClass() *throwClauseClass_ {
	return throwClauseClassReference_
}

var throwClauseClassReference_ = &throwClauseClass_{
	// Initialize the class constants.
}
