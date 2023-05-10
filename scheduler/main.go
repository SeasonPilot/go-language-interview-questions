// 调度器示例：创建两个goroutine并交替打印数字
package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2) // 设置逻辑处理器为1
	go printNums(1)       // 创建第一个goroutine
	go printNums(2)       // 创建第二个goroutine
	for {                 // 主goroutine阻塞
	}
}

func printNums(g int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("goroutine %d: %d\n", g, i)
		runtime.Gosched() // 主动让出CPU给其他可运行的goroutine.每次打印一个数字后，当前goroutine就会暂停运行，让另一个goroutine继续运行。这样就可以实现交替打印数字的效果。
	}
}
