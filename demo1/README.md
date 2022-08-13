# Demo 1

Compiling a single proto Go file.

- `brew install protobuf`
- `brew install protoc-gen-go`
- `protoc -I=$(pwd) --go_out=$(pwd) $(pwd)/payment.proto`
