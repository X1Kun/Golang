package main

import (
	"context"
	"fmt"
	"net"

	pb "gRPC/hello_server/proto"

	"google.golang.org/grpc"
)

// hello server
// 定义一个 server 结构体，用于实现 .proto 文件中定义的服务接口
type server struct {
	// 嵌入 UnimplementedSayHelloServer 是 gRPC 的一个最佳实践，
	// 它可以确保在未来向 .proto 添加新方法时，旧的服务端实现依然能编译通过。
	pb.UnimplementedSayHelloServer
}

// SayHello 方法是 SayHelloServer 接口的具体实现。
// 这就是服务的核心业务逻辑。
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	// 从请求中获取名字，并拼接一个 "hello" 字符串返回。
	fmt.Println("hello " + req.RequestName)
	return &pb.HelloResponse{ResponseMsg: "hello " + req.RequestName}, nil
}

func main() {
	// 开启端口
	// 监听本地的 9090 TCP 端口
	listen, _ := net.Listen("tcp", ":9090")

	// 创建grpc服务
	// 初始化一个新的 gRPC 服务器实例
	grpcServer := grpc.NewServer()

	// 在grpc服务端中去注册我们自己编写的服务
	// 将我们自己实现的 server 注册到 gRPC 服务器上。
	// RegisterSayHelloServer 是由 protoc 自动生成的函数。
	pb.RegisterSayHelloServer(grpcServer, &server{})

	// 启动服务
	// 让 gRPC 服务器开始在之前监听的端口上提供服务。
	// Serve 会阻塞当前 goroutine，直到服务器停止。
	err := grpcServer.Serve(listen)
	if err != nil {
		// 在实际项目中，这里应该有错误处理，例如记录日志
		return
	}
}
