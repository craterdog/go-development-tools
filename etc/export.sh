echo "Exporting the following modules:"

module=go-syntax-notation/v5/
echo "  ${module}"
rm -rf ../${module}
cp -r ./tst/${module} ../${module}

module=go-class-model/v5/
echo "  ${module}"
rm -rf ../${module}
cp -r ./tst/${module} ../${module}

module=go-collection-framework/v5/
echo "  ${module}"
cp ./tst/${module}/Module.go ../${module}/
cp ./tst/${module}/agent/Package.go ../${module}/agent/
cp ./tst/${module}/collection/Package.go ../${module}/collection/

module=go-code-generation/v5/
echo "  ${module}"
cp ./tst/${module}/Module.go ../${module}/
cp ./tst/${module}/analyzer/Package.go ../${module}/analyzer/
cp ./tst/${module}/synthesizer/Package.go ../${module}/synthesizer/
cp ./tst/${module}/generator/Package.go ../${module}/generator/

echo "Done."
