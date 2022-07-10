DESTDIR :=
PREFIX := /usr/local

.PHONY: all
all:
	go build -o bin/tt src/*.go

.PHONY: install
install:
	install -d $(DESTDIR)$(PREFIX)/bin
	install -d $(DESTDIR)$(PREFIX)/share/man/man1
	install -m755 bin/tt $(DESTDIR)$(PREFIX)/bin
	install -m644 tt.1.gz $(DESTDIR)$(PREFIX)/share/man/man1

.PHONY: uninstall
uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/tt
	rm -f $(DESTDIR)$(PREFIX)/share/man/man1/tt.1.gz

.PHONY: assets
assets:
	python3 ./scripts/themegen.py
	./scripts/pack themes/ words/ quotes/ > src/packed.go
	pandoc -s -t man -o - man.md|gzip > tt.1.gz

.PHONY: rel
rel:
	GOOS=darwin GOARCH=amd64 go build -o bin/tt-osx src/*.go
	GOOS=windows GOARCH=amd64 go build -o bin/tt.exe src/*.go
	GOOS=linux GOARCH=amd64 go build -o bin/tt-linux src/*.go
	GOOS=linux GOARCH=arm go build -o bin/tt-linux_arm src/*.go
	GOOS=linux GOARCH=arm64 go build -o bin/tt-linux_arm64 src/*.go
