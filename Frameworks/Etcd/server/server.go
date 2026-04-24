package main

import (
	"Etcd/registry"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

var (
	LocalServerName = "server"
)

type Server struct {
	Name   string
	Host   string
	Port   int
	end    func()        // 注册到etcd后将返回一个取消注册的函数
	stop   chan struct{} // 监听服务是否停止
	Status bool          // 服务是否正常
}

func NewServer(name string, host string, port int) *Server {
	return &Server{
		Name:   name,
		Host:   host,
		Port:   port,
		stop:   make(chan struct{}),
		Status: true,
		end:    func() {},
	}
}

// Update 监听服务状态变化，接收到停止信号后，关闭服务
func (s *Server) Update() {
	for _ = range s.stop {
		s.Stop()
		return
	}
}

// Stop 停止服务
func (s *Server) Stop() {
	s.Status = false
	s.end()
}

// Start 启动服务
func (s *Server) Start() error {
	// TODO: 启动grpc服务
	// TODO: 将服务注册到etcd
	// TODO: 监听服务状态变化确定更新服务列表

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Port))
	if err != nil {
		return err
	}

	err = grpc.NewServer().Serve(listen)
	if err != nil {
		return err
	}
	return nil
}

// Register 注册服务
func (s *Server) Register() error {
	agent, err := registry.NewAgent([]string{"127.0.0.1:2379"}, 5, s.stop)
	if err != nil {
		return err
	}

	f, err := agent.Register(s.Name, s.Host, strconv.Itoa(s.Port))
	if err != nil {
		return err
	}

	s.end = f
	return nil
}
