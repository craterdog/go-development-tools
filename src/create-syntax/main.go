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

package main

import (
	fmt "fmt"
	uti "github.com/craterdog/go-missing-utilities/v7"
	osx "os"
	sts "strings"
)

func main() {
	var tool = getTool()
	var version = getVersion()
	fmt.Printf("Tool: %v %v\n", tool, version)
	var directory = retrieveArguments()
	createSyntax(directory)
	fmt.Println("Done.")
}

func createSyntax(
	directory string,
) {
	uti.RemakeDirectory(directory)
	var filename = directory + "syntax.cdsn"
	fmt.Printf("  Generating %v...\n", filename)
	var source = exampleSyntax()
	uti.WriteFile(filename, source)
}

func getTool() string {
	var tool = osx.Args[0]
	tool = sts.TrimPrefix(tool, "./")
	return tool
}

func getVersion() string {
	var modFile = "./go.mod"
	var source = uti.ReadFile(modFile)
	var lines = sts.Split(source, "\n")
	var version = sts.Split(lines[6], " ")[1]
	return version
}

func retrieveArguments() (
	directory string,
) {
	if len(osx.Args) < 2 {
		var tool = getTool()
		fmt.Printf("  Usage: %v <directory>\n", tool)
		osx.Exit(1)
	}
	directory = osx.Args[1]
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	return
}

func exampleSyntax() string {
	return `!>
................................................................................
.    Copyright (c) 2009-2025 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
<!

!>
EXAMPLE NOTATION
This document contains a formal definition of an example syntax notation using
Crater Dog Syntax Notation™ (CDSN).

For more information on CDSN see the wiki at:
  - https://github.com/craterdog/go-syntax-notation/wiki

┌──────────────────────────────────────────────────────────────────────────────┐
│                               RULE DEFINITIONS                               │
└──────────────────────────────────────────────────────────────────────────────┘
<!
$Document: Component{3..}  ! An inline comment.

$Component:
    Intrinsic
    List

$Intrinsic:
    number
    rune  ! A multiline comment.
    text

$List: "[" Component Additional* "]"

$Additional: "," Component

!>
┌──────────────────────────────────────────────────────────────────────────────┐
│                            EXPRESSION DEFINITIONS                            │
└──────────────────────────────────────────────────────────────────────────────┘
<!
$number: '0' | '-'? ORDINAL

$rune: ''' ~[CONTROL] "'"  ! Any single printable unicode character.

$text: '"' ~['"' CONTROL]+ '"'

!>
┌──────────────────────────────────────────────────────────────────────────────┐
│                             FRAGMENT DEFINITIONS                             │
└──────────────────────────────────────────────────────────────────────────────┘
<!
$ORDINAL: ['1'..'9'] DIGIT*

`
}
