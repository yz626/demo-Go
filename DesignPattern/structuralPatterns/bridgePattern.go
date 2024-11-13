package structuralPatterns

import (
	"fmt"
)

// 桥接模式
// 将业务的一个大类拆分成不同层次结构，从而独立的进行开发。
// 第一层：抽象层
// 第二层：实现层，继承抽象层
// 能够减少耦合，方便扩展

// example：
// 用描述一段经历作为例子
// 一段经历为存在三个方面：
// 经历本身：旅游、探险
// 		抽象层：Experience
// 		实现层：travel、adventure
// 交通方式：飞机、汽车
// 		抽象层：Traffic
// 		实现层：airplane、car
// 活动的开展方式：冲浪、攀岩
// 		抽象层：Location
// 		实现层：seaside、mountain

// Traffic 交通工具
type Traffic interface {
	Transport() string
}

// airplane 飞机
type airplane struct{}

func (a *airplane) Transport() string {
	return "by airplane"
}

var _ Traffic = (*airplane)(nil)

// car 汽车
type car struct{}

func (c *car) Transport() string {
	return "by car"
}

var _ Traffic = (*car)(nil)

// Location 活动
type Location interface {
	Name() string   // 活动名称
	Sports() string // 活动方式
}

// nameLocation 被命名的地点，统一按照此种引用类型
type nameLocation struct {
	name string
}

func (n *nameLocation) Name() string {
	return n.name
}

// seaside 海滨
type seaside struct {
	nameLocation
}

func NewSeaside(name string) *seaside {
	return &seaside{nameLocation{name}}
}

// Sports 冲浪
func (s *seaside) Sports() string {
	return "surfing"
}

// mountain 山地
type mountain struct {
	nameLocation
}

func NewMountain(name string) *mountain {
	return &mountain{nameLocation{name}}
}

// Sports 攀岩
func (m *mountain) Sports() string {
	return "climbing"
}

// Experience 经历
type Experience interface {
	Describe() string // 描述
}

// travel 旅游经历
type travel struct {
	subject  string   // 主题
	traffic  Traffic  // 交通工具
	location Location // 活动
}

func NewTravel(subject string, traffic Traffic, location Location) *travel {
	return &travel{subject, traffic, location}
}

func (t *travel) Describe() string {
	return fmt.Sprintf("I have a %s experience in %s by %s", t.subject, t.location.Name(), t.traffic.Transport())
}

// adventure 探险经历
type adventure struct {
	subject string
	travel
}

func NewAdventure(subject string, traffic Traffic, location Location) *adventure {
	return &adventure{subject, travel{subject, traffic, location}}
}

func (a *adventure) Describe() string {
	return fmt.Sprintf("I have a %s experience in %s by %s", a.subject, a.location.Name(), a.traffic.Transport())
}
