package main

import (
	"context"
	"fmt"
	"net"

	protodata "github.com/rockyskate/spider/proto/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type helloServer struct {
	protodata.UnimplementedGreeterServer
}

func (s *helloServer) SayHello(ctx context.Context, req *protodata.HelloRequest) (*protodata.HelloReply, error) {
	//获取到的客户端传来的值
	fmt.Println("SayHello被调用了")
	name := req.GetName()
	fmt.Printf("Data from client : %s\n", name)
	//将值传回到客户端
	return &protodata.HelloReply{
		Message: "我是服务端",
	}, nil
}

func (s *helloServer) SayRepeatHello(ctx context.Context, req *protodata.RepeatHelloRequest) (*protodata.HelloReply, error) {
	//获取到的客户端传来的值
	fmt.Println("SayRepeatHello被调用了")
	name := req.GetName()
	count := int(req.GetCount())
	for i := 0; i < count; i++ {
		fmt.Printf("%d、Data from client : %s\n", i+1, name)
	}
	//将值传回到客户端
	return &protodata.HelloReply{
		Message: fmt.Sprintf("我是服务端，名称：%v, 次数: %d", name, count),
	}, nil
}

func main() {
	//创建rpc 服务器
	server := grpc.NewServer()
	protodata.RegisterGreeterServer(server, &helloServer{})
	reflection.Register(server)
	//创建 Listen，监听 TCP 端口
	listen, err := net.Listen("tcp", ":8100")
	if err != nil {
		fmt.Println(err)
	}
	defer listen.Close()
	//开启服务
	err = server.Serve(listen)
	if err != nil {
		fmt.Println(err)
	}
}
