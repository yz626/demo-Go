package main

import "fmt"

type Person struct {
	Name     string
	Age      int
	*Address // 嵌套字段，通过这种方式，可以让Person结构体直接使用Address结构体的方法和属性
}

type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

func (a *Address) FullAddress() string {
	return a.Street + " " + a.City + " " + a.State + " " + a.Zip
}

func main() {
	p := Person{
		Name: "John Doe",
		Age:  30,
		Address: &Address{
			Street: "123 Main St",
			City:   "AnyTown",
			State:  "CA",
			Zip:    "12345",
		},
	}
	// Person类型可以直接访问Address类型的方法
	fmt.Println(p.FullAddress())
	// 实际上是通过p.Address.FullAddress()来调用的
	fmt.Println(p.Address.FullAddress())
}
