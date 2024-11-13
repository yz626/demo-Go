package creationalPatterns

import (
	"fmt"
	"testing"
)

// 生成器/创建者模式
// 它允许逐步构建一个复杂的对象。该模式允许用户以相同的构建步骤来创建多种不同的表示。
// 生成器模式通常包括以下部分：
// Builder：定义构建过程的接口。
// Concrete Builders：实现构建过程，负责组装各个部件。
// Director：控制构建过程，调用具体的生成器来创建对象的不同部分。
// Product：最终被创建出来的复杂对象。
// 生成器模式适用于当创建复杂对象的算法应该独立于组成该对象的部分以及它们的装配方式时。

// example：
// 以摊煎饼为例子，为了避免每次都调用相同的构建步骤
// 通过实现不同的构造器，可以生成不同种类的煎饼。

// Quantity 分量
type Quantity int

const (
	Small  Quantity = 1
	Middle Quantity = 5
	Large  Quantity = 10
)

type PancakeBuilder interface {
	// PutPaste 放面糊
	PutPaste(quantity Quantity)
	// PutEgg 放鸡蛋
	PutEgg(num int)
	// PutWafer 放薄脆
	PutWafer()
	// PutFlavour 放调料 Coriander香菜，Shallot葱 Sauce酱
	PutFlavour(hasCoriander, hasShallot, hasSauce bool)
	// Build 摊煎饼
	Build() *Pancake2
}

// Pancake2  煎饼
type Pancake2 struct {
	pasteQuantity Quantity // 面糊分量
	eggNum        int      // 鸡蛋数量
	wafer         string   // 薄脆
	hasCoriander  bool     // 是否放香菜
	hasShallot    bool     // 是否放葱
	hasSauce      bool     // 是否放酱
}

// normalPancakeBuilder 正常煎饼创建器
type normalPancakeBuilder struct {
	pasteQuantity Quantity // 面糊量
	eggNum        int      // 鸡蛋数量
	friedWafer    string   // 油炸薄脆
	hasCoriander  bool     // 是否放香菜
	hasShallot    bool     // 是否放葱
	hasHotSauce   bool     // 是否放辣味酱
}

func NewNormalPancakeBuilder() *normalPancakeBuilder {
	return &normalPancakeBuilder{}
}

func (n *normalPancakeBuilder) PutPaste(quantity Quantity) {
	n.pasteQuantity = quantity
}

func (n *normalPancakeBuilder) PutEgg(num int) {
	n.eggNum = num
}

func (n *normalPancakeBuilder) PutWafer() {
	n.friedWafer = "油炸的薄脆"
}

func (n *normalPancakeBuilder) PutFlavour(hasCoriander, hasShallot, hasSauce bool) {
	n.hasCoriander = hasCoriander
	n.hasShallot = hasShallot
	n.hasHotSauce = hasSauce
}

func (n *normalPancakeBuilder) Build() *Pancake2 {
	return &Pancake2{
		pasteQuantity: n.pasteQuantity,
		eggNum:        n.eggNum,
		wafer:         n.friedWafer,
		hasCoriander:  n.hasCoriander,
		hasShallot:    n.hasShallot,
		hasSauce:      n.hasHotSauce,
	}
}

// healthyPancakeBuilder 健康煎饼创建器
type healthyPancakeBuilder struct {
	milletPasteQuantity Quantity // 小米面糊量
	chaiEggNum          int      // 柴鸡蛋数量
	nonFriedWafer       string   // 非油炸薄脆
	hasCoriander        bool     // 是否放香菜
	hasShallot          bool     // 是否放葱
}

func NewHealthyPancakeBuilder() *healthyPancakeBuilder {
	return &healthyPancakeBuilder{}
}

func (n *healthyPancakeBuilder) PutPaste(quantity Quantity) {
	n.milletPasteQuantity = quantity
}

func (n *healthyPancakeBuilder) PutEgg(num int) {
	n.chaiEggNum = num
}

func (n *healthyPancakeBuilder) PutWafer() {
	n.nonFriedWafer = "非油炸的薄脆"
}

func (n *healthyPancakeBuilder) PutFlavour(hasCoriander, hasShallot, _ bool) {
	n.hasCoriander = hasCoriander
	n.hasShallot = hasShallot
}

func (n *healthyPancakeBuilder) Build() *Pancake2 {
	return &Pancake2{
		pasteQuantity: n.milletPasteQuantity,
		eggNum:        n.chaiEggNum,
		wafer:         n.nonFriedWafer,
		hasCoriander:  n.hasCoriander,
		hasShallot:    n.hasShallot,
		hasSauce:      false,
	}
}

// Pancake2Cook 摊煎饼师傅
type Pancake2Cook struct {
	builder PancakeBuilder
}

func NewPancakeCook(builder PancakeBuilder) *Pancake2Cook {
	return &Pancake2Cook{
		builder: builder,
	}
}

// SetPancakeBuilder 重新设置煎饼构造器
func (p *Pancake2Cook) SetPancakeBuilder(builder PancakeBuilder) {
	p.builder = builder
}

// MakePancake 摊一个一般煎饼
func (p *Pancake2Cook) MakePancake() *Pancake2 {
	p.builder.PutPaste(Middle)
	p.builder.PutEgg(1)
	p.builder.PutWafer()
	p.builder.PutFlavour(true, true, true)
	return p.builder.Build()
}

// MakeBigPancake 摊一个巨无霸煎饼
func (p *Pancake2Cook) MakeBigPancake() *Pancake2 {
	p.builder.PutPaste(Large)
	p.builder.PutEgg(3)
	p.builder.PutWafer()
	p.builder.PutFlavour(true, true, true)
	return p.builder.Build()
}

// MakePancakeForFlavour 摊一个自选调料霸煎饼
func (p *Pancake2Cook) MakePancakeForFlavour(hasCoriander, hasShallot, hasSauce bool) *Pancake2 {
	p.builder.PutPaste(Large)
	p.builder.PutEgg(3)
	p.builder.PutWafer()
	p.builder.PutFlavour(hasCoriander, hasShallot, hasSauce)
	return p.builder.Build()
}

func TestBuilderPattern(t *testing.T) {
	normal := NewNormalPancakeBuilder()

	// 创建一个煎饼师傅
	// 制作一个普通煎饼
	pancakeCook := NewPancakeCook(normal)
	pancake := pancakeCook.MakePancake()
	fmt.Println(pancake)

	// 修改煎饼师傅，制作一个健康煎饼
	healthy := NewHealthyPancakeBuilder()
	pancakeCook.SetPancakeBuilder(healthy)
	pancake = pancakeCook.MakePancake()
	fmt.Println(pancake)
}
