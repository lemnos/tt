all:
	go build -o bin/tt *.go
install:
	install -m755 bin/tt /usr/local/bin
