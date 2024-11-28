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
echo

echo "Done."
