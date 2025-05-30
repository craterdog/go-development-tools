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
	mod "github.com/craterdog/go-class-model/v7"
	gen "github.com/craterdog/go-code-generation/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
	osx "os"
	sts "strings"
)

func main() {
	var tool = getTool()
	var version = getVersion()
	fmt.Printf("Tool: %v %v\n", tool, version)
	var moduleName, directory, packageName = retrieveArguments()
	var model = validateModelFile(directory + packageName)
	generatePackage(moduleName, directory, packageName, model)
	fmt.Println("Done.")
}

func generatePackage(
	moduleName string,
	directory string,
	packageName string,
	model mod.ModelLike,
) {
	uti.RemakeDirectory(directory + packageName)
	var filename = directory + packageName + "/package_api.go"
	mod.ValidateModel(model)
	var source = mod.FormatModel(model)
	uti.WriteFile(filename, source)
	fmt.Println("  Generating the following classes...")
	var existing string
	var interfaceDeclarations = model.GetInterfaceDeclarations()
	var classSection = interfaceDeclarations.GetClassSection()
	var classDeclarations = classSection.GetClassDeclarations().GetIterator()
	for classDeclarations.HasNext() {
		var classDeclaration = classDeclarations.GetNext()
		var className = classDeclaration.GetDeclaration().GetName()
		className = sts.TrimSuffix(className, "ClassLike")
		className = uti.MakeLowerCase(className)
		fmt.Printf("    %s/%s.go\n", packageName, className)
		var source = gen.GenerateClass(
			moduleName,
			packageName,
			className,
			existing,
			model,
		)
		className = uti.MakeUpperCase(className)
		var filename = directory + packageName + "/" + className + ".go"
		uti.WriteFile(filename, source)
	}
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

func parseModel(modelFile string) mod.ModelLike {
	if !uti.PathExists(modelFile) {
		fmt.Printf(
			"The model file %v does not exist, aborting...",
			modelFile,
		)
		osx.Exit(1)
	}
	var source = uti.ReadFile(modelFile)
	var model = mod.ParseSource(source)
	return model
}

func retrieveArguments() (
	moduleName string,
	directory string,
	packageName string,
) {
	if len(osx.Args) < 4 {
		var tool = getTool()
		fmt.Printf("  Usage: %v <moduleName> <directory> <packageName>\n", tool)
		osx.Exit(1)
	}
	moduleName = osx.Args[1]
	directory = osx.Args[2]
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	packageName = osx.Args[3]
	return
}

func validateModelFile(directory string) mod.ModelLike {
	fmt.Println("  Validating the package_api.go file...")
	var modelFile = directory + "/package_api.go"
	var model = parseModel(modelFile)
	mod.ValidateModel(model)
	return model
}
