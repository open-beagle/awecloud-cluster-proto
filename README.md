# awecloud-cluster-proto

## proto

```bash
# install
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

# update
protoc \
  --go_out=. \
  --go_opt=paths=source_relative \
  --go-grpc_out=. \
  --go-grpc_opt=paths=source_relative \
  cluster_registry.proto
```

## tag

```bash
git tag v0.6.1 -f
git push origin v0.6.1 -f
```
