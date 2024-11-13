package structuralPatterns

// 组合模式
// 结构型设计模式
// 使用它将对象组合成树状结构，并且能像使用独立对象一样使用它们

// IComponent 组件接口,假设为行政区为例
// 行政区之间组成的树状结构
type IComponent interface {
	Name() string
	GDP() float64
	Population() int
}

// City 城市，作为叶子节点，当然也可以是容器，例如：城市之间的组合
// 可以继续向下扩展另外的叶子的节点
// 只需要满足组件接口即可
type City struct {
	name       string
	gdp        float64
	population int
}

var _ IComponent = (*City)(nil)

func NewCity(name string, population int, gdp float64) *City {
	return &City{name, gdp, population}
}

func (c *City) Name() string {
	return c.name
}

func (c *City) GDP() float64 {
	return c.gdp
}

func (c *City) Population() int {
	return c.population
}

// Province 省份,作为容器
// 依旧可以向上继续扩展
// 作为其他更大的容器的一个节点
// 因为Province 实现了IComponent接口，因此可以作为其他容器的节点
type Province struct {
	name string
	city map[string]IComponent
}

var _ IComponent = (*Province)(nil)

func NewProvince(name string) *Province {
	return &Province{
		name: name,
		city: make(map[string]IComponent),
	}
}

func (p *Province) Name() string {
	return p.name
}

func (p *Province) GDP() float64 {
	var res = 0.0
	for _, item := range p.city {
		res += item.GDP()
	}
	return res
}

func (p *Province) Population() int {
	var res = 0
	for _, item := range p.city {
		res += item.Population()
	}
	return res
}

// Add 组装, 假设向Province中添加多个City
// 这里参数使用IComponent组件接口
func (p *Province) Add(c ...IComponent) {
	for _, item := range c {
		p.city[item.Name()] = item
	}
}

// Remove 删除
// 这里需要使用递归删除
// 由于Province的city字段存储的是IComponent接口类型
// 因此也有可能是Province或者其他容器
// 要递归便利每个容器查找删除
func (p *Province) Remove(name string) {
	for itemName, itemCity := range p.city {
		if itemName == name {
			delete(p.city, itemName)
			return
		}
		// 如果是Province类型，则递归查找删除
		if item, ok := itemCity.(*Province); ok {
			item.Remove(name)
		}
	}
}
