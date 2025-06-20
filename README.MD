# Go-i18n
通过 gRPC 服务来提供和管理微服务的多语言应用。

1. 复制仓库到你的项目下。
 ```SHELL
   git clone https://github.com/nicksnyder/go-i18n.git
   cd go-i18n
   go mod tidy
``` 
2. 生成protobuf文件
```SHELL
   cd proto
   protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative i18n.proto
```
3. 运行服务
```SHELL
   go run main.go
```