# Demo 6

Using buf.build.

- `brew install proto-gen-go`
- `brew install bufbuild/buf/buf`

```bash
# Validate everytrhing is setup correct
buf build

# Print detected packages
buf build --exclude-source-info -o -#format=json | jq '.file[] | .package'

# Add foo package outside of the v1 structure
buf lint

# Make a breaking change
buf breaking --against "../.git#branch=main,subdir=demo6"

# Generate code
buf generate
```
