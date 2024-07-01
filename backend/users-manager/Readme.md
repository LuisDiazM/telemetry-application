```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

Generación protos para users
```
protoc --go_out=. --go_opt=paths=source_relative                                                                  --go-grpc_out=. --go-grpc_opt=paths=source_relative                                                                  infraestructure/server/protos/users/user.proto
```