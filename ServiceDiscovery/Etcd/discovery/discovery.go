package discovery

import (
	"Etcd/notify"
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type EtcdDiscovery struct {
	serviceName    string           // 服务名称
	etcdClient     *clientv3.Client // etcd客户端
	*notify.Update                  // 监听服务列表的变化，然后通知客户端更新服务列表
}

func NewEtcdDiscovery(etcdAddr []string, serviceName string, update chan struct{}) (*EtcdDiscovery, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   etcdAddr,
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		return nil, err
	}

	return &EtcdDiscovery{
		serviceName: serviceName,
		etcdClient:  client,
		Update:      notify.NewUpdate(update),
	}, nil
}

// Discover 从etcd中获取服务调用的gRPC
// 并不一定需要使用，因为gRPC本身已经实现了负载均衡，这里只是为了演示
func (e *EtcdDiscovery) Discover() (*grpc.ClientConn, error) {
	builder, err := resolver.NewBuilder(e.etcdClient)
	if err != nil {
		return nil, err
	}

	// 创建gRPC连接etcd获取服务列表
	// gRPC连接etcd的前缀固定为 etcd:///
	// grpc.WithResolvers(builder): 指定解析器
	// grpc.WithTransportCredentials(insecure.NewCredentials()): 指定使用不安全的传输
	return grpc.NewClient("etcd:///"+e.serviceName, grpc.WithResolvers(builder), grpc.WithTransportCredentials(insecure.NewCredentials()))
}

// GetServiceList 从etcd中获取服务列表
func (e *EtcdDiscovery) GetServiceList() ([]string, error) {
	manager, err := endpoints.NewManager(e.etcdClient, e.serviceName)
	if err != nil {
		return nil, err
	}

	list, err := manager.List(context.Background())
	if err != nil {
		return nil, err
	}

	var res []string
	for _, v := range list {
		res = append(res, v.Addr)
	}
	return res, nil
}

// Watch 监听服务列表的变化然后通知客户端更新列表
func (e *EtcdDiscovery) Watch() {
	watch := clientv3.NewWatcher(e.etcdClient)
	watchChan := watch.Watch(context.Background(), e.serviceName, clientv3.WithPrefix())

	for watchRes := range watchChan {
		for _, event := range watchRes.Events {
			switch event.Type {
			case clientv3.EventTypePut:
				e.Update.Update()
			case clientv3.EventTypeDelete:
				e.Update.Update()
			}
		}
	}
}
