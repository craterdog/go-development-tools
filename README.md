<img src="https://craterdog.com/images/CraterDog.png" width="50%">

## Go Development Tools

### Overview
This project provides a set of Go based command-line tools that can be used to
do the following:
 * validate and reformat a language grammar defined in a `Syntax.cdsn`file;
 * validate and reformat a Go class model declared in a `Package.go` file;
 * generate `ast/Package.go` and `grammar/Package.go` files based on a `Syntax.cdsn`
   file;
 * and generate a `Module.go` file based on a set of `Package.go` files.

### Quick Links
For more information on this project click on the following links:
 * [project documentation](https://github.com/craterdog/go-development-tools/wiki)
 * [syntax notation](https://github.com/craterdog/go-syntax-notation/wiki)
 * [class model](https://github.com/craterdog/go-class-model/wiki)
 * [code generation](https://github.com/craterdog/go-code-generation/wiki)

### Getting Started
To install the language grammar tools do the following from a terminal window:
```
$ git clone git@github.com:craterdog/go-development-tools.git
$ cd go-development-tools/
$ etc/build.sh
$ ls
LICENSE		bin		go.mod		src
README.md	etc		go.sum		tst
```

<H5 align="center"> Copyright © 2009 - 2024  Crater Dog Technologies™. All rights reserved. </H5>
