# Demo 5

Package structure best practicies.

- `brew install protoc`
  - Well Known types will be here: `/opt/homebrew/include/google/protobuf`
- `brew install proto-gen-go`

```bash
protoc \
  -I=$(pwd) \
  -I=/opt/homebrew/include \
  --go_out=$(pwd) \
  --go_opt=paths=source_relative \
  $(find ~+ -type f -name '*.proto')
```
