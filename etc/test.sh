module=github.com/craterdog/example

for directory in `ls -d ./tst/*/v*`; do
	bin/format-syntax $directory/Syntax.cdsn
	echo
done

#echo "Generating the go.mod file:"
#cd $output
#cat <<EOF > go.mod
#module github.com/craterdog/example
#
#go 1.22
#EOF
#go mod tidy
#echo
#
#echo "Running lint on the generated files:"
#golangci-lint run
#cd - >/dev/null
#echo
