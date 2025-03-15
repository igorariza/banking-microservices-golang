BUILDPATH=$(CURDIR)
API_NAME=transaction-grpc
API_VERSION=1.0.0
API_PORT=8080
API_PORT_GRPC=50051

ctl:
	@cd ./rpc && goctl rpc protoc  transaction_api.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.

go-ctl:
	@cd ./rpc && go run transactionapi.go -f etc/transactionapi.yaml

coverage:
	@echo "Coverfile... required 60% coverage"
	@go test ./... --coverprofile coverfile_out >> /dev/null
	@go tool cover -func coverfile_out
	@go tool cover -func coverfile_out | grep total | awk '{print substr($$3, 1, length($$3)-1)}' > coverage.txt

test:
	@echo "Running tests..."
	@go test -v --coverprofile=coverage.out ./... ./...

.PHONY: test