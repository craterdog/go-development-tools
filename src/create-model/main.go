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
	mod "github.com/craterdog/go-class-model/v8"
	gen "github.com/craterdog/go-code-generation/v8"
	uti "github.com/craterdog/go-essential-utilities/v8"
	osx "os"
	sts "strings"
)

func main() {
	var tool = getTool()
	var version = getVersion()
	fmt.Printf("Tool: %v %v\n", tool, version)
	var moduleName, wikiPath, packageName, directory = retrieveArguments()
	createPackage(moduleName, wikiPath, packageName, directory)
	fmt.Println("Done.")
}

func createPackage(
	moduleName string,
	wikiPath string,
	packageName string,
	directory string,
) {
	uti.RemakeDirectory(directory + packageName)
	var filename = directory + packageName + "/package_api.go"
	fmt.Printf("  Generating %v...\n", filename)
	var model = gen.CreatePackage(
		moduleName,
		wikiPath,
		packageName,
	)
	mod.ValidateModel(model)
	var source = mod.FormatModel(model)
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
	moduleName string,
	wikiPath string,
	packageName string,
	directory string,
) {
	if len(osx.Args) < 5 {
		var tool = getTool()
		fmt.Printf("  Usage: %v <moduleName> <wikiPath> <packageName> <directory>\n", tool)
		osx.Exit(1)
	}
	moduleName = osx.Args[1]
	wikiPath = osx.Args[2]
	packageName = osx.Args[3]
	directory = osx.Args[4]
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	return
}
