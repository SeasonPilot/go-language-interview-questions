package main

import "fmt"

// 如果一个变量的指针被多个函数或线程引用，例如：
func main() {
	var x int64 = 1   // x 分配在栈上
	var y *int64 = &x // y 分配在堆上，x 发生逃逸
	go func() {
		fmt.Println(*y) // 一个线程引用了 y
	}()
	go func() {
		*y = 2 // 另一个线程也引用了 y
	}()
}
