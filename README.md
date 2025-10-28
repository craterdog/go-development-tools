[<img src="https://craterdog.com/images/CraterDog.png" width="30%">](https://craterdog.com)

## Go Development Tools

### Overview
This project provides a set of Go based command-line tools that can be used to
do the following:
 * Validate and reformat a language grammar defined in a `syntax.cdsn`file;
 * Generate a fully working project based on the `syntax.cdsn` file;
 * Validate and reformat a Go class model declared in a `package_api.go` file;
 * Generate class skeletons for each class defined in the `package_api.go` file;
 * And generate a `module_api.go` file based on the set of `package_api.go`
   files defined for a project.

### Quick Links
For more information on this project and the projects on which it depends, click
on the following links:
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

### Contributing
Project contributors are always welcome. Check out the contributing guidelines
[here](https://github.com/craterdog/go-development-tools/blob/main/.github/CONTRIBUTING.md).

<H5 align="center"> Copyright © 2009-2025. Crater Dog Technologies™. All rights reserved. </H5>
