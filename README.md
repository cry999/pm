# PM

PM is the project management tool.

## build

```sh
export PATH=$(pwd)/bin:$PATH
go build -o bin/ ./cmd/tools/...
go generate ./...
go build -o bin/server ./cmd/server
```
