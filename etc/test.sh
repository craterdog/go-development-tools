wikiPath=https://github.com/craterdog/example.wiki

for directory in `ls -d ./tst/go-*/v*`; do
	moduleName=github.com/craterdog/${directory:6}
	bin/format-syntax $directory/Syntax.cdsn
	echo
	bin/generate-packages $moduleName $wikiPath $directory/ force
	echo
	cd $directory
	gofmt -s -w . >/dev/null
	cat <<EOF > go.mod
module $moduleName

go 1.22
EOF
	go mod tidy >/dev/null 2>&1
	golangci-lint run
	cd - >/dev/null
done
