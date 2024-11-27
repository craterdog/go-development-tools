echo "Removing the old binaries..."
rm -rf ./bin
mkdir -p ./bin/
echo

echo "Compiling the following tools:"
echo "	format-model"
go build -o ./bin/ ./src/format-model

echo "	format-syntax"
go build -o ./bin/ ./src/format-syntax

echo "	generate-classes"
go build -o ./bin/ ./src/generate-classes

echo "	generate-module"
go build -o ./bin/ ./src/generate-module

echo "	generate-packages"
go build -o ./bin/ ./src/generate-packages
echo

echo "Done."
