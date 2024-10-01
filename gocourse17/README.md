# Example of REST and gRPC servers

Run REST and gRPC servers:

```sh
go run cmd/server/*.go
```

Run the GRPC client:
```sh
go run cmd/client/*.go
```

## gRPC

Regenerate GRPC:

```sh
protoc --go_out=./internal/grpc/grpcapi --go_opt=paths=source_relative --go-grpc_out=./internal/grpc/grpcapi --go-grpc_opt=paths=source_relative api.proto
```

Ensure the `protoc` and Go plugins for the protocol compiler are installed. See [instructions](https://grpc.io/docs/languages/go/quickstart/#prerequisites).
