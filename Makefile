default: test

test:
	@go test -cover -race ./...