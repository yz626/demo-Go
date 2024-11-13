package DesignPattern

import "fmt"

// 函数选项模式
// 针对于配置选项过多的场景，使用函数选项模式

// example-1:

// Options 函数选项
// 这里可以进一步封装，将函数选项封装成一个结构体
// 再定义一个接口，将接口对外暴露，并且将函数选项作为接口中函数的参数
type Options func(u *User)

func WithName(name string) Options {
	return func(u *User) {
		u.Name = name
	}
}

func WithAge(age int) Options {
	return func(u *User) {
		u.Age = age
	}
}

func WithGender(gender string) Options {
	return func(u *User) {
		u.Gender = gender
	}
}

func WithAddress(address string) Options {
	return func(u *User) {
		u.Address = address
	}
}

type User struct {
	Name    string
	Age     int
	Gender  string
	Address string
}

// NewUser 构造函数函数
// 将函数选项作为参数，只需要传入需要配置的选项
func NewUser(options ...Options) *User {
	u := &User{
		Name:    "default",
		Age:     18,
		Gender:  "male",
		Address: "default",
	}

	for _, option := range options {
		option(u)
	}

	return u
}

// NewUser 函数
// 但是该构造函数参数过长，创建时候需要传递很多参数
// 过于复杂
//func NewUser(name string, age int, gender string, address string) *User {
//	return &User{
//		Name:    name,
//		Age:     age,
//		Gender:  gender,
//		Address: address,
//	}
//}

func main() {
	// 只需要传入需要配置的选项
	user := NewUser(WithName("lll"), WithGender("female"))
	fmt.Println(user)
}
