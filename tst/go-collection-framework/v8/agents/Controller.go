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

package agents

import (
	uti "github.com/craterdog/go-missing-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func ControllerClass() ControllerClassLike {
	return controllerClass()
}

// Constructor Methods

func (c *controllerClass_) Controller(
	events []Event,
	transitions map[State]Transitions,
	initialState State,
) ControllerLike {
	if uti.IsUndefined(events) {
		panic("The \"events\" attribute is required by this class.")
	}
	if uti.IsUndefined(transitions) {
		panic("The \"transitions\" attribute is required by this class.")
	}
	if uti.IsUndefined(initialState) {
		panic("The \"initialState\" attribute is required by this class.")
	}
	var instance = &controller_{
		// Initialize the instance attributes.
		events_:      events,
		transitions_: transitions,
	}
	return instance
}

// Constant Methods

func (c *controllerClass_) Invalid() State {
	return c.invalid_
}

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *controller_) GetClass() ControllerClassLike {
	return controllerClass()
}

func (v *controller_) ProcessEvent(
	event Event,
) State {
	var result_ State
	// TBD - Add the method implementation.
	return result_
}

// Attribute Methods

func (v *controller_) GetState() State {
	return v.state_
}

func (v *controller_) SetState(
	state State,
) {
	if uti.IsUndefined(state) {
		panic("The \"state\" attribute is required by this class.")
	}
	v.state_ = state
}

func (v *controller_) GetEvents() []Event {
	return v.events_
}

func (v *controller_) GetTransitions() map[State]Transitions {
	return v.transitions_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type controller_ struct {
	// Declare the instance attributes.
	state_       State
	events_      []Event
	transitions_ map[State]Transitions
}

// Class Structure

type controllerClass_ struct {
	// Declare the class constants.
	invalid_ State
}

// Class Reference

func controllerClass() *controllerClass_ {
	return controllerClassReference_
}

var controllerClassReference_ = &controllerClass_{
	// Initialize the class constants.
	// invalid_: constantValue,
}
