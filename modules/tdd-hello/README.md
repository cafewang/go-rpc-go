# TDD 风格 gRPC Hello 模块学习指南

## 学习路线图

本模块通过测试驱动开发（TDD）的方式，带你一步步构建一个最简单的 gRPC client+server 应用。

## 目录结构

```
modules/tdd-hello/
├── proto/
│   └── hello.proto          # proto 文件定义
├── internal/service/
│   ├── greeter_service.go   # 服务实现
│   └── greeter_service_test.go  # 单元测试（TDD 核心）
├── cmd/
│   ├── server/
│   │   └── main.go         # gRPC 服务器
│   └── client/
│       └── main.go         # gRPC 客户端
├── pb/                      # 空的（proto 代码生成在这里）
└── proto/                   # proto 代码生成在这里
└── go.mod                   # Go 模块定义
```

## TDD 三步法

### 第 1 步：测试优先（RED 阶段）
- 先写测试，测试会失败（红色）
- 测试定义了我们要实现的功能

### 第 2 步：实现功能（GREEN 阶段）
- 编写最少的代码让测试通过
- 测试变绿（通过）

### 第 3 步：重构优化（REFACTOR 阶段）
- 在测试保护下重构代码
- 保持测试始终通过

## 学习步骤

### 步骤 1：理解 proto 定义

**文件**: `proto/hello.proto`

**任务**：
1. 阅读 proto 文件
2. 理解服务定义
3. 理解消息结构

**内容**：
- `SayHello` 方法：接收 `HelloRequest`，返回 `HelloReply`
- `HelloRequest`: 包含 `name` 字段
- `HelloReply`: 包含 `message` 字段

**验证**：
```bash
# 检查 proto 文件语法（ protoc 会自动验证）
cd modules/tdd-hello
protoc --go_out=. --go-grpc_out=. proto/hello.proto
```

---

### 步骤 2：运行失败的测试（RED）

**文件**: `internal/service/greeter_service_test.go`

**测试用例**：
1. `TestGreeterService_SayHello_Success_WithValidName` - 测试结果
2. `TestGreeterService_SayHello_Error_EmptyName` - 错误处理
3. `TestGreeterService_SayHello_Success_WithSpecialChars` - 特殊字符测试

**操作**：
```bash
cd modules/tdd-hello/internal/service
go test -v
```

**预期输出**：
```
--- FAIL: TestGreeterService_SayHello_Success_WithValidName
    预期无错误，但得到：rpc error: code = Unimplemented...
--- FAIL: TestGreeterService_SayHello_Error_EmptyName
    预期有错误，但得到 nil
```

**重点**：看到这个失败输出，说明测试已经运行！

---

### 步骤 3：实现服务让测试通过（GREEN）

**文件**: `internal/service/greeter_service.go`

**实现步骤**：
1. 检查请求是否为 nil
2. 检查 name 是否为空
3. 返回错误信息（如果 name 为空）
4. 构造问候语 "Hello, {name}"
5. 返回成功响应

**参考实现**（已完成）：
```go
func (s *GreeterService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
    if req == nil {
        return nil, status.Error(codes.InvalidArgument, "request is nil")
    }
    
    name := req.GetName()
    if name == "" {
        return nil, status.Error(codes.InvalidArgument, "name is required")
    }
    
    message := fmt.Sprintf("Hello, %s", name)
    return &pb.HelloReply{Message: message}, nil
}
```

**运行测试**：
```bash
go test -v
```

**预期输出**：
```
go test: 1 packages ok
```

**恭喜**！所有测试变绿（通过）！

---

### 步骤 4：启动 gRPC 服务器

**文件**: `cmd/server/main.go`

**服务器代码**（已完成）：
```go
func main() {
    port := flag.Int("port", 50051, "gRPC server port")
    flag.Parse()
    
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
    server := grpc.NewServer()
    gr := &service.GreeterService{}
    pb.RegisterGreeterServer(server, gr)
    
    log.Printf("服务器启动，监听端口 %d", *port)
    server.Serve(lis)
}
```

