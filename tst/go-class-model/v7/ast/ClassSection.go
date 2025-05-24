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
	col "github.com/craterdog/go-collection-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func ClassSectionClass() ClassSectionClassLike {
	return classSectionClass()
}

// Constructor Methods

func (c *classSectionClass_) ClassSection(
	delimiter string,
	classDeclarations col.ListLike[ClassDeclarationLike],
) ClassSectionLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(classDeclarations) {
		panic("The \"classDeclarations\" attribute is required by this class.")
	}
	var instance = &classSection_{
		// Initialize the instance attributes.
		delimiter_:         delimiter,
		classDeclarations_: classDeclarations,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *classSection_) GetClass() ClassSectionClassLike {
	return classSectionClass()
}

// Attribute Methods

func (v *classSection_) GetDelimiter() string {
	return v.delimiter_
}

func (v *classSection_) GetClassDeclarations() col.ListLike[ClassDeclarationLike] {
	return v.classDeclarations_
}

// PROTECTED INTERFACE

// Instance Structure

type classSection_ struct {
	// Declare the instance attributes.
	delimiter_         string
	classDeclarations_ col.ListLike[ClassDeclarationLike]
}

// Class Structure

type classSectionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func classSectionClass() *classSectionClass_ {
	return classSectionClassReference_
}

var classSectionClassReference_ = &classSectionClass_{
	// Initialize the class constants.
}
