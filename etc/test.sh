for directory in `ls -d ./tst/go-*`; do
	moduleName=github.com/craterdog/${directory:6}/v5
	wikiPath=https://github.com/craterdog/${directory:6}/wiki
	bin/format-syntax ${directory}/v5/Syntax.cdsn
	echo
	bin/generate-packages ${moduleName} ${wikiPath} ${directory}/v5/
	echo
	cd ${directory}/v5
	gofmt -w . >/dev/null
	cat <<EOF > go.mod
module ${moduleName}

go 1.22
EOF
	go mod tidy >/dev/null 2>&1
	golangci-lint run
	cd - >/dev/null
done
