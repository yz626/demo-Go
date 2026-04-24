package main

import "Etcd/discovery"

func main() {
	etcdDiscover, err := discovery.NewEtcdDiscovery([]string{"127.0.0.1:2379"}, "server", make(chan struct{}))
	if err != nil {
		panic(err)
	}

	list, err := etcdDiscover.GetServiceList()
	if err != nil {
		return
	}

	for _, v := range list {
		println(v)
	}
}
