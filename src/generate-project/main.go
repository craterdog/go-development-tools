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
	gen "github.com/craterdog/go-code-generation/v6"
	fra "github.com/craterdog/go-collection-framework/v5"
	uti "github.com/craterdog/go-missing-utilities/v2"
	not "github.com/craterdog/go-syntax-notation/v6"
	osx "os"
	sts "strings"
)

func main() {
	var tool = getTool()
	var version = getVersion()
	fmt.Printf("Tool: %v %v\n", tool, version)
	var moduleName, wikiPath, directory = retrieveArguments()
	var syntaxFile = directory + "Syntax.cdsn"
	var syntax = parseSyntax(syntaxFile)
	validateSyntax(syntax)
	generateAstPackage(moduleName, wikiPath, directory, syntax)
	generateGrammarPackage(moduleName, wikiPath, directory, syntax)
	generateModule(moduleName, wikiPath, directory)
	fmt.Println("Done.")
}

func generateAstPackage(
	moduleName string,
	wikiPath string,
	directory string,
	syntax not.SyntaxLike,
) {
	var packageName = "ast"
	uti.RemakeDirectory(directory + packageName)
	var filename = directory + packageName + "/Package.go"
	fmt.Printf("  Generating %v...\n", filename)
	var model = gen.GenerateAstPackage(
		moduleName,
		wikiPath,
		syntax,
	)
	var source = mod.FormatModel(model)
	uti.WriteFile(filename, source)
	generateAstClasses(moduleName, packageName, directory, model)
}

func generateAstClasses(
	moduleName string,
	packageName string,
	directory string,
	model mod.ModelLike,
) {
	var interfaceDeclarations = model.GetInterfaceDeclarations()
	var classSection = interfaceDeclarations.GetClassSection()
	var classDeclarations = classSection.GetClassDeclarations().GetIterator()
	for classDeclarations.HasNext() {
		var classDeclaration = classDeclarations.GetNext()
		var className = classDeclaration.GetDeclaration().GetName()
		className = sts.TrimSuffix(className, "ClassLike")
		className = uti.MakeLowerCase(className)
		var source = gen.GenerateAstClass(
			moduleName,
			className,
			model,
		)
		var filename = directory + packageName + "/" + className + ".go"
		uti.WriteFile(filename, source)
	}
}

func generateFormatter(
	moduleName string,
	directory string,
	syntax not.SyntaxLike,
) {
	var existing string
	var filename = directory + "grammar/formatter.go"
	if uti.PathExists(filename) {
		existing = uti.ReadFile(filename)
	}
	var source = gen.GenerateFormatterClass(
		moduleName,
		existing,
		syntax,
	)
	uti.WriteFile(filename, source)
}

func generateGrammarPackage(
	moduleName string,
	wikiPath string,
	directory string,
	syntax not.SyntaxLike,
) {
	var filename = directory + "grammar/Package.go"
	fmt.Printf("  Generating %v...\n", filename)
	var model = gen.GenerateGrammarPackage(
		moduleName,
		wikiPath,
		syntax,
	)
	var source = mod.FormatModel(model)
	uti.WriteFile(filename, source)
	generateFormatter(moduleName, directory, syntax)
	generateParser(moduleName, directory, syntax)
	generateProcessor(moduleName, directory, syntax)
	generateScanner(moduleName, directory, syntax)
	generateToken(moduleName, directory, syntax)
	generateValidator(moduleName, directory, syntax)
	generateVisitor(moduleName, directory, syntax)
}

func generateModule(
	moduleName string,
	wikiPath string,
	directory string,
) {
	var packages = []string{
		"ast",
		"grammar",
	}
	var models = fra.Catalog[string, mod.ModelLike]()
	for _, packageName := range packages {
		var filename = directory + packageName + "/Package.go"
		var source = uti.ReadFile(filename)
		var model = mod.ParseSource(source)
		models.SetValue(packageName, model)
	}
	var filename = directory + "Module.go"
	fmt.Printf("  Generating %v...\n", filename)
	var existing string
	if uti.PathExists(filename) {
		existing = uti.ReadFile(filename)
	}
	var source = gen.GenerateModule(
		moduleName,
		wikiPath,
		existing,
		models,
	)
	uti.WriteFile(filename, source)
}

func generateParser(
	moduleName string,
	directory string,
	syntax not.SyntaxLike,
) {
	var source = gen.GenerateParserClass(
		moduleName,
		syntax,
	)
	var filename = directory + "grammar/parser.go"
	uti.WriteFile(filename, source)
}

func generateProcessor(
	moduleName string,
	directory string,
	syntax not.SyntaxLike,
) {
	var source = gen.GenerateProcessorClass(
		moduleName,
		syntax,
	)
	var filename = directory + "grammar/processor.go"
	uti.WriteFile(filename, source)
}

func generateScanner(
	moduleName string,
	directory string,
	syntax not.SyntaxLike,
) {
	var source = gen.GenerateScannerClass(
		moduleName,
		syntax,
	)
	var filename = directory + "grammar/scanner.go"
	uti.WriteFile(filename, source)
}

func generateToken(
	moduleName string,
	directory string,
	syntax not.SyntaxLike,
) {
	var source = gen.GenerateTokenClass(
		moduleName,
		syntax,
	)
	var filename = directory + "grammar/token.go"
	uti.WriteFile(filename, source)
}

func generateValidator(
	moduleName string,
	directory string,
	syntax not.SyntaxLike,
) {
	var existing string
	var filename = directory + "grammar/validator.go"
	if uti.PathExists(filename) {
		existing = uti.ReadFile(filename)
	}
	var source = gen.GenerateValidatorClass(
		moduleName,
		existing,
		syntax,
	)
	uti.WriteFile(filename, source)
}

func generateVisitor(
	moduleName string,
	directory string,
	syntax not.SyntaxLike,
) {
	var source = gen.GenerateVisitorClass(
		moduleName,
		syntax,
	)
	var filename = directory + "grammar/visitor.go"
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

func parseSyntax(syntaxFile string) not.SyntaxLike {
	fmt.Printf("  Parsing %v...\n", syntaxFile)
	if !uti.PathExists(syntaxFile) {
		fmt.Println("The syntax file does not exist, aborting...")
		osx.Exit(1)
	}
	var source = uti.ReadFile(syntaxFile)
	var syntax = not.ParseSource(source)
	return syntax
}

func retrieveArguments() (
	moduleName string,
	wikiPath string,
	directory string,
) {
	if len(osx.Args) < 4 {
		var tool = getTool()
		fmt.Printf("  Usage: %v <moduleName> <wikiPath> <directory>\n", tool)
		osx.Exit(1)
	}
	moduleName = osx.Args[1]
	wikiPath = osx.Args[2]
	directory = osx.Args[3]
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	return
}

func validateSyntax(
	syntax not.SyntaxLike,
) {
	fmt.Println("  Validating the syntax...")
	not.ValidateSyntax(syntax)
}
