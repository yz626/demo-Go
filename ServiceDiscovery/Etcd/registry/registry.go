package registry

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type Registry interface {
	Register(serviceName string, IP string, port int) error
}

type EtcdAgent struct {
	client *clientv3.Client
}

func NewAgent(etcdAddr []string, ttl int) (Registry, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   etcdAddr,
		DialTimeout: time.Second * 5,
	})

	if err != nil {
		return nil, err
	}

	return &EtcdAgent{
		client: client,
	}, nil
}

func (e *EtcdAgent) Register(serviceName string, IP string, port int) error {
	// TODO: implement
	return nil
}
