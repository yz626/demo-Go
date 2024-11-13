package creationalPatterns

import (
	"fmt"
	"sync"
)

// 单例模式
// 保证一个类仅有一个实例，并提供一个访问它的全局访问点
// 单例拥有与全局变量相同的优缺点
// 尽管它们非常有用，但却会破坏代码的模块化特性

// example:
// 创建一个地球对象，无法导出，通过特定方法访问。

type earth struct {
	desc string
}

func (e *earth) String() string {
	return e.desc
}

// theEarth 唯一的地球对象
var theEarth *earth

// TheEarth 获取唯一的地球对象
func TheEarth() *earth {
	if theEarth == nil {
		once := sync.Once{}
		once.Do(func() {
			theEarth = &earth{desc: "地球"}
		})
	}
	return theEarth
}

func TestSingleton() {
	fmt.Println(TheEarth().String())
}
