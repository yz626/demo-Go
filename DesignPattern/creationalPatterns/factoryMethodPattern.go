package creationalPatterns

import "fmt"

//  工厂方法模式
// 在父类中提供一个创建对象的方法， 允许子类决定实例化对象的类型

// example:
// 摊煎饼的小贩需要先摊个煎饼，再卖出去，摊煎饼就可以类比为一个工厂方法，
// 根据顾客的喜好摊出不同口味的煎饼。

// Pancake 煎饼
type Pancake interface {
	ShowFlour() string // 展示面粉
	Value() int        // 展示价格
}

// PancakeCook 煎饼厨师
type PancakeCook interface {
	MakePancake() Pancake
}

// PancakeVendor 煎饼摊
type PancakeVendor struct {
	PancakeCook
}

func NewPancakeVendor(cook PancakeCook) *PancakeVendor {
	return &PancakeVendor{cook}
}

// SellPancake 卖煎饼, 先制作再卖出去
func (p *PancakeVendor) SellPancake() int {
	return p.MakePancake().Value()
}

// 煎饼的各种实现

// CornPancake 玉米煎饼
type CornPancake struct{}

func (c *CornPancake) ShowFlour() string {
	return "玉米面"
}

func (c *CornPancake) Value() int {
	return 10
}

func NewCornPancake() *CornPancake {
	return &CornPancake{}
}

var _ Pancake = (*CornPancake)(nil)

// MilletPancake 小麦煎饼
type MilletPancake struct{}

func (m *MilletPancake) ShowFlour() string {
	return "小麦面"
}

func (m *MilletPancake) Value() int {
	return 20
}

func NewMilletPancake() *MilletPancake {
	return &MilletPancake{}
}

var _ Pancake = (*MilletPancake)(nil)

// CornPancakeVendor 玉米面煎饼厨师
type CornPancakeVendor struct{}

func (c *CornPancakeVendor) MakePancake() Pancake {
	return NewCornPancake()
}

func NewCornPancakeVendor() *CornPancakeVendor {
	return &CornPancakeVendor{}
}

// MilletPancakeVendor 小麦煎饼厨师
type MilletPancakeVendor struct{}

func (m *MilletPancakeVendor) MakePancake() Pancake {
	return NewMilletPancake()
}

func NewMilletPancakeVendor() *MilletPancakeVendor {
	return &MilletPancakeVendor{}
}

// TestFactoryMethodPattern 测试工厂方法模式
func TestFactoryMethodPattern() {
	pancakeVendor := NewPancakeVendor(NewCornPancakeVendor())
	fmt.Println(pancakeVendor.SellPancake())

	pancakeVendor = NewPancakeVendor(NewMilletPancakeVendor())
	fmt.Println(pancakeVendor.SellPancake())
}
