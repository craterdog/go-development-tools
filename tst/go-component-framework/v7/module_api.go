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

/*
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│             This "module_api.go" file was automatically generated.           │
│      Updates to any part of this file—other than the Module Description      │
│             and the Global Functions sections may be overwritten.            │
└──────────────────────────────────────────────────────────────────────────────┘

Package "module" declares type aliases for the commonly used types declared in
the packages contained in this module.  It also provides constructors for each
commonly used class that is exported by the module.  Each constructor delegates
the actual construction process to its corresponding concrete class declared in
the corresponding package contained within this module.

For detailed documentation on this entire module refer to the wiki:
  - https://github.com/craterdog/go-component-framework/wiki
*/
package module

import (
	age "github.com/craterdog/go-component-framework/v7/agent"
	ele "github.com/craterdog/go-component-framework/v7/element"
	str "github.com/craterdog/go-component-framework/v7/string"
	uti "github.com/craterdog/go-missing-utilities/v7"
	uri "net/url"
)

// TYPE ALIASES

// Element

type (
	Units = ele.Units
)

const (
	Degrees  = ele.Degrees
	Radians  = ele.Radians
	Gradians = ele.Gradians
)

type (
	AngleClassLike       = ele.AngleClassLike
	BooleanClassLike     = ele.BooleanClassLike
	DurationClassLike    = ele.DurationClassLike
	GlyphClassLike       = ele.GlyphClassLike
	MomentClassLike      = ele.MomentClassLike
	NumberClassLike      = ele.NumberClassLike
	PercentageClassLike  = ele.PercentageClassLike
	ProbabilityClassLike = ele.ProbabilityClassLike
	ResourceClassLike    = ele.ResourceClassLike
	SymbolClassLike      = ele.SymbolClassLike
)

type (
	AngleLike       = ele.AngleLike
	BooleanLike     = ele.BooleanLike
	DurationLike    = ele.DurationLike
	GlyphLike       = ele.GlyphLike
	MomentLike      = ele.MomentLike
	NumberLike      = ele.NumberLike
	PercentageLike  = ele.PercentageLike
	ProbabilityLike = ele.ProbabilityLike
	ResourceLike    = ele.ResourceLike
	SymbolLike      = ele.SymbolLike
)

type (
	Continuous = ele.Continuous
	Discrete   = ele.Discrete
	Factored   = ele.Factored
	Polarized  = ele.Polarized
	Temporal   = ele.Temporal
)

// String

type (
	Identifier  = str.Identifier
	Instruction = str.Instruction
	Line        = str.Line
	Character   = str.Character
)

type (
	BinaryClassLike    = str.BinaryClassLike
	BytecodeClassLike  = str.BytecodeClassLike
	NameClassLike      = str.NameClassLike
	NarrativeClassLike = str.NarrativeClassLike
	PatternClassLike   = str.PatternClassLike
	QuoteClassLike     = str.QuoteClassLike
	TagClassLike       = str.TagClassLike
	VersionClassLike   = str.VersionClassLike
)

type (
	BinaryLike    = str.BinaryLike
	BytecodeLike  = str.BytecodeLike
	NameLike      = str.NameLike
	NarrativeLike = str.NarrativeLike
	PatternLike   = str.PatternLike
	QuoteLike     = str.QuoteLike
	TagLike       = str.TagLike
	VersionLike   = str.VersionLike
)

type (
	Accessible[V any] = str.Accessible[V]
	Sequential[V any] = str.Sequential[V]
)

// Agent

type (
	Event       = age.Event
	Rank        = age.Rank
	Slot        = age.Slot
	State       = age.State
	Transitions = age.Transitions
)

const (
	LesserRank  = age.LesserRank
	EqualRank   = age.EqualRank
	GreaterRank = age.GreaterRank
)

type (
	RankingFunction[V any] = age.RankingFunction[V]
)

type (
	CollatorClassLike[V any] = age.CollatorClassLike[V]
	ControllerClassLike      = age.ControllerClassLike
	EncoderClassLike         = age.EncoderClassLike
	GeneratorClassLike       = age.GeneratorClassLike
	IteratorClassLike[V any] = age.IteratorClassLike[V]
	SorterClassLike[V any]   = age.SorterClassLike[V]
)

type (
	CollatorLike[V any] = age.CollatorLike[V]
	ControllerLike      = age.ControllerLike
	EncoderLike         = age.EncoderLike
	GeneratorLike       = age.GeneratorLike
	IteratorLike[V any] = age.IteratorLike[V]
	SorterLike[V any]   = age.SorterLike[V]
)

// CLASS ACCESSORS

// Element

func AngleClass() AngleClassLike {
	return ele.AngleClass()
}

func Angle(
	radians float64,
) AngleLike {
	return AngleClass().Angle(
		radians,
	)
}

