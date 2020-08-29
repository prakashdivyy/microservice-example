# product-service

## How to Run

```
# Install dependency
go mod download

# Install protoc and protoc-gen-go
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}

# compile protobuf
protoc --proto_path=product --proto_path=../proto --go_out=plugins=grpc:rpc products.proto

# run server
go run server.go

# run client
go run client.go
```
