package structuralPatterns

import (
	"fmt"
	"testing"
)

// 适配器模式
// 结构型设计模式
// 通过一个中间件（适配器）将一个类的接口转换成客户希望的另外一个接口
// 适配器模式使得原本由于接口不兼容而不能一起工作的那些类可以一起工作
// 它会接收对于一个对象的调用， 并将其转换为另一个对象可识别的格式和接口

// example:
// 以充电宝接口为例，充电宝可以给不同手机或其他设备充电
// 但是设备的充电接口和充电宝的接口并不相同，
// 连接充电宝的通用接口就作为一个适配器

// HuaweiPlug 华为充电接口
type HuaweiPlug interface {
	ConnectTypeC() string
}

// HuaweiPhone 华为手机
type HuaweiPhone struct {
	model string
}

func NewHuaweiPhone(model string) *HuaweiPhone {
	return &HuaweiPhone{model}
}

// ConnectTypeC 实现华为充电接口
func (h *HuaweiPhone) ConnectTypeC() string {
	return "Connect TypeC"
}

var _ HuaweiPlug = (*HuaweiPhone)(nil)

// ApplePlug 苹果充电接口
type ApplePlug interface {
	ConnectLightning() string
}

// IPhone 苹果手机
type IPhone struct {
	model string
}

func NewIPhone(model string) *IPhone {
	return &IPhone{model}
}

// ConnectLightning 实现苹果充电接口
func (i *IPhone) ConnectLightning() string {
	return "Connect Lightning"
}

var _ ApplePlug = (*IPhone)(nil)

// CommonPlug 通用的USB接口
type CommonPlug interface {
	ConnectUSB() string
}

// HuaweiPlugAdapter 华为适配通用USB充电插槽
type HuaweiPlugAdapter struct {
	huaweiPhone *HuaweiPhone
}

func NewHuaweiPlugAdapter(huaweiPhone *HuaweiPhone) *HuaweiPlugAdapter {
	return &HuaweiPlugAdapter{huaweiPhone}
}

// ConnectUSB 华为适配器实现通用USB充电接口
func (h *HuaweiPlugAdapter) ConnectUSB() string {
	return h.huaweiPhone.ConnectTypeC()
}

var _ CommonPlug = (*HuaweiPlugAdapter)(nil)

type ApplePlugAdapter struct {
	iphone *IPhone
}

func NewApplePlugAdapter(iphone *IPhone) *ApplePlugAdapter {
	return &ApplePlugAdapter{iphone}
}

// ConnectUSB 苹果适配器实现通用USB充电接口
func (a *ApplePlugAdapter) ConnectUSB() string {
	return a.iphone.ConnectLightning()
}

var _ CommonPlug = (*ApplePlugAdapter)(nil)

// PowerBank 充电宝
type PowerBank struct {
	brand string
}

// Charge 支持通用USB接口充电
func (p *PowerBank) Charge(plug CommonPlug) string {
	return plug.ConnectUSB()
}

// 测试程序

func TestAdapterPattern(t *testing.T) {
	huaweiPhone := NewHuaweiPhone("Huawei P30")
	iphone13MaxPro := NewIPhone("iPhone 13 Max Pro")

	powerBank := &PowerBank{brand: "XXX"}
	fmt.Println(powerBank.Charge(NewHuaweiPlugAdapter(huaweiPhone)))
	fmt.Println(powerBank.Charge(NewApplePlugAdapter(iphone13MaxPro)))
}
