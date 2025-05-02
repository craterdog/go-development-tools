<img src="https://craterdog.com/images/CraterDog.png" width="50%">

## Go Development Tools

### Overview
This project provides a set of Go based command-line tools that can be used to
do the following:
 * Validate and reformat a Go class model declared in a `package_api.go` file;
 * Validate and reformat a language grammar defined in a `syntax.cdsn`file;
 * Generate `ast/package_api.go` and `grammar/package_api.go` files based on a `syntax.cdsn`
   file;
 * And generate a `module_api.go` file based on a set of `package_api.go` files.

### Quick Links
For more information on this project and the projects on which it depends, click
on the following links:
 * [project documentation](https://github.com/craterdog/go-development-tools/wiki)
 * [class model](https://github.com/craterdog/go-class-model/wiki)
 * [syntax notation](https://github.com/craterdog/go-syntax-notation/wiki)
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
