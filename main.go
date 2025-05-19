package main

import (
	"flag"
	"fmt"
	"i18n-service/config"
	"i18n-service/proto"
	"i18n-service/rpc"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "."
	}
	inf := flag.String("c", configPath, "config path")
	flag.Parse()
	if *inf == "" {
		log.Fatalf("config path is required")
	}
	// 加载配置文件
	cfg, err := config.LoadConfig(*inf)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	// Create a new ConfigManager
	configManager, err := config.NewConfigManager(&cfg.Apollo)
	if err != nil {
		fmt.Println("Error creating ConfigManager:", err)
		return
	}
	configManager.Start()
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
	rpcServer := rpc.NewCulturesRpc(configManager)

	proto.RegisterI18NServiceServer(grpcServer, rpcServer)

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
