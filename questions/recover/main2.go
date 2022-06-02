package main

import (
	"fmt"
	"time"
)

func G() {
	defer func() { // 其它协程里挂的 defer 语句不作保证
		fmt.Println("执行不到!!!")
		//goroutine外进行recover
		if err := recover(); err != nil {
			fmt.Println("捕获异常:", err)
		}
		fmt.Println("c")
	}()

	//创建goroutine调用F函数
	go F()

	time.Sleep(time.Second)
}

func F() {
	defer func() {
		fmt.Println("b")
	}()

	//goroutine内部抛出panic
	panic("a")
}

func main() {
	G()
}

// panic 会停掉当前正在执行的程序，不只是当前协程。在这之前，它会有序地执行完当前协程 defer 列表里的语句，其它协程里挂的 defer 语句不作保证。
/*
	5、recover都是在当前的goroutine里进行捕获的，这就是说，对于创建goroutine的外层函数，
	如果goroutine内部发生panic并且内部没有用recover，外层函数是无法用recover来捕获的，这样会造成程序崩溃

	链接：https://www.jianshu.com/p/63e3d57f285f
*/
