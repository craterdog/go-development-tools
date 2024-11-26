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
	mod "github.com/craterdog/go-class-model/v5"
	gen "github.com/craterdog/go-code-generation/v5"
	uti "github.com/craterdog/go-missing-utilities/v2"
	osx "os"
	sts "strings"
)

func main() {
	var tool = getTool()
	var version = getVersion()
	fmt.Printf("Tool: %v %v\n", tool, version)
	var moduleName, directory, packageName, force = retrieveArguments()
	var model = validateModelFile(directory)
	generatePackage(moduleName, directory, packageName, model, force)
	fmt.Println("Done.")
}

func generatePackage(
	moduleName string,
	directory string,
	packageName string,
	model mod.ModelLike,
	force bool,
) {
	if force {
		remakeDirectory(directory+packageName, force)
		var filename = directory + packageName + "/Package.go"
		var source = mod.FormatModel(model)
		writeFile(filename, source, force)
	}
	var generator = gen.ClassGenerator()
	var interfaceDeclarations = model.GetInterfaceDeclarations()
	var classSection = interfaceDeclarations.GetClassSection()
	var classDeclarations = classSection.GetClassDeclarations().GetIterator()
	for classDeclarations.HasNext() {
		var classDeclaration = classDeclarations.GetNext()
		var className = classDeclaration.GetDeclaration().GetName()
		className = sts.TrimSuffix(className, "ClassLike")
		className = uti.MakeLowerCase(className)
		var classSynthesizer = gen.ClassSynthesizer(model, className)
		var source = generator.GenerateClass(
			moduleName,
			packageName,
			className,
			classSynthesizer,
		)
		var filename = directory + packageName + "/" + className + ".go"
		writeFile(filename, source, force)
	}
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

func parseModel(modelFile string) mod.ModelLike {
	if !pathExists(modelFile) {
		fmt.Printf(
			"The model file %v does not exist, aborting...",
			modelFile,
		)
		osx.Exit(1)
	}
	var source = readFile(modelFile)
	var model = mod.ParseSource(source)
	return model
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

func remakeDirectory(
	directory string,
	force bool,
) {
	if force {
		var err = osx.RemoveAll(directory)
		if err != nil {
			panic(err)
		}
	}
	var err = osx.MkdirAll(directory, 0755)
	if err != nil {
		panic(err)
	}
}

func retrieveArguments() (
	moduleName string,
	packageName string,
	directory string,
	force bool,
) {
	if len(osx.Args) < 4 {
		var tool = getTool()
		fmt.Printf("  Usage: %v <moduleName> <directory> <packageName> [force]\n", tool)
		osx.Exit(1)
	}
	moduleName = osx.Args[1]
	directory = osx.Args[2]
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	packageName = osx.Args[3]
	force = len(osx.Args) == 5
	return
}

func validateModelFile(directory string) mod.ModelLike {
	fmt.Println("  Validating the Package.go file...")
	var modelFile = directory + "Packet.go"
	var model = parseModel(modelFile)
	mod.ValidateModel(model)
	return model
}

func writeFile(
	filename string,
	source string,
	force bool,
) {
	if pathExists(filename) && !force {
		fmt.Printf(
			"    The following package already exists, not overwriting it:\n%v\n",
			filename,
		)
		return
	}
	var bytes = []byte(source)
	var err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
}
