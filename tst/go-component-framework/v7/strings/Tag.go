/*
................................................................................
.    Copyright (c) 2009-2025 Crater Dog Technologies™.  All Rights Reserved.   .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package strings

import (
	bin "encoding/binary"
	fmt "fmt"
	age "github.com/craterdog/go-component-framework/v7/agents"
	uti "github.com/craterdog/go-missing-utilities/v7"
	reg "regexp"
	sli "slices"
)

// CLASS INTERFACE

// Access Function

func TagClass() TagClassLike {
	return tagClass()
}

// Constructor Methods

func (c *tagClass_) Tag(
	bytes []byte,
) TagLike {
	c.validateSize(uti.Cardinal(len(bytes)))
	return tag_(bytes)
}

func (c *tagClass_) TagWithSize(
	size uti.Cardinal,
) TagLike {
	c.validateSize(size)
	var generator = age.GeneratorClass().Generator()
	var bytes = generator.RandomBytes(size)
	return tag_(bytes)
}

func (c *tagClass_) TagFromSequence(
	sequence Sequential[byte],
) TagLike {
	c.validateSize(sequence.GetSize())
	return tag_(sequence.AsArray())
}

func (c *tagClass_) TagFromString(
	source string,
) TagLike {
	var matches = c.matcher_.FindStringSubmatch(source)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the tag constructor method: %s",
			source,
		)
		panic(message)
	}
	var base32 = matches[1] // Strip off the leading "#".
	var encoder = age.EncoderClass().Encoder()
	var bytes = encoder.Base32Decode(base32)
	return tag_(bytes)
}

// Constant Methods

// Function Methods

func (c *tagClass_) Concatenate(
	first TagLike,
	second TagLike,
) TagLike {
	var firstBytes = first.AsArray()
	var secondBytes = second.AsArray()
	var allBytes = make(
		[]byte,
		len(firstBytes)+len(secondBytes),
	)
	copy(allBytes, firstBytes)
	copy(allBytes[len(firstBytes):], secondBytes)
	return c.Tag(allBytes)
}

// INSTANCE INTERFACE

// Principal Methods

func (v tag_) GetClass() TagClassLike {
	return tagClass()
}

func (v tag_) AsIntrinsic() []byte {
	return []byte(v)
}

func (v tag_) AsString() string {
	var encoder = age.EncoderClass().Encoder()
	return "#" + encoder.Base32Encode(v)
}

func (v tag_) GetHash() uint64 {
	return bin.BigEndian.Uint64(v)
}

// Attribute Methods

// Searchable[byte] Methods

func (v tag_) ContainsValue(
	value byte,
) bool {
	return sli.Index(v, value) > -1
}

func (v tag_) ContainsAny(
	values Sequential[byte],
) bool {
	var iterator = values.GetIterator()
	for iterator.HasNext() {
		var value = iterator.GetNext()
		if v.ContainsValue(value) {
			// This set contains at least one of the values.
			return true
		}
	}
	// This set does not contain any of the values.
	return false
}

func (v tag_) ContainsAll(
	values Sequential[byte],
) bool {
	var iterator = values.GetIterator()
	for iterator.HasNext() {
		var value = iterator.GetNext()
		if !v.ContainsValue(value) {
			// This set is missing at least one of the values.
			return false
		}
	}
	// This set does contains all of the values.
	return true
}

// Sequential[byte] Methods

func (v tag_) IsEmpty() bool {
	return len(v) == 0
}

func (v tag_) GetSize() uti.Cardinal {
	return uti.Cardinal(len(v))
}

func (v tag_) AsArray() []byte {
	return uti.CopyArray(v)
}

func (v tag_) GetIterator() age.IteratorLike[byte] {
	var array = uti.CopyArray(v)
	var class = age.IteratorClass[byte]()
	var iterator = class.Iterator(array)
	return iterator
}

// Accessible[byte] Methods

func (v tag_) GetValue(
	index uti.Index,
) byte {
	var size = v.GetSize()
	var goIndex = uti.RelativeToZeroBased(index, size)
	return v[goIndex]
}

func (v tag_) GetValues(
	first uti.Index,
	last uti.Index,
) Sequential[byte] {
	var size = v.GetSize()
	var goFirst = uti.RelativeToZeroBased(first, size)
	var goLast = uti.RelativeToZeroBased(last, size)
	return tag_(v[goFirst : goLast+1])
}

func (v tag_) GetIndex(
	value byte,
) uti.Index {
	var index uti.Index
	var iterator = v.GetIterator()
	for iterator.HasNext() {
		index++
		var candidate = iterator.GetNext()
		if candidate == value {
			// Found the value.
			return index
		}
	}
	// The value was not found.
	return 0
}

// PROTECTED INTERFACE

func (v tag_) String() string {
	return v.AsString()
}

// Private Methods

func (c *tagClass_) validateSize(
	size uti.Cardinal,
) {
	if size < 8 {
		var message = fmt.Sprintf(
			"A tag must be at least eight bytes long: %v",
			size,
		)
		panic(message)
	}
}

// NOTE:
// These private constants are used to define the private regular expression
// matcher that is used to match legal string patterns for this intrinsic type.
// Unfortunately there is no way to make them private to this class since they
// must be TRUE Go constants to be used in this way.  We append an underscore to
// each name to lessen the chance of a name collision with other private Go
// class constants in this package.
const (
	base32_ = base10_ + "|[A-DF-HJ-NP-TV-Z]"
)

// Instance Structure

type tag_ []byte

// Class Structure

type tagClass_ struct {
	// Declare the class constants.
	matcher_ *reg.Regexp
}

// Class Reference

func tagClass() *tagClass_ {
	return tagClassReference_
}

var tagClassReference_ = &tagClass_{
	// Initialize the class constants.
	matcher_: reg.MustCompile("^#((?:" + base32_ + ")+)"),
}
