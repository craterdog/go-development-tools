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
moduleName=github.com/craterdog/${directory:6}/v6
wikiPath=https://github.com/craterdog/${directory:6}/wiki
bin/format-syntax ${directory}/v6/Syntax.cdsn
echo
bin/generate-project ${moduleName} ${wikiPath} ${directory}/v6/
echo
bin/generate-module ${moduleName} ${wikiPath} ${directory}/v6/ ast grammar
echo
cd ${directory}/v6
gofmt -w . >/dev/null
cat <<EOF > go.mod
module ${moduleName}

go 1.23
EOF
go mod tidy >/dev/null 2>&1
golangci-lint run
cd - >/dev/null

directory=./tst/go-class-model
moduleName=github.com/craterdog/${directory:6}/v5
wikiPath=https://github.com/craterdog/${directory:6}/wiki
bin/format-syntax ${directory}/v5/Syntax.cdsn
echo
bin/generate-project ${moduleName} ${wikiPath} ${directory}/v5/
echo
bin/generate-module ${moduleName} ${wikiPath} ${directory}/v5/ ast grammar
echo
cd ${directory}/v5
gofmt -w . >/dev/null
cat <<EOF > go.mod
module ${moduleName}

go 1.23
EOF
go mod tidy >/dev/null 2>&1
golangci-lint run
cd - >/dev/null

directory=./tst/go-collection-framework
moduleName=github.com/craterdog/${directory:6}/v5
wikiPath=https://github.com/craterdog/${directory:6}/wiki
bin/format-model ${directory}/v5/agent/Package.go
echo
bin/format-model ${directory}/v5/collection/Package.go
echo
bin/generate-classes ${moduleName} ${directory}/v5/ agent
echo
bin/generate-classes ${moduleName} ${directory}/v5/ collection
echo
bin/generate-module ${moduleName} ${wikiPath} ${directory}/v5/ agent collection
echo
cd ${directory}/v5
gofmt -w . >/dev/null
cat <<EOF > go.mod
module ${moduleName}

go 1.23
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
bin/format-syntax ${directory}/v2/Syntax.cdsn
echo
bin/create-model ${moduleName} ${wikiPath} ${packageName} ${directory}/v2
echo
bin/format-model ${directory}/v2/example/Package.go
echo
