package service

import (
	"context"
	"testing"

	pb "github.com/example/tdd-hello/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ============================================
// TODO 步骤 1：实现服务
// ============================================
// 完成事项：
// 1. 创建 GreeterService 结构体
// 2. 实现 SayHello 方法
// 3. 验证 name 不能为空
// 4. 返回 "Hello, {name}" 格式的问候语
// 5. 空 name 时返回 codes.InvalidArgument 错误
//
// 通过后测试：
// - TestGreeterService_SayHello_Success_WithValidName
// - TestGreeterService_SayHello_Error_EmptyName
// - TestGreeterService_SayHello_Success_WithSpecialChars

// TestGreeterService_SayHello_Success_WithValidName
// 测试目标：当提供有效的用户名时，应该返回正确的问候语
// TODO: 实现服务代码使此测试通过
func TestGreeterService_SayHello_Success_WithValidName(t *testing.T) {
	// 初始化服务
	service := &GreeterService{}

	// 创建请求
	req := &pb.HelloRequest{Name: "Alice"}

	// 调用 SayHello
	resp, err := service.SayHello(context.Background(), req)

	// 验证无错误
	if err != nil {
		t.Fatalf("预期无错误，但得到：%v", err)
	}

	// 验证返回消息
	expected := "Hello, Alice"
	if resp.GetMessage() != expected {
		t.Errorf("预期问候语 '%s'，但得到 '%s'", expected, resp.GetMessage())
	}
}

// TestGreeterService_SayHello_Error_EmptyName
// 测试目标：当提供空用户名时，应该返回 InvalidArgument 错误
// TODO: 实现服务代码使此测试通过
func TestGreeterService_SayHello_Error_EmptyName(t *testing.T) {
	// 初始化服务
	service := &GreeterService{}

	// 创建空 name 请求
	req := &pb.HelloRequest{Name: ""}

	// 调用 SayHello
	_, err := service.SayHello(context.Background(), req)

	// 验证有错误
	if err == nil {
		t.Fatal("预期有错误，但得到 nil")
	}

	// 验证错误码是 InvalidArgument
	if status.Code(err) != codes.InvalidArgument {
		t.Errorf("预期错误码 InvalidArgument，但得到 %v", status.Code(err))
	}

	// 验证错误消息包含 "name is required"
	if err.Error() != "rpc error: code = InvalidArgument desc = name is required" {
		t.Errorf("预期错误消息包含 'name is required'，但得到 '%s'", err.Error())
	}
}

// TestGreeterService_SayHello_Success_WithSpecialChars
// 测试目标：支持特殊字符和中文名字
// TODO: 实现服务代码使此测试通过
func TestGreeterService_SayHello_Success_WithSpecialChars(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{"Bob", "Hello, Bob"},
		{"张三", "Hello, 张三"},
		{"John_Doe", "Hello, John_Doe"},
		{"user@example.com", "Hello, user@example.com"},
		{"  空格  ", "Hello,   空格  "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &GreeterService{}
			req := &pb.HelloRequest{Name: tt.name}

			resp, err := service.SayHello(context.Background(), req)

			if err != nil {
				t.Fatalf("预期无错误，但得到：%v", err)
			}

			if resp.GetMessage() != tt.expected {
				t.Errorf("预期问候语 '%s'，但得到 '%s'", tt.expected, resp.GetMessage())
			}
		})
	}
}

// ============================================
// TODO 步骤 2：实现服务器
// ============================================
// 完成事项：
// 1. 在 cmd/server/main.go 中创建 gRPC 服务器
// 2. 注册 GreeterService
// 3. 监听端口 50051
// 4. 启动服务器
//
// 通过后测试：
// - TestServer_StartsSuccessfully
// - TestServer_CanHandleRequests

// ============================================
// TODO 步骤 3：实现客户端
// ============================================
// 完成事项：
// 1. 在 cmd/client/main.go 中创建 gRPC 客户端
// 2. 连接服务器（localhost:50051）
// 3. 调用 SayHello 方法
// 4. 打印响应
//
// 通过后测试：
// - TestClient_CanConnectAndCall

// ============================================
// TODO 步骤 4：集成测试
// ============================================
// 完成事项：
// 1. 启动服务器
// 2. 创建客户端连接
// 3. 调用 SayHello
// 4. 验证整个流程
//
// 通过后测试：
// - TestEndToEnd_HelloRPC