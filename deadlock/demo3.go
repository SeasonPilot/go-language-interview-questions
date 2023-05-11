package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2) // 创建一个有缓冲区的channel，容量为2
	ch <- 42                // 向channel发送数据
	fmt.Println("Sent to channel", 42)
	ch <- 43 // 向channel发送数据
	fmt.Println("Sent to channel", 43)
	// 此时缓冲区已经满了
	close(ch) // 关闭channel
	v := <-ch // 从channel接收数据
	fmt.Println("Received from channel", v)
	v = <-ch // 从channel接收数据
	fmt.Println("Received from channel", v)
	v = <-ch // 从channel接收数据
	fmt.Println("Received from channel", v)
	// 这一行代码会执行，因为从关闭的channel接收数据会得到零值，而不会阻塞。但是这样就造成了内存泄漏，因为没有goroutine来清理channel中剩余的数据。

}
