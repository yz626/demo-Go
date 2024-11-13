package main

import "fmt"

// Message 信息
type Message string

func NewMessage() Message {
	return Message("Hi there!")
}

// Greeter 问候
type Greeter struct {
	Message Message
}

// NewGreeter 创建一个 Greeter，需要依赖 Message对象
func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

func (g Greeter) Greet() Message {
	return g.Message
}

// Event 事件
type Event struct {
	Greeter Greeter
}

// NewEvent 创建一个 Event，需要依赖 Greeter对象
func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

// 创建一个 Event对象，需要依赖Greeter对象,而创建Greeter对象需要依赖 Message对象，
// 这之间具有依赖关系，直接通过创建一个Message对象然后再创建一个 Greeter对象，
// 最后创建一个 Event对象，这种方式过于麻烦,通过定义额外的封装函数，解决创建依赖问题。

func main() {
	event := InitializeEvent()

	event.Start()
}
