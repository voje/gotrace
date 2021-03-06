.PHONY: build test

build:
	go build -o ./build/gotrace

test: build
	./build/gotrace -w ./testdata/main.go
	go run ./testdata/main.go
	mv ./testdata/main.go.gotrace.orig ./testdata/main.go
