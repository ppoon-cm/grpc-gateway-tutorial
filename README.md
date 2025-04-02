# grpc-gateway-tutorial

## Overview

This is a basic Hello World gRPC-Gateway repository.  It derives from the tutorial at https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/introduction/.

## Quick Start

- `docker run -it --name grpcgw golang:latest /bin/bash`
- `mkdir /app ; cd /app`
- `git clone https://github.com/ppoon-cm/grpc-gateway-tutorial.git`
- `cd grpc-gateway-tutorial`
- `go main.go`
- Follow the instructions in the `Make HTTP and gRPC Requests` section below to make HTTP and gRPC requests.

## Installation and Setup for Development

- Go 1.24+

- Install dependencies

    ```
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    go install google.golang.org/grpc/reflection
    ```

- Install Command-Line Tools
  - **grpcurl** - Make gRPC calls

    ```
    go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
    ```
  - **buf** - A developer tool that enables building and management of Protobuf APIs

    ```
    BIN="/usr/local/bin" && \
    VERSION="1.50.0" && \
    curl -sSL \
    "https://github.com/bufbuild/buf/releases/download/v${VERSION}/buf-$(uname -s)-$(uname -m)" \
    -o "${BIN}/buf" && \
    chmod +x "${BIN}/buf"
    ```

## Usage

### Start the gRPC-Gateway

```
go run cmd/main.go
```

### Make HTTP and gRPC Requests

#### HTTP
```
curl -X POST -k http://localhost:8080/v1/example/echo -d '{"name": " hello"}'
```

#### gPRC
```
grpcurl -plaintext -d '{"name": "Alice"}' 0.0.0.0:5566 helloworld.Greeter/SayHello
```

# Notes

## Modification of gRPC-Gateway Tutorial

This repository derives from the gRPC-Gateway tutorial at https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/introduction/.  If you follow the instructions from there step-by-step, it will not work. It is extremely outdated, and is not even consistent with the repo that it refers to for the full source code at https://github.com/iamrajiv/helloworld-grpc-gateway.

### Generating Stubs

- I had to run `buf config migrate` on the sample `buf.yaml` and `buf.gen.yaml` files to upgrade the config files to `v2`.
- Had to replace the `out` field value of the plugins to `gen/go`.
- In the `proto/hello_world.proto` file, I had to insert the line:
    ```
    option go_package = "github.com/myuser/myrepo/proto";
    ```

## Swagger UI (OpenAPI)

Following instructions at https://blainsmith.com/articles/go-grpc-gateway-openapi/, we `git clone https://github.com/swagger-api/swagger-ui` and copied the entire `/dist` directory into the `/third_party/swagger-ui` directory. Initial attempts to use `git sparse-checkout` and `git submodule` did not yield any discernable benefits.
