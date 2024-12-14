directory=./tst/go-syntax-notation
moduleName=github.com/craterdog/${directory:6}/v5
wikiPath=https://github.com/craterdog/${directory:6}/wiki
bin/format-syntax ${directory}/v5/Syntax.cdsn
echo
bin/generate-packages ${moduleName} ${wikiPath} ${directory}/v5/
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

directory=./tst/go-class-model
moduleName=github.com/craterdog/${directory:6}/v5
wikiPath=https://github.com/craterdog/${directory:6}/wiki
bin/format-syntax ${directory}/v5/Syntax.cdsn
echo
bin/generate-packages ${moduleName} ${wikiPath} ${directory}/v5/
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
bin/format-model ${directory}/v5/collection/Package.go
echo
bin/generate-module ${moduleName} ${wikiPath} ${directory}/v5/ agent collection
echo
bin/generate-classes ${moduleName} ${directory}/v5/ agent
bin/generate-classes ${moduleName} ${directory}/v5/ collection
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
