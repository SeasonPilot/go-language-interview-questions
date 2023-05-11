package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		//close(ch) // 关闭channel
	}()

	for {
		v := <-ch // 从channel中读取数据
		fmt.Println(v)
	}
}
