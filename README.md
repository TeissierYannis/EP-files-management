[![codecov](https://codecov.io/gh/TeissierYannis/EP-files-management/graph/badge.svg?token=MIBICOTDXG)](https://codecov.io/gh/TeissierYannis/EP-files-management)

# EP-files-management

Build protoc for go
```bash
protoc --go_out=. --go-grpc_out=. proto/file_service.proto
```

Build protoc for NodeJS
```bash
protoc --js_out=import_style=commonjs,binary:. \
       --grpc-web_out=import_style=typescript,mode=grpcwebtext:. \
       proto/file_service.proto
```