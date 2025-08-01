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

func InterfaceDeclarationsClass() InterfaceDeclarationsClassLike {
	return interfaceDeclarationsClass()
}

// Constructor Methods

func (c *interfaceDeclarationsClass_) InterfaceDeclarations(
	classSection ClassSectionLike,
	instanceSection InstanceSectionLike,
	aspectSection AspectSectionLike,
) InterfaceDeclarationsLike {
	if uti.IsUndefined(classSection) {
		panic("The \"classSection\" attribute is required by this class.")
	}
	if uti.IsUndefined(instanceSection) {
		panic("The \"instanceSection\" attribute is required by this class.")
	}
	if uti.IsUndefined(aspectSection) {
		panic("The \"aspectSection\" attribute is required by this class.")
	}
	var instance = &interfaceDeclarations_{
		// Initialize the instance attributes.
		classSection_:    classSection,
		instanceSection_: instanceSection,
		aspectSection_:   aspectSection,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *interfaceDeclarations_) GetClass() InterfaceDeclarationsClassLike {
	return interfaceDeclarationsClass()
}

// Attribute Methods

func (v *interfaceDeclarations_) GetClassSection() ClassSectionLike {
	return v.classSection_
}

func (v *interfaceDeclarations_) GetInstanceSection() InstanceSectionLike {
	return v.instanceSection_
}

func (v *interfaceDeclarations_) GetAspectSection() AspectSectionLike {
	return v.aspectSection_
}

// PROTECTED INTERFACE

// Instance Structure

type interfaceDeclarations_ struct {
	// Declare the instance attributes.
	classSection_    ClassSectionLike
	instanceSection_ InstanceSectionLike
	aspectSection_   AspectSectionLike
}

// Class Structure

type interfaceDeclarationsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func interfaceDeclarationsClass() *interfaceDeclarationsClass_ {
	return interfaceDeclarationsClassReference_
}

var interfaceDeclarationsClassReference_ = &interfaceDeclarationsClass_{
	// Initialize the class constants.
}
