Name=codesync

install:
	go build -o ${Name} .
	mv ${Name} /usr/local/bin

build:
	mkdir bin
	GOOS=windows GOARCH=amd64 go build -o bin/${Name}\ OS=windows\ Arch=amd64.exe .
	GOOS=windows GOARCH=386 go build -o bin/${Name}\ OS=windows\ Arch=386.exe .
	GOOS=darwin GOARCH=amd64 go build -o bin/${Name}\ OS=darwin\ Arch=amd64 .
	GOOS=darwin GOARCH=arm64 go build -o bin/${Name}\ OS=darwin\ Arch=arm64 .
	GOOS=linux GOARCH=amd64 go build -o bin/${Name}\ OS=linux\ Arch=amd64 .
	GOOS=linux GOARCH=386 go build -o bin/${Name}\ OS=linux\ Arch=386 .
	GOOS=linux GOARCH=arm64 go build -o bin/${Name}\ OS=linux\ Arch=arm64 .
	GOOS=linux GOARCH=arm go build -o bin/${Name}\ OS=linux\ Arch=arm .