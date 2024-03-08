version=0.1

.PHONY: build bindir clean
build: clean bindir
	CGO_ENABLED=0 go build -ldflags="-s -w -X 'main.majorVersion=${version}'" -o ./bin/gc

bindir:
	mkdir ./bin

clean:
	rm -rf ./bin