func AngleFromString(
	string_ string,
) AngleLike {
	return AngleClass().AngleFromString(
		string_,
	)
}

func BooleanClass() BooleanClassLike {
	return ele.BooleanClass()
}

func Boolean(
	boolean bool,
) BooleanLike {
	return BooleanClass().Boolean(
		boolean,
	)
}

func BooleanFromString(
	string_ string,
) BooleanLike {
	return BooleanClass().BooleanFromString(
		string_,
	)
}

func DurationClass() DurationClassLike {
	return ele.DurationClass()
}

func Duration(
	milliseconds int,
) DurationLike {
	return DurationClass().Duration(
		milliseconds,
	)
}

func DurationFromString(
	string_ string,
) DurationLike {
	return DurationClass().DurationFromString(
		string_,
	)
}

func GlyphClass() GlyphClassLike {
	return ele.GlyphClass()
}

func Glyph(
	rune_ rune,
) GlyphLike {
	return GlyphClass().Glyph(
		rune_,
	)
}

func GlyphFromInteger(
	integer int,
) GlyphLike {
	return GlyphClass().GlyphFromInteger(
		integer,
	)
}

func GlyphFromString(
	string_ string,
) GlyphLike {
	return GlyphClass().GlyphFromString(
		string_,
	)
}

func MomentClass() MomentClassLike {
	return ele.MomentClass()
}

func Moment(
	milliseconds int,
) MomentLike {
	return MomentClass().Moment(
		milliseconds,
	)
}

func MomentFromString(
	string_ string,
) MomentLike {
	return MomentClass().MomentFromString(
		string_,
	)
}

func NumberClass() NumberClassLike {
	return ele.NumberClass()
}

func Number(
	complex_ complex128,
) NumberLike {
	return NumberClass().Number(
		complex_,
	)
}

func NumberFromPolar(
	magnitude float64,
	phase float64,
) NumberLike {
	return NumberClass().NumberFromPolar(
		magnitude,
		phase,
	)
}

func NumberFromRectangular(
	real_ float64,
	imaginary float64,
) NumberLike {
	return NumberClass().NumberFromRectangular(
		real_,
		imaginary,
	)
}

func NumberFromString(
	string_ string,
) NumberLike {
	return NumberClass().NumberFromString(
		string_,
	)
}

func PercentageClass() PercentageClassLike {
	return ele.PercentageClass()
}

func Percentage(
	float float64,
) PercentageLike {
	return PercentageClass().Percentage(
		float,
	)
}

func PercentageFromInteger(
	integer int,
) PercentageLike {
	return PercentageClass().PercentageFromInteger(
		integer,
	)
}

func PercentageFromString(
	string_ string,
) PercentageLike {
	return PercentageClass().PercentageFromString(
		string_,
	)
}

func ProbabilityClass() ProbabilityClassLike {
	return ele.ProbabilityClass()
}

func Probability(
	float float64,
) ProbabilityLike {
	return ProbabilityClass().Probability(
		float,
	)
}

func ProbabilityFromBoolean(
	boolean bool,
) ProbabilityLike {
	return ProbabilityClass().ProbabilityFromBoolean(
		boolean,
	)
}

func ProbabilityFromString(
	string_ string,
) ProbabilityLike {
	return ProbabilityClass().ProbabilityFromString(
		string_,
	)
}

func ResourceClass() ResourceClassLike {
	return ele.ResourceClass()
}

func Resource(
	string_ string,
) ResourceLike {
	return ResourceClass().Resource(
		string_,
	)
}

func ResourceFromUri(
	url *uri.URL,
) ResourceLike {
	return ResourceClass().ResourceFromUri(
		url,
	)
}

func SymbolClass() SymbolClassLike {
	return ele.SymbolClass()
}

func Symbol(
	string_ string,
) SymbolLike {
	return SymbolClass().Symbol(
		string_,
	)
}

// String

func BinaryClass() BinaryClassLike {
	return str.BinaryClass()
}

func Binary(
	bytes []byte,
) BinaryLike {
	return BinaryClass().Binary(
		bytes,
	)
}

func BinaryFromSequence(
	sequence str.Sequential[byte],
) BinaryLike {
	return BinaryClass().BinaryFromSequence(
		sequence,
	)
}

func BinaryFromString(
	string_ string,
) BinaryLike {
	return BinaryClass().BinaryFromString(
		string_,
	)
}

func BytecodeClass() BytecodeClassLike {
	return str.BytecodeClass()
}

func Bytecode(
	instructions []str.Instruction,
) BytecodeLike {
	return BytecodeClass().Bytecode(
		instructions,
	)
}

func BytecodeFromSequence(
	sequence str.Sequential[str.Instruction],
) BytecodeLike {
	return BytecodeClass().BytecodeFromSequence(
		sequence,
	)
}

func BytecodeFromString(
	string_ string,
) BytecodeLike {
	return BytecodeClass().BytecodeFromString(
		string_,
	)
}

