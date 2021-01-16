all:
	go build -o bin/tt *.go
install:
	install -m755 bin/tt /usr/local/bin
	install -m755 tt.1.gz /usr/share/man/man1
assets:
	python3 ./scripts/themegen.py
	./scripts/pack themes/ words/ quotes/ > packed.go
	pandoc -s -t man -o - man.md|gzip > tt.1.gz
rel:
	GOOS=darwin GOARCH=amd64 go build -o bin/tt-osx *.go
	GOOS=windows GOARCH=amd64 go build -o bin/tt.exe *.go
	GOOS=linux GOARCH=amd64 go build -o bin/tt-linux *.go
	GOOS=linux GOARCH=386 go build -o bin/tt-linux_386 *.go
