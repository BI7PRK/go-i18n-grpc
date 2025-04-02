package main

import (
	"flag"
	"fmt"
	"i18n-service/config"
	"i18n-service/proto"
	"i18n-service/rpc"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"xorm.io/xorm"
)

func main() {
	cfg, ex := config.LoadAppConfig()
	if ex != nil {
		panic(ex)
	}

	// 初始化 MySQL 和 xorm
	db, err := xorm.NewEngine("mysql", "root:123456@/culturei18n?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 定义命令行参数
	portPtr := flag.Int("p", 50001, "gRPC service port")
	flag.Parse()
	if *portPtr != cfg.Server.Port {
		cfg.Server.Port = *portPtr
	}
	port := fmt.Sprintf(":%d", cfg.Server.Port)
	// 初始化 gRPC 服务器
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	rpcServer := rpc.NewCulturesRpc()

	proto.RegisterI18NServiceServer(grpcServer, rpcServer)

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
