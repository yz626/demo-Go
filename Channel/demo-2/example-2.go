package main

func main() {
	numChan := make(chan int, 5)
	strChan := make(chan string, 1)

	go func() {
		for i := 0; i < 5; i++ {
			numChan <- i
		}
		strChan <- "hello"
	}()

	// 目前这段代码输出是随机的，并不是按照发送的顺序输出的
	// numChan和strChan是两个独立队列,但是都已经准备完毕
	// 现在我想要实现数字的优先级高于字符串，如何？
	// 由于select会在已经准备读取的队列中随机选择一个，无法确定哪一个先输出
	// 我们在低优先级（strChan）进行处理的时候，先读取高优先级（numChan）进行处理
	// 直到高优先级队列中的数据全部处理完，再处理低优先级队列
	//for {
	//	select {
	//	case num := <-numChan:
	//		println(num)
	//	case str := <-strChan:
	//		println(str)
	//		return
	//	}
	//}

	// 改写为如下
	for {
		select {
		case num := <-numChan:
			println(num)
		case str := <-strChan:
		goal:
			// 将高优先级任务处理完之后再处理低优先级
			for {
				// default 语句只有在所有的 case 语句都没有准备好的时候才会执行
				// 因此 这里会执行外面因随机选取错过的已经准备就绪的高优先级任务
				select {
				case num := <-numChan:
					println(num)
				default:
					println("跳出")
					// 跳出内层for循环
					break goal
				}
			}
			// 执行原本的低优先级任务
			println(str)
			return
		}
	}
}
