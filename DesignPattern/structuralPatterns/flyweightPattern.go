package structuralPatterns

import (
	"bytes"
	"fmt"
)

// 享元模式
// 结构型设计模式
// 主要是减少创建对象的数量，以减少内存消耗和提高性能。
// 允许你在消耗少量内存的情况下支持大量对象，
// 通过共享多个对象的部分状态来实现上述功能，
// 换句话来说，享元会将不同对象的相同数据进行缓存以节省内存。

// example：
// 出租车调度系统中，每辆车的位置是动态变化的，但是车辆信息是固定的
// 车辆信息可以作为一个享元对象
// 在车辆定位中共享车辆信息减少内存消耗。

// Taxi 享元对象，出租车，保存不变的内在属性信息
type Taxi struct {
	color        string // 颜色
	brand        string // 品牌
	company      string // 公司
	licensePlate string // 车牌号
}

// LocateFor 获取定位信息
func (t *Taxi) LocateFor(monitorMap string, x, y int) string {
	return fmt.Sprintf("%s,对于车牌号%s,%s,%s品牌,所属%s公司,定位(%d,%d)", monitorMap,
		t.licensePlate, t.color, t.brand, t.company, x, y)
}

// TaxiFactory 出租车工厂单例
var taxiFactoryInstance = &TaxiFactory{
	taxis: make(map[string]*Taxi),
}

// GetTaxiFactory 获取出租车工厂单例
func GetTaxiFactory() *TaxiFactory {
	return taxiFactoryInstance
}

// TaxiFactory 出租车工厂类
// 保存所有的出租车对象
type TaxiFactory struct {
	taxis map[string]*Taxi
}

// GetTaxi 获取出租车对象
// 当查询一个出租车对象时，如果对象不存在，则创建一个对象，并保存到map中
// 通过这种方式，如果缓存中存在，则直接返回，达到共享的效果。
func (t *TaxiFactory) GetTaxi(color, brand, company, licensePlate string) *Taxi {
	if _, ok := t.taxis[licensePlate]; !ok {
		t.taxis[licensePlate] = &Taxi{
			color:        color,
			brand:        brand,
			company:      company,
			licensePlate: licensePlate,
		}
	}
	return t.taxis[licensePlate]
}

// TaxiPosition 定位信息
type TaxiPosition struct {
	x    int
	y    int
	taxi *Taxi
}

func NewTaxiPosition(x, y int, taxi *Taxi) *TaxiPosition {
	return &TaxiPosition{x: x, y: y, taxi: taxi}
}

// LocateFor 获取定位信息
func (t *TaxiPosition) LocateFor(monitorMap string) string {
	return t.taxi.LocateFor(monitorMap, t.x, t.y)
}

// TaxiDispatcher 出租车调度系统
type TaxiDispatcher struct {
	name   string
	traces map[string][]*TaxiPosition // 存储出租车当天的轨迹信息，key为车牌号
}

func NewTaxiDispatcher(name string) *TaxiDispatcher {
	return &TaxiDispatcher{name: name, traces: make(map[string][]*TaxiPosition)}
}

// AddTrace 添加轨迹信息
func (t *TaxiDispatcher) AddTrace(licensePlate, color, brand, company string, x, y int) {
	// 获取出租车对象
	taxi := GetTaxiFactory().GetTaxi(color, brand, company, licensePlate)
	// 添加轨迹信息
	t.traces[licensePlate] = append(t.traces[licensePlate], NewTaxiPosition(x, y, taxi))
}

// ShowTraces 展示轨迹信息
func (t *TaxiDispatcher) ShowTraces(licensePlate string) string {
	bytesBuf := bytes.Buffer{}
	for _, trace := range t.traces[licensePlate] {
		bytesBuf.WriteString(trace.LocateFor(t.name))
		bytesBuf.WriteString("\n")
	}
	return bytesBuf.String()
}
