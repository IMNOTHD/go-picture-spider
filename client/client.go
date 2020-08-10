package main

import (
	"context"
	"log"

	pb "go-picture-spider/api"

	"google.golang.org/grpc"
)

const Address string = ":8000"

var grpcClient pb.PictureSpiderClient

func main() {
	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	grpcClient = pb.NewPictureSpiderClient(conn)
	testPictureSpider()
}

// sayHelloWorld 调用服务端Route方法
func testPictureSpider() {
	// 创建发送结构体
	req := pb.PictureUrlRequest{
		Type: "yande",
		Id:   "557301",
	}
	// 调用我们的服务(Route方法)
	// 同时传入了一个 context.Context ，在有需要时可以让我们改变RPC的行为，比如超时/取消一个正在运行的RPC
	res, err := grpcClient.PictureUrl(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call Route err: %v", err)
	}
	// 打印返回值
	log.Println(res)
}
