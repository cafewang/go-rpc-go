package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/example/tdd-hello/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ============================================
// TODO 步骤 3：实现客户端
// ============================================
// 
// 任务清单：
// 
// 1. 创建 gRPC 客户端连接
// 2. 连接到服务器（默认 localhost:50051）
// 3. 创建 GreeterClient
// 4. 调用 SayHello 方法
// 5. 打印响应
// 
// 完成后运行的命令：
//   go run ./cmd/client -name "世界"
//   预期输出：Hello, 世界
// 
// 然后运行测试：
//   go test ./cmd/client -v
//   预期结果：TestClient_CanConnectAndCall 通过

func main() {
	// 添加命令行参数
	name := flag.String("name", "World", "要问候的名字")
	addr := flag.String("addr", "localhost:50051", "服务器地址")
	flag.Parse()

	// TODO: 创建 gRPC 客户端连接
	// 使用 grpc.Dial 和 insecure.NewCredentials()
	
	// TODO: 检查连接是否有错误
	
	// TODO: 创建 GreeterClient
	// 使用 pb.NewGreeterClient(conn)
	
	// TODO: 创建上下文（带超时）
	// 使用 context.WithTimeout
	
	// TODO: 调用 SayHello 方法
	// 使用 client.SayHello(ctx, &pb.HelloRequest{Name: *name})
	
	// TODO: 打印响应消息
	
	// TODO: 确保关闭连接（defer conn.Close()）
	
	// TODO: 处理错误情况

	// 步骤 1: 创建连接
	conn, err := grpc.NewClient(
		*addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("无法连接服务器 %s: %v", *addr, err)
	}
	defer conn.Close()

	// 步骤 2: 创建客户端
	client := pb.NewGreeterClient(conn)

	// 步骤 3: 创建上下文（带 5 秒超时）
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 步骤 4: 调用 SayHello
	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("调用 SayHello 失败：%v", err)
	}

	// 步骤 5: 打印响应
	fmt.Printf("响应：%s\n", resp.GetMessage())
}

// ============================================
// 客户端测试
// ============================================
// 将此代码放在 client_test.go 中测试客户端