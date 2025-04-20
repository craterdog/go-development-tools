echo "Importing the following modules:"
rm -rf ./tst
mkdir -p ./tst/

module=go-syntax-notation/v6
echo "  ${module}"
mkdir -p ./tst/${module}/
cp -r ../${module}/ ./tst/${module}/

module=go-class-model/v5
echo "  ${module}"
mkdir -p ./tst/${module}/
cp -r ../${module}/ ./tst/${module}/

module=go-collection-framework/v5
echo "  ${module}"
mkdir -p ./tst/${module}/
cp -r ../${module}/ ./tst/${module}/
rm ./tst/${module}/Module_test.go
rm ./tst/${module}/agent/Package_test.go
rm ./tst/${module}/collection/Package_test.go
rm -f ./tst/${module}/.goignore

echo "Done."
