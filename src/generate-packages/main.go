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
	var moduleName, wikiPath, directory, force = retrieveArguments()
	var syntax = validateSyntaxFile(directory)
	generateAstPackage(moduleName, wikiPath, directory, syntax, force)
	generateGrammarPackage(moduleName, wikiPath, directory, syntax, force)
	generateModuleFile(moduleName, wikiPath, directory, force)
	fmt.Println("Done.")
}

func generateAstPackage(
	moduleName string,
	wikiPath string,
	directory string,
	syntax not.SyntaxLike,
	force bool,
) {
	fmt.Println("  Generating the /ast/Package.go file...")
	remakeDirectory(directory, force)
	var generator = gen.PackageGenerator()
	var packageName = "ast"
	var astSynthesizer = gen.AstSynthesizer(syntax)
	var source = generator.GeneratePackage(
		moduleName,
		wikiPath,
		packageName,
		astSynthesizer,
	)
	var filename = directory + packageName + "/Package.go"
	writeFile(filename, source, force)
	var astModel = parseModel(source)
	generateClasses(moduleName, directory, packageName, astModel, force)
}

func generateClasses(
	moduleName string,
	directory string,
	packageName string,
	astModel mod.ModelLike,
	force bool,
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
		writeFile(filename, source, force)
	}
}

func generateFormatter(
	moduleName string,
	directory string,
	model mod.ModelLike,
	force bool,
) {
	var generator = gen.ClassGenerator()
	var packageName = "grammar"
	var className = "formatter"
	var formatterSynthesizer = gen.FormatterSynthesizer(model)
	var source = generator.GenerateClass(
		moduleName,
		packageName,
		className,
		formatterSynthesizer,
	)
	var filename = directory + packageName + "/" + className + ".go"
	source = replaceMethodicalMethods(filename, source, force)
	writeFile(filename, source, force)
}

func generateGrammarPackage(
	moduleName string,
	wikiPath string,
	directory string,
	syntax not.SyntaxLike,
	force bool,
) {
	fmt.Println("  Generating the /grammar/Package.go file...")
	remakeDirectory(directory, force)
	var generator = gen.PackageGenerator()
	var packageName = "grammar"
	var grammarSynthesizer = gen.AstSynthesizer(syntax)
	var source = generator.GeneratePackage(
		moduleName,
		wikiPath,
		packageName,
		grammarSynthesizer,
	)
	var filename = directory + packageName + "/Package.go"
	writeFile(filename, source, force)
	var grammarModel = parseModel(source)
	generateFormatter(moduleName, directory, grammarModel, force)
	generateParser(moduleName, directory, grammarModel, force)
	generateProcessor(moduleName, directory, grammarModel, force)
	generateScanner(moduleName, directory, grammarModel, force)
	generateToken(moduleName, directory, grammarModel, force)
	generateValidator(moduleName, directory, grammarModel, force)
	generateVisitor(moduleName, directory, grammarModel, force)
}

func generateModuleFile(
	moduleName string,
	wikiPath string,
	directory string,
	force bool,
) {
	fmt.Println("  Generating the /Module.go file...")
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
	var filename = directory + "Module.go"
	source = replaceGlobalFunctions(filename, source, force)
	writeFile(filename, source, force)
}

func generateParser(
	moduleName string,
	directory string,
	model mod.ModelLike,
	force bool,
) {
	var generator = gen.ClassGenerator()
	var packageName = "grammar"
	var className = "parser"
	var parserSynthesizer = gen.ParserSynthesizer(model)
	var source = generator.GenerateClass(
		moduleName,
		packageName,
		className,
		parserSynthesizer,
	)
	var filename = directory + packageName + "/" + className + ".go"
	writeFile(filename, source, force)
}

func generateProcessor(
	moduleName string,
	directory string,
	model mod.ModelLike,
	force bool,
) {
	var generator = gen.ClassGenerator()
	var packageName = "grammar"
	var className = "processor"
	var processorSynthesizer = gen.ProcessorSynthesizer(model)
	var source = generator.GenerateClass(
		moduleName,
		packageName,
		className,
		processorSynthesizer,
	)
	var filename = directory + packageName + "/" + className + ".go"
	writeFile(filename, source, force)
}

func generateScanner(
	moduleName string,
	directory string,
	model mod.ModelLike,
	force bool,
) {
	var generator = gen.ClassGenerator()
	var packageName = "grammar"
	var className = "scanner"
	var scannerSynthesizer = gen.ScannerSynthesizer(model)
	var source = generator.GenerateClass(
		moduleName,
		packageName,
		className,
		scannerSynthesizer,
	)
	var filename = directory + packageName + "/" + className + ".go"
	writeFile(filename, source, force)
}

func generateToken(
	moduleName string,
	directory string,
	model mod.ModelLike,
	force bool,
) {
	var generator = gen.ClassGenerator()
	var packageName = "grammar"
	var className = "token"
	var tokenSynthesizer = gen.TokenSynthesizer(model)
	var source = generator.GenerateClass(
		moduleName,
		packageName,
		className,
		tokenSynthesizer,
	)
	var filename = directory + packageName + "/" + className + ".go"
	writeFile(filename, source, force)
}

func generateValidator(
	moduleName string,
	directory string,
	model mod.ModelLike,
	force bool,
) {
	var generator = gen.ClassGenerator()
	var packageName = "grammar"
	var className = "validator"
	var validatorSynthesizer = gen.ValidatorSynthesizer(model)
	var source = generator.GenerateClass(
		moduleName,
		packageName,
		className,
		validatorSynthesizer,
	)
	var filename = directory + packageName + "/" + className + ".go"
	source = replaceMethodicalMethods(filename, source, force)
	writeFile(filename, source, force)
}

func generateVisitor(
	moduleName string,
	directory string,
	model mod.ModelLike,
	force bool,
) {
	var generator = gen.ClassGenerator()
	var packageName = "grammar"
	var className = "visitor"
	var visitorSynthesizer = gen.VisitorSynthesizer(model)
	var source = generator.GenerateClass(
		moduleName,
		packageName,
		className,
		visitorSynthesizer,
	)
	var filename = directory + packageName + "/" + className + ".go"
	writeFile(filename, source, force)
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
	if !pathExists(syntaxFile) {
		fmt.Printf(
			"The syntax file %v does not exist, aborting...",
			syntaxFile,
		)
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

func replaceGlobalFunctions(
	filename string,
	generatedSource string,
	force bool,
) string {
	if !pathExists(filename) || !force {
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
	force bool,
) string {
	if !pathExists(filename) || !force {
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
	force bool,
) {
	if len(osx.Args) < 4 {
		var tool = getTool()
		fmt.Printf("  Usage: %v <moduleName> <wikiPath> <directory> [force]\n", tool)
		osx.Exit(1)
	}
	moduleName = osx.Args[1]
	wikiPath = osx.Args[2]
	directory = osx.Args[3]
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	force = len(osx.Args) == 5
	return
}

func validateSyntaxFile(directory string) not.SyntaxLike {
	fmt.Println("  Validating the Syntax.cdsn file...")
	var syntaxFile = directory + "Syntax.cdsn"
	var syntax = parseSyntax(syntaxFile)
	not.ValidateSyntax(syntax)
	return syntax
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
