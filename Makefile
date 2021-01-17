all:
	go build -o bin/tt src/*.go
install:
	install -m755 bin/tt /usr/local/bin
	install -m755 tt.1.gz /usr/share/man/man1
assets:
	python3 ./scripts/themegen.py
	./scripts/pack themes/ words/ quotes/ > src/packed.go
	pandoc -s -t man -o - man.md|gzip > tt.1.gz
rel:
	GOOS=darwin GOARCH=amd64 go build -o bin/tt-osx src/*.go
	GOOS=windows GOARCH=amd64 go build -o bin/tt.exe src/*.go
	GOOS=linux GOARCH=amd64 go build -o bin/tt-linux src/*.go
	GOOS=linux GOARCH=arm go build -o bin/tt-linux_arm src/*.go
	GOOS=linux GOARCH=arm64 go build -o bin/tt-linux_arm64 src/*.go
