package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/example/tdd-hello/internal/service"
	pb "github.com/example/tdd-hello/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ============================================
// TODO 步骤 2：实现服务器
// ============================================
// 
// 任务清单：
// 
// 1. 创建一个 gRPC 服务器
// 2. 注册 GreeterService
// 3. 监听指定端口（默认 50051）
// 4. 启动服务器
// 5. 优雅关闭（使用 context）
// 
// 完成后运行的命令：
//   go run ./cmd/server
//   预期输出：服务器启动成功，监听在 :50051
// 
// 然后运行测试：
//   go run ./cmd/server &  # 后台运行
//   go test ./cmd/server -v
//   预期结果：TestServer_StartsSuccessfully 通过

func main() {
	// TODO: 添加命令行参数，允许指定端口
	port := flag.Int("port", 50051, "gRPC server port")
	flag.Parse()

	// TODO: 创建监听
	// 使用 net.Listen("tcp", fmt.Sprintf(":%d", *port))
	
	// TODO: 创建 gRPC 服务器
	// 使用 grpc.NewServer()
	
	// TODO: 创建 GreeterService 实例
	// 使用 service.GreeterService{}
	
	// TODO: 注册服务
	// pb.RegisterGreeterServer(server, greeterService)
	
	// TODO: 打印服务器启动信息
	// log.Printf("服务器启动，监听端口 %d", *port)
	
	// TODO: 启动服务器
	// 使用 server.Serve(lis)
	
	// TODO: 添加优雅关闭处理
	// 使用 defer server.GracefulStop()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("无法监听端口 %d: %v", *port, err)
	}

	server := grpc.NewServer()
	gr := &service.GreeterService{}
	pb.RegisterGreeterServer(server, gr)

	log.Printf("服务器启动，监听端口 %d", *port)
	defer server.GracefulStop()

	if err := server.Serve(lis); err != nil {
		log.Fatalf("服务器启动失败：%v", err)
	}
}

// ============================================
// 服务器集成测试
// ============================================
// 将此代码放在 server_test.go 中测试服务器