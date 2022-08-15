# Demo 6

Using buf.build.

- `brew install protoc`
  - Well Known types will be here: `/opt/homebrew/include/google/protobuf`
- `brew install proto-gen-go`
- `brew install bufbuild/buf/buf`

```bash
buf build

buf build --exclude-source-info -o -#format=json | jq '.file[] | .package'

# Add foo package outside of the v1 structure

buf lint

buf breaking
```
