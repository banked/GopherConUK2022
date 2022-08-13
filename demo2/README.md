# Demo

Importing packages.

- `brew install protoc`
  - Well Known types will be here: `/opt/homebrew/include/google/protobuf`
- `brew install proto-gen-go`
- `protoc -I=$(pwd) -I=/opt/homebrew/include --go_out=$(pwd) $(pwd)/payment.proto $(pwd)/amount.proto`
