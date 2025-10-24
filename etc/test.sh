etc/build.sh

rm -rf ./tst/
mkdir ./tst/

module=go-example-project/v2
directory=./tst/${module}
moduleName=github.com/craterdog/${module}
wikiPath=https://github.com/craterdog/go-example-project/wiki
mkdir -p ${directory}
bin/create-syntax ${directory}
echo
bin/format-syntax ${directory}/syntax.cdsn
echo
bin/generate-project ${moduleName} ${wikiPath} ${directory}
cd ${directory}
gofmt -w . >/dev/null
echo "Updating the dependencies..."
cat <<EOF > go.mod
module ${moduleName}

go 1.25
EOF
go mod tidy >/dev/null 2>&1
golangci-lint run
cd - >/dev/null
echo "Done."
echo

module=go-example-package/v2
directory=./tst/${module}
moduleName=github.com/craterdog/${module}
wikiPath=https://github.com/craterdog/go-example-package/wiki
packageName=example
mkdir -p ${directory}
bin/create-model ${moduleName} ${wikiPath} ${packageName} ${directory}
echo
bin/format-model ${directory}/example/package_api.go
echo
bin/generate-classes ${moduleName} ${directory} example
echo
bin/generate-module ${moduleName} ${wikiPath} ${directory} example
cd ${directory}
gofmt -w . >/dev/null
echo "Updating the dependencies..."
cat <<EOF > go.mod
module ${moduleName}

go 1.25
EOF
go mod tidy >/dev/null 2>&1
golangci-lint run
cd - >/dev/null
echo "Done."
echo

