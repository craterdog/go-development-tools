echo "Importing the following modules:"
rm -rf ./tst
mkdir -p ./tst/

module=go-syntax-notation/v7
echo "  ${module}"
mkdir -p ./tst/${module}/
cp -r ../${module}/ ./tst/${module}/

module=go-class-model/v7
echo "  ${module}"
mkdir -p ./tst/${module}/
cp -r ../${module}/ ./tst/${module}/

module=go-component-framework/v7
echo "  ${module}"
mkdir -p ./tst/${module}/
cp -r ../${module}/ ./tst/${module}/
rm ./tst/${module}/module_test.go
#rm -r ./tst/${module}/collection
rm -f ./tst/${module}/.goignore

echo "Done."
