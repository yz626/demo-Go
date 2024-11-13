package creationalPatterns

import "fmt"

// 抽象工厂模式
// 创建一个工厂，该工厂可以创建一系列产品，无需指定创建的具体产品。
// 抽象工厂定义了创建产品的接口，而具体工厂则实现了该接口。

// example：
// 创建两个产品族：Furniture 和 Vehicle。每个产品族包含多个具体的产品

// Furniture 接口
type Furniture interface {
	GetMaterial() string
}

// Vehicle 接口
type Vehicle interface {
	GetBrand() string
}

// WoodFurniture 实现 Furniture 接口
type WoodFurniture struct{}

func (w *WoodFurniture) GetMaterial() string {
	return "Wood"
}

// MetalFurniture 实现 Furniture 接口
type MetalFurniture struct{}

func (m *MetalFurniture) GetMaterial() string {
	return "Metal"
}

// Car 实现 Vehicle 接口
type Car struct{}

func (c *Car) GetBrand() string {
	return "Toyota"
}

// Bike 实现 Vehicle 接口
type Bike struct{}

func (b *Bike) GetBrand() string {
	return "Honda"
}

// AbstractFactory 抽象工厂接口
type AbstractFactory interface {
	CreateFurniture() Furniture
	CreateVehicle() Vehicle
}

// ModernFactory 现代工厂，实现 AbstractFactory 接口
type ModernFactory struct{}

func (mf *ModernFactory) CreateFurniture() Furniture {
	return &MetalFurniture{}
}

func (mf *ModernFactory) CreateVehicle() Vehicle {
	return &Car{}
}

// ClassicFactory 古典工厂，实现 AbstractFactory 接口
type ClassicFactory struct{}

func (cf *ClassicFactory) CreateFurniture() Furniture {
	return &WoodFurniture{}
}

func (cf *ClassicFactory) CreateVehicle() Vehicle {
	return &Bike{}
}

func TestAbstractFactoryPattern() {
	// 创建现代风格的工厂
	modernFactory := &ModernFactory{}
	modernFurniture := modernFactory.CreateFurniture()
	modernVehicle := modernFactory.CreateVehicle()

	fmt.Println("Modern Furniture Material:", modernFurniture.GetMaterial())
	fmt.Println("Modern Vehicle Brand:", modernVehicle.GetBrand())

	// 创建经典风格的工厂
	classicFactory := &ClassicFactory{}
	classicFurniture := classicFactory.CreateFurniture()
	classicVehicle := classicFactory.CreateVehicle()

	fmt.Println("Classic Furniture Material:", classicFurniture.GetMaterial())
	fmt.Println("Classic Vehicle Brand:", classicVehicle.GetBrand())
}
