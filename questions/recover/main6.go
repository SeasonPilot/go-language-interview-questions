package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err) // 多个panic只会捕捉最后一个
		}
	}()

	defer func() {
		panic("three")
	}()

	defer func() {
		panic("two")
	}()

	panic("one")
}

// https://www.runoob.com/note/38704
// 在调试程序时，通过 panic 来打印堆栈，方便定位错误。
