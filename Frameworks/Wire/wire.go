//go:build wireinject

package main

// InitializeEvent 注入依赖
func InitializeEvent() Event {
	// 通过 wire.Build 构建依赖
	// 然后执行Wire命令，生成wire_gen.go文件
	panic(wire.Build(NewEvent, NewGreeter, NewMessage))
	return Event{}
}
