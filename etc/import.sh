echo "Removing the old test results..."
rm -rf ./tst
mkdir -p ./tst/
echo

echo "Importing the following modules:"
echo "  go-syntax-notation"
mkdir -p ./tst/go-syntax-notation/v5/
cp -r ../go-syntax-notation/v5/ ./tst/go-syntax-notation/v5/

echo "  go-class-model"
mkdir -p ./tst/go-class-model/v5/
cp -r ../go-class-model/v5/ ./tst/go-class-model/v5/
echo

echo "Done."
