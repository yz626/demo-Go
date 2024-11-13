package main

import (
	"context"
	"fmt"
	"gRPC-jwt/auth"
	token "gRPC-jwt/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 创建连接
	conn, err := grpc.NewClient("127.0.0.1:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// 创建客户端
	c := token.NewTokenServiceClient(conn)
	login, err := c.Login(context.Background(), &token.LoginRequest{
		Username: "admin",
		Password: "123456",
	})
	if err != nil {
		return
	}

	if login.Status != "200" {
		return
	}

	conn2, err := grpc.NewClient("127.0.0.1:8081",
		grpc.WithPerRPCCredentials(auth.NewToken(login.Token)),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn2.Close()

	c2 := token.NewTokenServiceClient(conn2)
	hello, err := c2.SayHello(context.Background(), &token.HelloRequest{})
	if err != nil {
		return
	}
	fmt.Println(hello.Message)
}
