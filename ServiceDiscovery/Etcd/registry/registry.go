package registry

import (
	"Etcd/notify"
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"time"
)

type Registry interface {
	Register(serviceName string, IP, port string) (func(), error)
}

type EtcdAgent struct {
	client *clientv3.Client
	*notify.Notify
}

func NewAgent(etcdAddr []string, ttl int, stop chan struct{}) (Registry, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   etcdAddr,
		DialTimeout: time.Second * time.Duration(ttl),
	})

	if err != nil {
		return nil, err
	}

	return &EtcdAgent{
		client: client,
		Notify: notify.NewNotify(stop),
	}, nil
}

func (e *EtcdAgent) Register(serviceName string, IP, port string) (func(), error) {
	// 创建租约
	lease, err := e.client.Lease.Grant(context.TODO(), 5)
	if err != nil {
		return nil, err
	}
	// 创建租约的KeepAlive
	// 该方法会定期进行续租，如果通道返回错误，说明续租失败
	leaseKeepAliveChan, err := e.client.KeepAlive(context.Background(), lease.ID)
	if err != nil {
		return nil, err
	}

	if err := e.addService(serviceName, IP, port); err != nil {
		return nil, err
	}

	// 监听续租的chan
	go func() {
		for {
			select {
			case _, ok := <-leaseKeepAliveChan:
				// 如果租约失效，通知服务关闭
				// 租约失效之后，etcd会自动删除无效的key，因此这里不需要手动删除key
				if !ok {
					e.Notify.Stop()
				}
			}
		}
	}()

	// 返回取消注册的函数
	return func() {
		err := e.removeService(serviceName, IP)
		if err != nil {
			return
		}
	}, nil
}

func (e *EtcdAgent) addService(serviceName string, IP, port string) error {
	// 获取指定的服务名对应的服务列表的Manager
	manager, err := endpoints.NewManager(e.client, serviceName)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("%s/%s", serviceName, IP)
	return manager.AddEndpoint(context.TODO(), key, endpoints.Endpoint{
		Addr:     fmt.Sprintf("%s:%s", IP, port),
		Metadata: "metadata",
	})
}

func (e *EtcdAgent) removeService(serviceName string, IP string) error {
	// 获取指定的服务名对应的服务列表的Manager
	manager, err := endpoints.NewManager(e.client, serviceName)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("%s/%s", serviceName, IP)
	err = manager.DeleteEndpoint(context.TODO(), key)
	if err != nil {
		return err
	}
	return nil
}
