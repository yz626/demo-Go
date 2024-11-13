package structuralPatterns

import (
	"bytes"
	"fmt"
)

// 代理模式
// 一个类型代表另一个类型的功能
// 通过引入一个代理对象来控制对于原来对象的访问
// 代理对象在客户端喝目标对象之间充当中介
// 负责将客户端对目标对象的请求转发给目标对象
// 同时可以进行额外的操作

// 代理模式和修饰器模式类似，都是通过代理对象来控制对目标对象的访问
// 但是有所不同
// 修饰器模式：一定会执行服务对象，是对于执行之前或者执行之后的结果进行加强
//			   服务对象基本是客户端创建好的，然后嵌套外层的修饰对象。
// 代理模式：  可能不会执行服务对象，同时什么时候创建服务对象，由代理对象决定。

// example:
// 房屋中介帮助房东卖房子

// HouseSeller 房屋出售者
type HouseSeller interface {
	SellHouse(address, buyer string) string
}

// HouseProxy 房产中介代理
type HouseProxy struct {
	seller HouseSeller
}

func NewHouseProxy(seller HouseSeller) *HouseProxy {
	return &HouseProxy{seller}
}

// viewHouse 查看房子的基本情况
func (h *HouseProxy) viewHouse(address, buyer string) string {
	return fmt.Sprintf("带买家%s看位于%s的房屋，并介绍基本情况", buyer, address)
}

// preBargain 初步沟通价格
func (h *HouseProxy) preBargain(address, buyer string) string {
	return fmt.Sprintf("讨价还价后，初步达成购买意向")
}

func (h *HouseProxy) SellHouse(address, buyer string) string {
	buf := bytes.Buffer{}
	buf.WriteString(h.viewHouse(address, buyer) + "\n")
	buf.WriteString(h.preBargain(address, buyer) + "\n")
	buf.WriteString(h.seller.SellHouse(address, buyer))
	return buf.String()
}

// HouseOwner 房东
type HouseOwner struct{}

func (h *HouseOwner) SellHouse(address, buyer string) string {
	return fmt.Sprintf("最终商讨价格后，与%s签署购买地址为%s的购房协议。", buyer, address)
}

var _ HouseSeller = (*HouseOwner)(nil)
