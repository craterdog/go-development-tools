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
#rm -rf ../${module}
cp ./tst/${module}/Module.go ../${module}/
echo

echo "Done."
