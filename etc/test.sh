etc/build.sh
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
echo "Updating the dependencies..."
cat <<EOF > go.mod
module ${moduleName}

go 1.24
EOF
go mod tidy >/dev/null 2>&1
golangci-lint run
cd - >/dev/null
echo "Done."
echo

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
echo "Updating the dependencies..."
cat <<EOF > go.mod
module ${moduleName}

go 1.24
EOF
go mod tidy >/dev/null 2>&1
golangci-lint run
cd - >/dev/null
echo "Done."
echo

directory=./tst/go-component-framework
moduleName=github.com/craterdog/${directory:6}/v7
wikiPath=https://github.com/craterdog/${directory:6}/wiki
bin/format-model ${directory}/v7/agent/package_api.go
echo
bin/format-model ${directory}/v7/collection/package_api.go
echo
bin/generate-classes ${moduleName} ${directory}/v7/ element
echo
bin/generate-classes ${moduleName} ${directory}/v7/ string
echo
bin/generate-classes ${moduleName} ${directory}/v7/ agent
echo
bin/generate-module ${moduleName} ${wikiPath} ${directory}/v7/ element string collection agent
echo
cd ${directory}/v7
gofmt -w . >/dev/null
echo "Updating the dependencies..."
cat <<EOF > go.mod
module ${moduleName}

go 1.24
EOF
go mod tidy >/dev/null 2>&1
golangci-lint run
cd - >/dev/null
echo "Done."
echo

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
