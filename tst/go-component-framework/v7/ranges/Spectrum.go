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

package ranges

import (
	fmt "fmt"
	age "github.com/craterdog/go-component-framework/v7/agents"
	ele "github.com/craterdog/go-component-framework/v7/elements"
	str "github.com/craterdog/go-component-framework/v7/strings"
	syn "sync"
)

// CLASS INTERFACE

// Access Function

func SpectrumClass[V ele.Continuous]() SpectrumClassLike[V] {
	return spectrumClass[V]()
}

// Constructor Methods

func (c *spectrumClass_[V]) Spectrum(
	left Bracket,
	minimum V,
	maximum V,
	right Bracket,
) SpectrumLike[V] {
	var instance = &spectrum_[V]{
		// Initialize the instance attributes.
		left_:    left,
		minimum_: minimum,
		maximum_: maximum,
		right_:   right,
	}
	instance.validateSpectrum()
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *spectrum_[V]) GetClass() SpectrumClassLike[V] {
	return spectrumClass[V]()
}

// Bounded[V] Methods

func (v *spectrum_[V]) GetLeft() Bracket {
	return v.left_
}

func (v *spectrum_[V]) SetLeft(left Bracket) {
	v.left_ = left
	v.validateSpectrum()
}

func (v *spectrum_[V]) GetMinimum() V {
	return v.minimum_
}

func (v *spectrum_[V]) SetMinimum(minimum V) {
	v.minimum_ = minimum
	v.validateSpectrum()
}

func (v *spectrum_[V]) GetMaximum() V {
	return v.maximum_
}

func (v *spectrum_[V]) SetMaximum(maximum V) {
	v.maximum_ = maximum
	v.validateSpectrum()
}

func (v *spectrum_[V]) GetRight() Bracket {
	return v.right_
}

func (v *spectrum_[V]) SetRight(right Bracket) {
	v.right_ = right
	v.validateSpectrum()
}

// Searchable[V] Methods

func (v *spectrum_[V]) ContainsValue(
	value V,
) bool {
	if v.minimum_.IsDefined() && value.AsFloat() < v.minimum_.AsFloat() {
		return false
	}
	if v.maximum_.IsDefined() && value.AsFloat() > v.maximum_.AsFloat() {
		return false
	}
	return true
}

func (v *spectrum_[V]) ContainsAny(
	values str.Sequential[V],
) bool {
	var iterator = values.GetIterator()
	for iterator.HasNext() {
		var value = iterator.GetNext()
		if v.ContainsValue(value) {
			// This spectrum contains at least one of the values.
			return true
		}
	}
	// This spectrum does not contain any of the values.
	return false
}

func (v *spectrum_[V]) ContainsAll(
	values str.Sequential[V],
) bool {
	var iterator = values.GetIterator()
	for iterator.HasNext() {
		var value = iterator.GetNext()
		if !v.ContainsValue(value) {
			// This spectrum is missing at least one of the values.
			return false
		}
	}
	// This spectrum does contains all of the values.
	return true
}

// PROTECTED INTERFACE

func (v *spectrum_[V]) String() string {
	var string_ string
	switch v.left_ {
	case Inclusive:
		string_ += "["
	case Exclusive:
		string_ += "("
	}
	if v.minimum_.IsDefined() {
		string_ += v.minimum_.AsString()
	}
	string_ += ".."
	if v.maximum_.IsDefined() {
		string_ += v.maximum_.AsString()
	}
	switch v.right_ {
	case Inclusive:
		string_ += "]"
	case Exclusive:
		string_ += ")"
	}
	return string_
}

// Private Methods

// This method ensures that the endpoints are valid.
func (v *spectrum_[V]) validateSpectrum() {
	// Validate the left bracket.
	switch v.left_ {
	case Inclusive:
	case Exclusive:
	default:
		var message = fmt.Sprintf(
			"Received an invalid left bracket for a spectrum: %v",
			v.left_,
		)
		panic(message)
	}

	// Validate the right bracket.
	switch v.right_ {
	case Inclusive:
	case Exclusive:
	default:
		var message = fmt.Sprintf(
			"Received an invalid right bracket for a spectrum: %v",
			v.right_,
		)
		panic(message)
	}

	// Validate the endpoints.
	if v.minimum_.IsDefined() && v.maximum_.IsDefined() {
		var collator = age.CollatorClass[V]().Collator()
		if collator.RankValues(v.minimum_, v.maximum_) != age.LesserRank {
			var message = fmt.Sprintf(
				"The minimum %v in a spectrum must be less than the maximum %v.",
				v.minimum_,
				v.maximum_,
			)
			panic(message)
		}
		var size = v.maximum_.AsFloat() - v.minimum_.AsFloat()
		if size <= 0 {
			var message = fmt.Sprintf(
				"The size of a spectrum must be greater than zero: %v.",
				size,
			)
			panic(message)
		}
	}
}

// Instance Structure

type spectrum_[V ele.Continuous] struct {
	// Declare the instance attributes.
	left_    Bracket
	minimum_ V
	maximum_ V
	right_   Bracket
}

// Class Structure

type spectrumClass_[V ele.Continuous] struct {
	// Declare the class constants.
}

// Class Reference

var spectrumMap_ = map[string]any{}
var spectrumMutex_ syn.Mutex

func spectrumClass[V ele.Continuous]() *spectrumClass_[V] {
	// Generate the name of the bound class type.
	var class *spectrumClass_[V]
	var name = fmt.Sprintf("%T", class)

	// Check for an existing bound class type.
	spectrumMutex_.Lock()
	var value = spectrumMap_[name]
	switch actual := value.(type) {
	case *spectrumClass_[V]:
		// This bound class type already exists.
		class = actual
	default:
		// Add a new bound class type.
		class = &spectrumClass_[V]{
			// Initialize the class constants.
		}
		spectrumMap_[name] = class
	}
	spectrumMutex_.Unlock()

	// Return a reference to the bound class type.
	return class
}
