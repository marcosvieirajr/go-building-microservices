# Currency Serrvice

The currency service is a gRPC service provides up to date exchange rates and currency conversion capabilities.

## Building protos

To build the gRPC client and server interfaces, first install Protocol buffer compiler, `protoc`:

<https://grpc.io/docs/protoc-installation/>

### Linux

```bash
apt install -y protobuf-compiler
```

### Mac

```shell
brew install protobuf
```

Then install Go plugins for the protocol compiler:

<https://grpc.io/docs/languages/go/quickstart/>

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Update your PATH so that the protoc compiler can find the plugins:

```shell
export PATH="$PATH:$(go env GOPATH)/bin"
```

Then run the build commmand:

```shell
protoc \
    --go_out=.      --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    --go-grpc_opt=require_unimplemented_servers=false \
    proto/*.proto
```

Or run the Makefile targe:

```shell
make proto
```

## Testing

To test the system install `grpccurl` which is a command line tool which can interact with gRPC API's

<https://github.com/fullstorydev/grpcurl>

```shell
go install github.com/fullstorydev/grpcurl/cmd/grpcurl
```

### Or Homebrew (Mac)

```shell
brew install grpcurl
```

### List Services

```shell
$ grpcurl --plaintext localhost:9092 list
Currency
grpc.reflection.v1alpha.ServerReflection
```

### List Methods

```shell
$ grpcurl --plaintext localhost:9092 list Currency        
Currency.GetRate
```

### Method detail for GetRate

```shell
$ grpcurl --plaintext localhost:9092 describe Currency.GetRate
Currency.GetRate is a method:
rpc GetRate ( .RateRequest ) returns ( .RateResponse );
```

### RateRequest detail

```shell
$ grpcurl --plaintext localhost:9092 describe .RateRequest
RateRequest is a message:
message RateRequest {
  .Currencies base = 1;
  .Currencies destination = 2;
}
```

### Execute a request

```shell
$ grpcurl --plaintext -d '{"base":"GBP", "destination": "USD"}' localhost:9092 Currency.GetRate
{
  "rate": 0.5
}
```

## References

- <https://golangrepo.com/repo/grpc-grpc-go-go-distributed-systems>

- <https://developers.google.com/protocol-buffers/docs/reference/proto3-spec>

- <https://github.com/grpc/grpc-go/blob/master/README.md>

- <https://grpc.io/docs/languages/go>
