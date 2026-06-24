package service

import (
	"context"
	"fmt"
	pb "github.com/example/tdd-hello/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GreeterService 实现 Greeter 服务
//
// TODO 步骤 1 完成：服务实现
//
// 实现说明：
// 1. 检查请求是否为 nil
// 2. 验证 name 字段是否为空
// 3. 如果 name 为空，返回 InvalidArgument 错误，消息为 "name is required"
// 4. 如果 name 有效，构造问候语 "Hello, {name}"
// 5. 返回 HelloReply 包含问候语
//
// 完成后，运行: go test ./internal/service
// 预期结果：所有测试通过

type GreeterService struct {
	pb.UnimplementedGreeterServer
}

// SayHello 实现 Greeter 服务的 SayHello 方法
//
// 实现步骤：
// 1. 检查请求是否为 nil
// 2. 检查 name 是否为空
// 3. 处理错误情况
// 4. 返回成功的响应
func (s *GreeterService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	// 步骤 1: 检查请求是否为 nil
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}

	// 步骤 2: 检查 name 是否为空
	name := req.GetName()
	if name == "" {
		// 步骤 3: 返回错误
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}

	// 步骤 4: 构造问候语并使用 fmt.Sprintf 确保格式正确
	message := fmt.Sprintf("Hello, %s", name)

	// 步骤 5: 返回成功的响应
	return &pb.HelloReply{Message: message}, nil
}