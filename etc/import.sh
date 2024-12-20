echo "Importing the following modules:"
rm -rf ./tst
mkdir -p ./tst/

module=go-syntax-notation/v5/
echo "  ${module}"
mkdir -p ./tst/${module}
cp -r ../${module} ./tst/${module}

module=go-class-model/v5/
echo "  ${module}"
mkdir -p ./tst/${module}
cp -r ../${module} ./tst/${module}

module=go-collection-framework/v5/
echo "  ${module}"
mkdir -p ./tst/${module}
cp ../${module}/Module.go ./tst/${module}/
mkdir -p ./tst/${module}/agent/
cp ../${module}/agent/Package.go ./tst/${module}/agent/
mkdir -p ./tst/${module}/collection/
cp ../${module}/collection/Package.go ./tst/${module}/collection/

echo "Done."
