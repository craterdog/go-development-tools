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
	not "github.com/craterdog/go-syntax-notation/v5"
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
	generateModuleFile(moduleName, wikiPath, directory)
	fmt.Println("Done.")
}

func generateAstPackage(
	moduleName string,
	wikiPath string,
	directory string,
	syntax not.SyntaxLike,
) {
	var packageName = "ast"
	var filename = directory + packageName + "/Package.go"
	fmt.Printf("  Generating %v...\n", filename)
	uti.RemakeDirectory(directory + packageName)
	var generator = gen.PackageGenerator()
	var astSynthesizer = gen.AstSynthesizer(syntax)
	var source = generator.GeneratePackage(
		moduleName,
		wikiPath,
		packageName,
		"", // The AST package does not yet exist.
		astSynthesizer,
	)
	uti.WriteFile(filename, source)
	var astModel = parseModel(source)
	generateClasses(moduleName, directory, packageName, astModel)
}

func generateClasses(
	moduleName string,
	directory string,
	packageName string,
	astModel mod.ModelLike,
) {
	var generator = gen.ClassGenerator()
	var interfaceDeclarations = astModel.GetInterfaceDeclarations()
	var classSection = interfaceDeclarations.GetClassSection()
	var classDeclarations = classSection.GetClassDeclarations().GetIterator()
	for classDeclarations.HasNext() {
		var classDeclaration = classDeclarations.GetNext()
		var className = classDeclaration.GetDeclaration().GetName()
		className = sts.TrimSuffix(className, "ClassLike")
		className = uti.MakeLowerCase(className)
		var nodeSynthesizer = gen.NodeSynthesizer(astModel, className)
		var source = generator.GenerateClass(
			moduleName,
			packageName,
			className,
			"", // The class will be completely overwritten.
			nodeSynthesizer,
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
	var generator = gen.ClassGenerator()
	var packageName = "grammar"
	var className = "formatter"
	var existing string
	var filename = directory + packageName + "/" + className + ".go"
	if uti.PathExists(filename) {
		existing = uti.ReadFile(filename)
	}
	var formatterSynthesizer = gen.FormatterSynthesizer(syntax)
	var generated = generator.GenerateClass(
		moduleName,
		packageName,
		className,
		existing,
		formatterSynthesizer,
	)
	uti.WriteFile(filename, generated)
}

func generateGrammarPackage(
	moduleName string,
	wikiPath string,
	directory string,
	syntax not.SyntaxLike,
) {
	var packageName = "grammar"
	var filename = directory + packageName + "/Package.go"
	fmt.Printf("  Generating %v...\n", filename)
	var generator = gen.PackageGenerator()
	var grammarSynthesizer = gen.GrammarSynthesizer(syntax)
	var source = generator.GeneratePackage(
		moduleName,
		wikiPath,
		packageName,
		"", // The grammar package does not yet exist.
		grammarSynthesizer,
	)
	uti.WriteFile(filename, source)
	generateFormatter(moduleName, directory, syntax)
	generateParser(moduleName, directory, syntax)
	generateProcessor(moduleName, directory, syntax)
	generateScanner(moduleName, directory, syntax)
	generateToken(moduleName, directory, syntax)
	generateValidator(moduleName, directory, syntax)
	generateVisitor(moduleName, directory, syntax)
}

func generateModuleFile(
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
	var moduleSynthesizer = gen.ModuleSynthesizer(models)
	var generator = gen.ModuleGenerator()
	var generated = generator.GenerateModule(
		moduleName,
		wikiPath,
		existing,
		moduleSynthesizer,
	)
	uti.WriteFile(filename, generated)
}

func generateParser(
	moduleName string,
	directory string,
	syntax not.SyntaxLike,
) {
	var generator = gen.ClassGenerator()
	var packageName = "grammar"
	var className = "parser"
	var parserSynthesizer = gen.ParserSynthesizer(syntax)
	var source = generator.GenerateClass(
		moduleName,
		packageName,
		className,
		"", // The parser class will be completely overwritten.
		parserSynthesizer,
	)
	var filename = directory + packageName + "/" + className + ".go"
	uti.WriteFile(filename, source)
}

func generateProcessor(
	moduleName string,
	directory string,
	syntax not.SyntaxLike,
) {
	var generator = gen.ClassGenerator()
	var packageName = "grammar"
	var className = "processor"
	var processorSynthesizer = gen.ProcessorSynthesizer(syntax)
	var source = generator.GenerateClass(
		moduleName,
		packageName,
		className,
		"", // The processor class will be completely overwritten.
		processorSynthesizer,
	)
	var filename = directory + packageName + "/" + className + ".go"
	uti.WriteFile(filename, source)
}

func generateScanner(
	moduleName string,
	directory string,
	syntax not.SyntaxLike,
) {
	var generator = gen.ClassGenerator()
	var packageName = "grammar"
	var className = "scanner"
	var scannerSynthesizer = gen.ScannerSynthesizer(syntax)
	var source = generator.GenerateClass(
		moduleName,
		packageName,
		className,
		"", // The scanner class will be completely overwritten.
		scannerSynthesizer,
	)
	var filename = directory + packageName + "/" + className + ".go"
	uti.WriteFile(filename, source)
}

func generateToken(
	moduleName string,
	directory string,
	syntax not.SyntaxLike,
) {
	var generator = gen.ClassGenerator()
	var packageName = "grammar"
	var className = "token"
	var tokenSynthesizer = gen.TokenSynthesizer(syntax)
	var source = generator.GenerateClass(
		moduleName,
		packageName,
		className,
		"", // The token class will be completely overwritten.
		tokenSynthesizer,
	)
	var filename = directory + packageName + "/" + className + ".go"
	uti.WriteFile(filename, source)
}

func generateValidator(
	moduleName string,
	directory string,
	syntax not.SyntaxLike,
) {
	var generator = gen.ClassGenerator()
	var packageName = "grammar"
	var className = "validator"
	var filename = directory + packageName + "/" + className + ".go"
	var existing string
	if uti.PathExists(filename) {
		existing = uti.ReadFile(filename)
	}
	var validatorSynthesizer = gen.ValidatorSynthesizer(syntax)
	var generated = generator.GenerateClass(
		moduleName,
		packageName,
		className,
		existing,
		validatorSynthesizer,
	)
	uti.WriteFile(filename, generated)
}

func generateVisitor(
	moduleName string,
	directory string,
	syntax not.SyntaxLike,
) {
	var generator = gen.ClassGenerator()
	var packageName = "grammar"
	var className = "visitor"
	var visitorSynthesizer = gen.VisitorSynthesizer(syntax)
	var source = generator.GenerateClass(
		moduleName,
		packageName,
		className,
		"", // The visitor class will be completely overwritten.
		visitorSynthesizer,
	)
	var filename = directory + packageName + "/" + className + ".go"
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

func parseModel(
	source string,
) mod.ModelLike {
	var model = mod.ParseSource(source)
	mod.ValidateModel(model)
	return model
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
