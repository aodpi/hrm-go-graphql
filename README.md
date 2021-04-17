# hrm-go-graphql
HRM grapqhl server with GO

## hrm-go-grpc

This server on go uses gRPC to communicate with hrm api server.

To generate the go code for gRPC protocol buffer files you need to:

* Make sure you have `protoc` installed. For more information consult this [documentation.](https://grpc.io/docs/protoc-installation/)

* Add your `.proto` files in the `src/grpc/proto` folder.

* Execute the following command (in `./src` cwd) to generate the required go source code to work with the gRPC client:
   
   ```bash
   $ protoc --go_out=./grpc/generated --go_opt=paths=source_relative \
            --go-grpc_out=./grpc/generated --go-grpc_opt=paths=source_relative \
            ./grpc/proto/*.proto
   ```
