all:
	go build -o bin/tt *.go
install:
	install -m755 bin/tt /usr/local/bin
rel:
	GOOS=darwin GOARCH=386 go build -o binaries/tt-osx_386 *.go
	GOOS=darwin GOARCH=amd64 go build -o binaries/tt-osx_amd64 *.go
	GOOS=windows GOARCH=386 go build -o binaries/tt-windows_386 *.go
	GOOS=windows GOARCH=amd64 go build -o binaries/tt-windows_amd64 *.go
	GOOS=linux GOARCH=386 go build -o binaries/tt-linux_386 *.go
	GOOS=linux GOARCH=amd64 go build -o binaries/tt-linux_amd64 *.go
