package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	protodata "github.com/rockyskate/spider/proto/go"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	var serviceHost = "127.0.0.1:8100"

	conn, err := grpc.Dial(serviceHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	client := protodata.NewGreeterClient(conn)
	rsp, err := client.SayHello(context.TODO(), &protodata.HelloRequest{
		Name: "Jakerson",
	})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rsp.Message)
	rsp, err = client.SayRepeatHello(context.TODO(), &protodata.RepeatHelloRequest{
		Name:  "Jakerson",
		Count: 2,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rsp.Message)
	fmt.Println("按回车键退出程序...")
	in := bufio.NewReader(os.Stdin)
	_, _, _ = in.ReadLine()
}
