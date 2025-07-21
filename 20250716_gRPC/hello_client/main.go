package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "gRPC/hello_client/proto"
)

func main() {
	// 连接到server端, 此处禁用安全传输, 没有加密和验证
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 建立连接 (创建客户端)
	client := pb.NewSayHelloClient(conn)

	// 执行rpc调用(这个方法在服务器端来实现并返回结果)
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "Kunkun"})
	fmt.Println(resp.ResponseMsg)
}
