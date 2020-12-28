all:
	go build -o bin/tt *.go
install:
	install -m755 bin/tt /usr/local/bin
assets:
	python3 tools/themegen.py | gofmt > generatedThemes.go 
rel:
	GOOS=darwin GOARCH=amd64 go build -o bin/tt-osx *.go
	GOOS=windows GOARCH=amd64 go build -o bin/tt.exe *.go
	GOOS=linux GOARCH=amd64 go build -o bin/tt-linux *.go
	GOOS=linux GOARCH=386 go build -o bin/tt-linux_386 *.go
