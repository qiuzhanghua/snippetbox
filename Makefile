default: clean build

clean:
	rm -rf snippetbox
	rm -rf bin/*

build:
	go build -o bin/ ./...

test:
	go test -v ./...

run:
	go run cmd/web/*.go
