package main

// 创建服务实例
// 启动服务
// 注册服务
// 停止服务

func main() {
	server := NewServer(LocalServerName, "127.0.0.1", 8080)
	err := server.Register()
	if err != nil {
		panic(err)
	}

	err = server.Start()
	if err != nil {
		return
	}
}
