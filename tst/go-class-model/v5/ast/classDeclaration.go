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

func ClassDeclarationClass() ClassDeclarationClassLike {
	return classDeclarationClassReference()
}

// Constructor Methods

func (c *classDeclarationClass_) ClassDeclaration(
	declaration DeclarationLike,
	classMethods ClassMethodsLike,
) ClassDeclarationLike {
	if uti.IsUndefined(declaration) {
		panic("The \"declaration\" attribute is required by this class.")
	}
	if uti.IsUndefined(classMethods) {
		panic("The \"classMethods\" attribute is required by this class.")
	}
	var instance = &classDeclaration_{
		// Initialize the instance attributes.
		declaration_:  declaration,
		classMethods_: classMethods,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *classDeclaration_) GetClass() ClassDeclarationClassLike {
	return classDeclarationClassReference()
}

// Attribute Methods

func (v *classDeclaration_) GetDeclaration() DeclarationLike {
	return v.declaration_
}

func (v *classDeclaration_) GetClassMethods() ClassMethodsLike {
	return v.classMethods_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type classDeclaration_ struct {
	// Declare the instance attributes.
	declaration_  DeclarationLike
	classMethods_ ClassMethodsLike
}

// Class Structure

type classDeclarationClass_ struct {
	// Declare the class constants.
}

// Class Reference

func classDeclarationClassReference() *classDeclarationClass_ {
	return classDeclarationClassReference_
}

var classDeclarationClassReference_ = &classDeclarationClass_{
	// Initialize the class constants.
}
