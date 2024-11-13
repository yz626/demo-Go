package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 创建链表
	ll := list.New()
	ll.PushBack("hello")
	ll.PushBack("world")
	fmt.Println(ll)
	fmt.Println(ll.Len())
}
