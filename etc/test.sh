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
bin/format-model ${directory}/v7/agents/package_api.go
echo
bin/generate-classes ${moduleName} ${directory}/v7/ agents
echo
bin/format-model ${directory}/v7/collections/package_api.go
echo
#bin/generate-classes ${moduleName} ${directory}/v7/ collections
#echo
bin/format-model ${directory}/v7/elements/package_api.go
echo
bin/generate-classes ${moduleName} ${directory}/v7/ elements
echo
bin/format-model ${directory}/v7/ranges/package_api.go
echo
#bin/generate-classes ${moduleName} ${directory}/v7/ ranges
#echo
bin/format-model ${directory}/v7/strings/package_api.go
echo
bin/generate-classes ${moduleName} ${directory}/v7/ strings
echo
bin/generate-module ${moduleName} ${wikiPath} ${directory}/v7/ agents collections elements ranges strings
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