func NameClass() NameClassLike {
	return str.NameClass()
}

func Name(
	identifiers []str.Identifier,
) NameLike {
	return NameClass().Name(
		identifiers,
	)
}

func NameFromSequence(
	sequence str.Sequential[str.Identifier],
) NameLike {
	return NameClass().NameFromSequence(
		sequence,
	)
}

func NameFromString(
	string_ string,
) NameLike {
	return NameClass().NameFromString(
		string_,
	)
}

func NarrativeClass() NarrativeClassLike {
	return str.NarrativeClass()
}

func Narrative(
	lines []str.Line,
) NarrativeLike {
	return NarrativeClass().Narrative(
		lines,
	)
}

func NarrativeFromSequence(
	sequence str.Sequential[str.Line],
) NarrativeLike {
	return NarrativeClass().NarrativeFromSequence(
		sequence,
	)
}

func NarrativeFromString(
	string_ string,
) NarrativeLike {
	return NarrativeClass().NarrativeFromString(
		string_,
	)
}

func PatternClass() PatternClassLike {
	return str.PatternClass()
}

func Pattern(
	characters []str.Character,
) PatternLike {
	return PatternClass().Pattern(
		characters,
	)
}

func PatternFromSequence(
	sequence str.Sequential[str.Character],
) PatternLike {
	return PatternClass().PatternFromSequence(
		sequence,
	)
}

func PatternFromString(
	string_ string,
) PatternLike {
	return PatternClass().PatternFromString(
		string_,
	)
}

func QuoteClass() QuoteClassLike {
	return str.QuoteClass()
}

func Quote(
	characters []str.Character,
) QuoteLike {
	return QuoteClass().Quote(
		characters,
	)
}

func QuoteFromSequence(
	sequence str.Sequential[str.Character],
) QuoteLike {
	return QuoteClass().QuoteFromSequence(
		sequence,
	)
}

func QuoteFromString(
	string_ string,
) QuoteLike {
	return QuoteClass().QuoteFromString(
		string_,
	)
}

func TagClass() TagClassLike {
	return str.TagClass()
}

func Tag(
	bytes []byte,
) TagLike {
	return TagClass().Tag(
		bytes,
	)
}

func TagWithSize(
	size uti.Cardinal,
) TagLike {
	return TagClass().TagWithSize(
		size,
	)
}

func TagFromSequence(
	sequence str.Sequential[byte],
) TagLike {
	return TagClass().TagFromSequence(
		sequence,
	)
}

func TagFromString(
	string_ string,
) TagLike {
	return TagClass().TagFromString(
		string_,
	)
}

func VersionClass() VersionClassLike {
	return str.VersionClass()
}

func Version(
	ordinals []uti.Ordinal,
) VersionLike {
	return VersionClass().Version(
		ordinals,
	)
}

func VersionFromSequence(
	sequence str.Sequential[uti.Ordinal],
) VersionLike {
	return VersionClass().VersionFromSequence(
		sequence,
	)
}

func VersionFromString(
	string_ string,
) VersionLike {
	return VersionClass().VersionFromString(
		string_,
	)
}

// Agent

func CollatorClass[V any]() CollatorClassLike[V] {
	return age.CollatorClass[V]()
}

func Collator[V any]() CollatorLike[V] {
	return CollatorClass[V]().Collator()
}

func CollatorWithMaximumDepth[V any](
	maximumDepth uti.Cardinal,
) CollatorLike[V] {
	return CollatorClass[V]().CollatorWithMaximumDepth(
		maximumDepth,
	)
}

func ControllerClass() ControllerClassLike {
	return age.ControllerClass()
}

func Controller(
	events []age.Event,
	transitions map[State]age.Transitions,
	initialState age.State,
) ControllerLike {
	return ControllerClass().Controller(
		events,
		transitions,
		initialState,
	)
}

func EncoderClass() EncoderClassLike {
	return age.EncoderClass()
}

func Encoder() EncoderLike {
	return EncoderClass().Encoder()
}

func GeneratorClass() GeneratorClassLike {
	return age.GeneratorClass()
}

func Generator() GeneratorLike {
	return GeneratorClass().Generator()
}

func IteratorClass[V any]() IteratorClassLike[V] {
	return age.IteratorClass[V]()
}

func Iterator[V any](
	array []V,
) IteratorLike[V] {
	return IteratorClass[V]().Iterator(
		array,
	)
}

func SorterClass[V any]() SorterClassLike[V] {
	return age.SorterClass[V]()
}

func Sorter[V any]() SorterLike[V] {
	return SorterClass[V]().Sorter()
}

func SorterWithRanker[V any](
	ranker age.RankingFunction[V],
) SorterLike[V] {
	return SorterClass[V]().SorterWithRanker(
		ranker,
	)
}

// GLOBAL FUNCTIONS
