build:
	@go build -o bin/go-hypefast cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/go-hypefast