.PHONY: proto

proto: 
	@ PATH=$(PATH):$(shell go env GOPATH)/bin protoc \
	--go_out=.      --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	--go-grpc_opt=require_unimplemented_servers=false \
	proto/*.proto

test:
	go test -v -cover ./...