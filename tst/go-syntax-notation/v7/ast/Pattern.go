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

func PatternClass() PatternClassLike {
	return patternClass()
}

// Constructor Methods

func (c *patternClass_) Pattern(
	alternatives AlternativesLike,
	optionalNote string,
) PatternLike {
	if uti.IsUndefined(alternatives) {
		panic("The \"alternatives\" attribute is required by this class.")
	}
	var instance = &pattern_{
		// Initialize the instance attributes.
		alternatives_: alternatives,
		optionalNote_: optionalNote,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *pattern_) GetClass() PatternClassLike {
	return patternClass()
}

// Attribute Methods

func (v *pattern_) GetAlternatives() AlternativesLike {
	return v.alternatives_
}

func (v *pattern_) GetOptionalNote() string {
	return v.optionalNote_
}

// PROTECTED INTERFACE

// Instance Structure

type pattern_ struct {
	// Declare the instance attributes.
	alternatives_ AlternativesLike
	optionalNote_ string
}

// Class Structure

type patternClass_ struct {
	// Declare the class constants.
}

// Class Reference

func patternClass() *patternClass_ {
	return patternClassReference_
}

var patternClassReference_ = &patternClass_{
	// Initialize the class constants.
}
