echo "Exporting the following modules:"

echo "  go-syntax-notation"
#rm -rf ../go-syntax-notation/v5
cp -r ./tst/go-syntax-notation/v5 ../go-syntax-notation/

echo "  go-class-model"
#rm -rf ../go-class-model/v5
cp -r ./tst/go-class-model/v5 ../go-class-model/
echo

echo "Done."
