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
	uti "github.com/craterdog/go-missing-utilities/v2"
	osx "os"
	sts "strings"
)

func main() {
	var tool = getTool()
	var version = getVersion()
	fmt.Printf("Tool: %v %v\n", tool, version)
	var modelFile = retrieveArguments()
	var model = parseModel(modelFile)
	validateModel(model)
	formatModelFile(modelFile, model)
	fmt.Println("Done.")
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
	fmt.Printf("  Parsing the following model file:\n    %v\n", modelFile)
	if !uti.PathExists(modelFile) {
		fmt.Println("The model file does not exist, aborting...")
		osx.Exit(1)
	}
	var source = uti.ReadFile(modelFile)
	var model = mod.ParseSource(source)
	return model
}

func formatModelFile(
	modelFile string,
	model mod.ModelLike,
) {
	fmt.Println("  Formatting the model file...")
	var source = mod.FormatModel(model)
	uti.WriteFile(modelFile, source)
}

func retrieveArguments() (
	modelFile string,
) {
	if len(osx.Args) < 2 {
		var tool = getTool()
		fmt.Printf("  Usage: %v <modelFile>\n", tool)
		osx.Exit(1)
	}
	modelFile = osx.Args[1]
	return
}

func validateModel(
	model mod.ModelLike,
) {
	fmt.Println("  Validating the model...")
	mod.ValidateModel(model)
}
