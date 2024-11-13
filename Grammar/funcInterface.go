package main

import (
	"fmt"
)

// 接口型函数只能应用于接口内部只有一个方法的情况。
// 接口型函数的作用：
// 1.既能够接收函数作为参数，也能够接收实现接口类型的结构体作为参数。
// 2.

type Getter interface {
	Get(key string) ([]byte, error)
}

// Handle 函数类型
type Handle func(key string) ([]byte, error)

// Get 实现 Getter 接口，通过函数类型实现
func (h Handle) Get(key string) ([]byte, error) {
	return h(key)
}

// Cache 实现 Getter 接口
type Cache struct {
	data map[string]interface{}
}

func (c *Cache) Get(key string) ([]byte, error) {
	if bytes, ok := c.data[key]; ok {
		return bytes.([]byte), nil
	}
	return nil, fmt.Errorf("key %s not exist", key)
}

func main() {
	// 函数类型实现了 Getter 接口
	var g Getter = Handle(func(key string) ([]byte, error) {
		return []byte(key), nil
	})
	// 通过接口调用
	result, err := g.Get("key")
	if err != nil {
		return
	}
	fmt.Println(string(result))

	// 结构体实现 Getter 接口
	var c Getter = &Cache{data: make(map[string]interface{})}
	result, err = c.Get("key")
	if err != nil {
		return
	}
	fmt.Println(string(result))
}
