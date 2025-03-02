build:
	@buf generate

run:
	@go run go/grpc-uberfx/cmd/main.go grpc