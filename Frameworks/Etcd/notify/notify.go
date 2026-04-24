package notify

type Notify struct {
	stop chan struct{}
}

func NewNotify(stop chan struct{}) *Notify {
	return &Notify{stop: stop}
}

// Stop 当etcd 服务挂掉时，通知服务停止
func (n *Notify) Stop() {
	n.stop <- struct{}{}
	close(n.stop)
}

type Update struct {
	update chan struct{}
}

func NewUpdate(update chan struct{}) *Update {
	return &Update{update: update}
}

func (u *Update) Update() {
	u.update <- struct{}{}
}
