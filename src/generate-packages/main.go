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
	col "github.com/craterdog/go-collection-framework/v4"
	uti "github.com/craterdog/go-missing-utilities/v2"
	not "github.com/craterdog/go-syntax-notation/v5"
	osx "os"
	reg "regexp"
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
	remakeDirectory(directory + packageName)
	var generator = gen.PackageGenerator()
	var astSynthesizer = gen.AstSynthesizer(syntax)
	var source = generator.GeneratePackage(
		moduleName,
		wikiPath,
		packageName,
		astSynthesizer,
	)
	writeFile(filename, source)
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
		var classSynthesizer = gen.ClassSynthesizer(astModel, className)
		var source = generator.GenerateClass(
			moduleName,
			packageName,
			className,
			classSynthesizer,
		)
		var filename = directory + packageName + "/" + className + ".go"
		writeFile(filename, source)
	}
}

func generateFormatter(
	moduleName string,
	directory string,
	syntax not.SyntaxLike,
) {
	var generator = gen.ProcessorGenerator()
	var packageName = "grammar"
	var className = "formatter"
	var formatterSynthesizer = gen.FormatterSynthesizer(syntax)
	var source = generator.GenerateClass(
		moduleName,
		packageName,
		className,
		formatterSynthesizer,
	)
	var filename = directory + packageName + "/" + className + ".go"
	source = replaceMethodicalMethods(filename, source)
	writeFile(filename, source)
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
		grammarSynthesizer,
	)
	writeFile(filename, source)
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
	var filename = directory + "Module.go"
	fmt.Printf("  Generating %v...\n", filename)
	var packages = []string{
		"ast",
		"grammar",
	}
	var models = col.Catalog[string, mod.ModelLike]()
	for _, packageName := range packages {
		var filename = directory + packageName + "/Package.go"
		var source = readFile(filename)
		var model = mod.ParseSource(source)
		models.SetValue(packageName, model)
	}
	var generator = gen.ModuleGenerator()
	var moduleSynthesizer = gen.ModuleSynthesizer(models)
	var source = generator.GenerateModule(
		moduleName,
		wikiPath,
		moduleSynthesizer,
	)
	source = replaceGlobalFunctions(filename, source)
	writeFile(filename, source)
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
		parserSynthesizer,
	)
	var filename = directory + packageName + "/" + className + ".go"
	writeFile(filename, source)
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
		processorSynthesizer,
	)
	var filename = directory + packageName + "/" + className + ".go"
	writeFile(filename, source)
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
		scannerSynthesizer,
	)
	var filename = directory + packageName + "/" + className + ".go"
	writeFile(filename, source)
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
		tokenSynthesizer,
	)
	var filename = directory + packageName + "/" + className + ".go"
	writeFile(filename, source)
}

func generateValidator(
	moduleName string,
	directory string,
	syntax not.SyntaxLike,
) {
	var generator = gen.ProcessorGenerator()
	var packageName = "grammar"
	var className = "validator"
	var validatorSynthesizer = gen.ValidatorSynthesizer(syntax)
	var source = generator.GenerateClass(
		moduleName,
		packageName,
		className,
		validatorSynthesizer,
	)
	var filename = directory + packageName + "/" + className + ".go"
	source = replaceMethodicalMethods(filename, source)
	writeFile(filename, source)
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
		visitorSynthesizer,
	)
	var filename = directory + packageName + "/" + className + ".go"
	writeFile(filename, source)
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

func parseModel(
	source string,
) mod.ModelLike {
	var model = mod.ParseSource(source)
	mod.ValidateModel(model)
	return model
}

func parseSyntax(syntaxFile string) not.SyntaxLike {
	fmt.Printf("  Parsing %v...\n", syntaxFile)
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

func remakeDirectory(
	directory string,
) {
	var err = osx.RemoveAll(directory)
	if err != nil {
		panic(err)
	}
	err = osx.MkdirAll(directory, 0755)
	if err != nil {
		panic(err)
	}
}

func replaceGlobalFunctions(
	filename string,
	generatedSource string,
) string {
	if !pathExists(filename) {
		return generatedSource
	}
	var matcher = reg.MustCompile(
		`// GLOBAL FUNCTIONS(.|\r?\n)*`,
	)
	var originalSource = readFile(filename)
	var originalFunctions = matcher.FindString(originalSource)
	var generatedFunctions = matcher.FindString(generatedSource)
	generatedSource = sts.ReplaceAll(
		generatedSource,
		generatedFunctions,
		originalFunctions,
	)
	return generatedSource
}

func replaceMethodicalMethods(
	filename string,
	generatedSource string,
) string {
	if !pathExists(filename) {
		return generatedSource
	}
	var matcher = reg.MustCompile(
		`// Methodical Methods(.|\r?\n)+// PROTECTED INTERFACE`,
	)
	var originalSource = readFile(filename)
	var originalFunctions = matcher.FindString(originalSource)
	var generatedFunctions = matcher.FindString(generatedSource)
	generatedSource = sts.ReplaceAll(
		generatedSource,
		generatedFunctions,
		originalFunctions,
	)
	return generatedSource
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
