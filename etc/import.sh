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

module=go-code-generation/v5/
echo "  ${module}"
mkdir -p ./tst/${module}
cp ../${module}/Module.go ./tst/${module}/
mkdir -p ./tst/${module}/analyzer/
cp ../${module}/analyzer/Package.go ./tst/${module}/analyzer/
mkdir -p ./tst/${module}/synthesizer/
cp ../${module}/synthesizer/Package.go ./tst/${module}/synthesizer/
mkdir -p ./tst/${module}/generator/
cp ../${module}/generator/Package.go ./tst/${module}/generator/

echo "Done."
