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

package main

import (
	fmt "fmt"
	not "github.com/craterdog/go-syntax-notation/v5"
	osx "os"
	sts "strings"
)

func main() {
	var tool = getTool()
	var version = getVersion()
	fmt.Printf("Tool: %v %v\n", tool, version)
	var syntaxFile = retrieveArguments()
	var syntax = parseSyntax(syntaxFile)
	validateSyntax(syntax)
	formatSyntaxFile(syntaxFile, syntax)
	fmt.Println("Done.")
}

func getTool() string {
	var tool = osx.Args[0]
	tool = sts.TrimPrefix(tool, "./")
	return tool
}

func getVersion() string {
	var modFile = "./go.mod"
	var source = readFile(modFile)
	var lines = sts.Split(source, "\n")
	var version = sts.Split(lines[6], " ")[1]
	return version
}

func parseSyntax(syntaxFile string) not.SyntaxLike {
	fmt.Printf("  Parsing the following syntax file:\n    %v\n", syntaxFile)
	if !pathExists(syntaxFile) {
		fmt.Println("The syntax file does not exist, aborting...")
		osx.Exit(1)
	}
	var source = readFile(syntaxFile)
	var syntax = not.ParseSource(source)
	return syntax
}

func pathExists(path string) bool {
	var _, err = osx.Stat(path)
	if err == nil {
		return true
	}
	if osx.IsNotExist(err) {
		return false
	}
	panic(err)
}

func readFile(
	filename string,
) string {
	var bytes, err = osx.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var source = string(bytes)
	return source
}

func formatSyntaxFile(
	syntaxFile string,
	syntax not.SyntaxLike,
) {
	fmt.Println("  Formatting the syntax file...")
	var source = not.FormatSyntax(syntax)
	writeFile(syntaxFile, source)
}

func retrieveArguments() (
	syntaxFile string,
) {
	if len(osx.Args) < 2 {
		var tool = getTool()
		fmt.Printf("  Usage: %v <syntaxFile>\n", tool)
		osx.Exit(1)
	}
	syntaxFile = osx.Args[1]
	return
}

func validateSyntax(syntax not.SyntaxLike) {
	fmt.Println("  Validating the syntax...")
	not.ValidateSyntax(syntax)
}

func writeFile(
	filename string,
	source string,
) {
	var bytes = []byte(source)
	var err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
}
