package main

import (
	"Etcd/discovery"
	"log"
	"sync"
)

type Client struct {
	name         string
	host         string // 本地服务器地址
	port         string // 端口
	mu           sync.Mutex
	remoteServer []string      // 远程服务器地址
	update       chan struct{} // 监听更新
}

func (c *Client) Run() {}

func NewClient(name, host, port string) *Client {
	return &Client{
		name:         name,
		host:         host,
		port:         port,
		update:       make(chan struct{}),
		remoteServer: []string{},
	}
}

func (c *Client) AddRemoteServer() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	// TODO: 重构，将client等长连接单独拎出来，避免每次都去etcd拉取
	agent, err := discovery.NewEtcdDiscovery([]string{"127.0.0.1:2379"}, "server", c.update)
	if err != nil {
		return err
	}

	list, err := agent.GetServiceList()
	if err != nil {
		return err
	}

	c.remoteServer = list
	log.Println("remote server list: ", list)
	return nil
}

func (c *Client) Watch() {

	for {
		select {
		case <-c.update:
			_ = c.AddRemoteServer()
		}
	}
}
