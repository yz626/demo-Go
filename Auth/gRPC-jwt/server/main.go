package main

import (
	"context"
	"fmt"
	"gRPC-jwt/auth"
	token "gRPC-jwt/proto"
	"gRPC-jwt/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
)

// 拦截器，用于打印日志
func loggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	fmt.Printf("grpc method: %s, %v\n", info.FullMethod, req)
	resp, err = handler(ctx, req)
	fmt.Printf("grpc method: %s, %v\n", info.FullMethod, resp)
	return
}

// 实现 TokenServiceServer 接口
type login struct {
	token.UnimplementedTokenServiceServer
}

func (l *login) Login(ctx context.Context, req *token.LoginRequest) (*token.LoginResponse, error) {
	if req.Username == "admin" && req.Password == "123456" {
		tokenStr, err := utils.CreateToken(req.Username)
		if err != nil {
			return &token.LoginResponse{
				Status: "500",
				Token:  "",
			}, err
		}
		return &token.LoginResponse{
			Status: "200",
			Token:  tokenStr,
		}, nil
	}

	return &token.LoginResponse{
		Status: "401",
		Token:  "",
	}, nil
}

func (l *login) SayHello(ctx context.Context, request *token.HelloRequest) (*token.HelloResponse, error) {
	name, err := auth.Auth{}.CheckToken(ctx)
	if err != nil {
		return nil, err
	}
	return &token.HelloResponse{
		Message: "hello " + name,
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}

	// 创建grpc服务, 并添加拦截器
	s := grpc.NewServer(grpc.Creds(insecure.NewCredentials()), grpc.UnaryInterceptor(loggerInterceptor))
	// 注册服务
	token.RegisterTokenServiceServer(s, &login{})
	// 启动服务
	s.Serve(listen)
}
