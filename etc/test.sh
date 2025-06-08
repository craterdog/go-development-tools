echo "Removing the old binaries..."
rm -rf ./bin
mkdir -p ./bin/
echo "Done."
echo

echo "Compiling the following tools:"
echo "  create-syntax"
go build -o ./bin/ ./src/create-syntax
echo "  format-syntax"
go build -o ./bin/ ./src/format-syntax
echo "  create-model"
go build -o ./bin/ ./src/create-model
echo "  format-model"
go build -o ./bin/ ./src/format-model
echo "  generate-classes"
go build -o ./bin/ ./src/generate-classes
echo "  generate-module"
go build -o ./bin/ ./src/generate-module
echo "  generate-project"
go build -o ./bin/ ./src/generate-project
echo "Done."
echo

directory=./tst/go-syntax-notation
moduleName=github.com/craterdog/${directory:6}/v7
wikiPath=https://github.com/craterdog/${directory:6}/wiki
bin/format-syntax ${directory}/v7/syntax.cdsn
echo
bin/generate-project ${moduleName} ${wikiPath} ${directory}/v7/
echo
bin/generate-module ${moduleName} ${wikiPath} ${directory}/v7/ ast grammar
echo
cd ${directory}/v7
gofmt -w . >/dev/null
cat <<EOF > go.mod
module ${moduleName}

go 1.23
EOF
go mod tidy >/dev/null 2>&1
golangci-lint run
cd - >/dev/null

directory=./tst/go-class-model
moduleName=github.com/craterdog/${directory:6}/v7
wikiPath=https://github.com/craterdog/${directory:6}/wiki
bin/format-syntax ${directory}/v7/syntax.cdsn
echo
bin/generate-project ${moduleName} ${wikiPath} ${directory}/v7/
echo
bin/generate-module ${moduleName} ${wikiPath} ${directory}/v7/ ast grammar
echo
cd ${directory}/v7
gofmt -w . >/dev/null
cat <<EOF > go.mod
module ${moduleName}

go 1.23
EOF
go mod tidy >/dev/null 2>&1
golangci-lint run
cd - >/dev/null

directory=./tst/go-component-framework
moduleName=github.com/craterdog/${directory:6}/v7
wikiPath=https://github.com/craterdog/${directory:6}/wiki
bin/format-model ${directory}/v7/agent/package_api.go
echo
bin/format-model ${directory}/v7/collection/package_api.go
echo
bin/generate-classes ${moduleName} ${directory}/v7/ element
echo
bin/generate-classes ${moduleName} ${directory}/v7/ collection
echo
bin/generate-classes ${moduleName} ${directory}/v7/ agent
echo
bin/generate-module ${moduleName} ${wikiPath} ${directory}/v7/ element collection agent
echo
cd ${directory}/v7
gofmt -w . >/dev/null
cat <<EOF > go.mod
module ${moduleName}

go 1.24
EOF
go mod tidy >/dev/null 2>&1
golangci-lint run
cd - >/dev/null

directory=./tst/go-example-project
moduleName=github.com/craterdog/${directory:6}/v2
wikiPath=https://github.com/craterdog/${directory:6}/wiki
packageName=example
bin/create-syntax ${directory}/v2
echo
bin/format-syntax ${directory}/v2/syntax.cdsn
echo
bin/create-model ${moduleName} ${wikiPath} ${packageName} ${directory}/v2
echo
bin/format-model ${directory}/v2/example/package_api.go
echo
bin/generate-classes ${moduleName} ${directory}/v2/ example
echo
