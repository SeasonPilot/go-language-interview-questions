package main

import (
	"fmt"
)

func main() {
	ch := make(chan int) // 创建一个无缓冲的channel
	go func() {
		fmt.Println("Hello") // 打印Hello
		// 没有向channel发送数据或者关闭channel
	}()

	<-ch // 从channel接收数据，阻塞main函数
}
