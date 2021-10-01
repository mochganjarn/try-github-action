build:
	@go build
run:
	@go run main.go
test:
	@go test -v ./...
lint:
	@golangci-lint run