**命令行启动**：
```bash
cd modules/tdd-hello
go run ./cmd/server
```

**预期输出**：
```
2026/06/24 服务器启动，监听端口 50051
```

**保持服务器运行**，进行测试。

---

### 步骤 5：启动 gRPC 客户端

**文件**: `cmd/client/main.go`

**客户端代码**（已完成）：
```go
func main() {
    name := flag.String("name", "World", "要问候的名字")
    addr := flag.String("addr", "localhost:50051", "服务器地址")
    flag.Parse()
    
    conn, _ := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
    client := pb.NewGreeterClient(conn)
    
    resp, _ := client.SayHello(ctx, &pb.HelloRequest{Name: *name})
    fmt.Printf("响应：%s\n", resp.GetMessage())
}
```

**命令行调用**：
```bash
# 新终端窗口
cd modules/tdd-hello
go run ./cmd/client -name "张三"
```

**预期输出**：
```
响应：Hello, 张三
```

**尝试其他名字**：
```bash
go run ./cmd/client -name "Alice"
go run ./cmd/client -name "Hello_GRPC_世界_å¥½"
```

---

### 步骤 6：端对端测试验证

**文件**: `internal/service/integration_test.go`

**运行集成测试**：
```bash
cd modules/tdd-hello
go test -v ./...
```

**预期输出**：
```
=== RUN   TestClientConnection
--- PASS: TestClientConnection (0.00s)
    端对端测试成功：Hello, 测试用户
PASS
```

---

## 学习要点总结

### ✅ 已完成的知识点

1. **proto 文件定义**
   - 服务定义：`service Greeter`
   - RPC 方法定义：`rpc SayHello`
   - 消息定义：`message HelloRequest/HelloReply`

2. **gRPC 代码生成**
   - 使用 `protoc` 生成 Go 代码
   - `hello.pb.go` - 消息结构
   - `hello_grpc.pb.go` - 服务接口

3. **TDD 开发流程**
   - 先写测试（RED）
   - 实现功能（GREEN）
   - 重构优化（REFACTOR）

4. **gRPC 服务器实现**
   - `grpc.NewServer()`
   - `RegisterGreeterServer()`
   - `server.Serve()`

5. **gRPC 客户端实现**
   - `grpc.NewClient()`
   - `NewGreeterClient()`
   - `SayHello()` 方法调用

6. **错误处理**
   - gRPC 状态码：`codes.InvalidArgument`
   - 错误信息返回

7. **测试验证**
   - 单元测试
   - 集成测试
   - 端对端测试

### 🔍 可以探索的扩展

1. **添加更多 RPC 方法**：
   - `SayHelloToAll(names)` - 返回多个问候
   - `StreamHello(name)` - 流式响应

2. **添加错误类型**：
   - 超时处理
   - 认证错误

3. **添加日志**：
   - 请求日志
   - 性能监控

4. **添加度量**：
   - 请求计数
   - 响应时间

---

## 下一步学习

你已经成功完成最简单的 gRPC hello 模块！

**接下来可以尝试**：
1. 查看其他模块（user, product, order 等）学习更复杂的场景
2. 修改这个模块，添加新的 RPC 方法
3. 学习 gRPC 的流式功能（server streaming, client streaming, bidirectional streaming）
4. 学习 gRPC 的拦截器（interceptor）
5. 学习 gRPC 的负载均衡

---

## 常见问题

**Q: 测试一直失败怎么办？**
A: 检查：
1. 包路径是否正确
2. 是否运行了 `go mod tidy`
3. prot 代码是否生成

**Q: 客户端连接不上服务器？**
A: 检查：
1. 服务器是否在运行
2. 端口是否冲突
3. 地址是否是 localhost:50051

**Q: 如何调试？**
A: 
1. 使用 `go test -v` 查看详细输出
2. 在代码中添加 `log.Printf`
3. 使用 `fmt.Printf` 打印变量

---

**祝你学习愉快！** 🎉