echo "Removing the old test results..."
rm -rf ./tst
mkdir -p ./tst/
echo

echo "Importing the following syntax notations:"
echo "  go-syntax-notation"
mkdir -p ./tst/go-syntax-notation/v5/
cp ../go-syntax-notation/v5/Syntax.cdsn ./tst/go-syntax-notation/v5/

echo "  go-class-model"
mkdir -p ./tst/go-class-model/v5/
cp ../go-class-model/v5/Syntax.cdsn ./tst/go-class-model/v5/
echo

echo "Done."
