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
	mod "github.com/craterdog/go-class-model/v5"
	gen "github.com/craterdog/go-code-generation/v5"
	fra "github.com/craterdog/go-collection-framework/v5"
	uti "github.com/craterdog/go-missing-utilities/v2"
	osx "os"
	sts "strings"
)

func main() {
	var tool = getTool()
	var version = getVersion()
	fmt.Printf("Tool: %v %v\n", tool, version)
	var moduleName, wikiPath, directory, packages = retrieveArguments()
	generateModule(moduleName, wikiPath, directory, packages)
	fmt.Println("Done.")
}

func generateModule(
	moduleName string,
	wikiPath string,
	directory string,
	packages []string,
) {
	var models = fra.Catalog[string, mod.ModelLike]()
	for _, packageName := range packages {
		var filename = directory + packageName + "/Package.go"
		var source = uti.ReadFile(filename)
		var model = mod.ParseSource(source)
		models.SetValue(packageName, model)
	}
	var generator = gen.ModuleGenerator()
	var existing string
	var filename = directory + "Module.go"
	if uti.PathExists(filename) {
		existing = uti.ReadFile(filename)
	}
	var moduleSynthesizer = gen.ModuleSynthesizer(
		moduleName,
		models,
	)
	var generated = generator.GenerateModule(
		moduleName,
		wikiPath,
		existing,
		moduleSynthesizer,
	)
	uti.WriteFile(filename, generated)
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
	directory string,
	packages []string,
) {
	if len(osx.Args) < 5 {
		var tool = getTool()
		fmt.Printf("  Usage: %v <moduleName> <wikiPath> <directory> <package>+ \n", tool)
		osx.Exit(1)
	}
	moduleName = osx.Args[1]
	wikiPath = osx.Args[2]
	directory = osx.Args[3]
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	packages = osx.Args[4:]
	return
}
