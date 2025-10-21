echo "Importing the following modules:"
rm -rf ./tst
mkdir -p ./tst/

module=go-syntax-notation/v8
echo "  ${module}"
mkdir -p ./tst/${module}/
cp -r ../${module}/ ./tst/${module}/

module=go-class-model/v8
echo "  ${module}"
mkdir -p ./tst/${module}/
cp -r ../${module}/ ./tst/${module}/

module=go-collection-framework/v8
echo "  ${module}"
mkdir -p ./tst/${module}/
cp -r ../${module}/ ./tst/${module}/
rm ./tst/${module}/module_test.go  # The unit tests won't compile.

echo "Done."
