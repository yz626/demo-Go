package main

import (
	"log"
	"time"
)

// 死锁分析：
// goroutine1: 持续向 in 输入
// 主循环中：进行中转，读取in ，写入out，或者处理错误通道的消息
// goroutine2: 持续从 out 输出
// 主要的原因在于，select和streamTextPreProcessStop
// 当select执行一个case时，如果在这个case中发生阻塞，那么select就会一直阻塞，导致死锁
// 代码中，当向streamTextPreProcessStop写入数据时，由于主循环中一直在处理in的读取
// 然后向out写入，但是在goroutine2中out的处理速度远小于in的发送
// 这就会导致主循环中会在向out写入的时候发生短时间的阻塞，直到goroutine2中处理完一个out
// 才能继续写入out，但是goroutine2中出发了向streamTextPreProcessStop写入
// 而streamTextPreProcessStop是一个无缓冲的channel
// 需要接收端准备好了才写入，否则就会发生阻塞
// 这里主循环中卡在了向out写入，streamTextPreProcessStop没有被处理
// 但是又只能写入完streamTextPreProcessStop后才能继续执行，导致死锁

func main() {
	testChannelChokeUp()
}

func testChannelChokeUp() {
	// 等待处理通道
	in := make(chan int, 20)
	// 处理通道
	out := make(chan int, 20)
	// 错误通道
	streamTextPreProcessStop := make(chan bool)

	// goroutine1: 持续向 in 输入
	go func() {
		for i := 0; i < 2000; i++ {
			log.Printf("in: %d", i)
			in <- i
		}
	}()

	// goroutine2: 持续从 out 输出
	// 并且当out元素为50时，触发错误
	go func() {
		for {
			select {
			case content, ok := <-out:
				log.Printf("out: %d", content)
				if !ok {
					log.Println("out channel error")
					break
				}
				time.Sleep(1 * time.Second)
				if content == 50 {
					// [这里会发生死锁]
					streamTextPreProcessStop <- true
				}
			}
		}
	}()

	// 主循环
	for {
		select {
		case <-streamTextPreProcessStop:
			log.Println("streamTextPreProcessStop")
		case content, ok := <-in:
			if !ok {
				log.Println("in channel error")
				break
			}
			out <- content
		}
	}
}
