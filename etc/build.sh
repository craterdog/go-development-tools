echo "Removing the old binaries..."
rm -rf ./bin
mkdir -p ./bin/
echo "Done."
echo

echo "Updating the dependencies..."
head -4 go.mod >go.deleteme
mv go.deleteme go.mod
go mod tidy >/dev/null 2>&1
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
