package main

import (
	"fmt"
	"time"
)

func main() {
	var m = [...]int{1, 2, 3}

	for i, v := range m {
		go func() {
			fmt.Println(i, v)
		}()
	}

	time.Sleep(time.Second * 3)
}

/*
	解析：
	for range 使⽤短变量声明(:=)的形式迭代变量，需要注意的是，变量 i、v 在每次循环体中都会被重⽤，⽽不是重新声明。
	各个 goroutine 中输出的 i、v 值都是 for range 循环结束后的 i、v 最终值，⽽不是各个goroutine启动时的i, v值。可以理解为闭包引⽤，使⽤的是上下⽂环境的值。
*/

// 两种可⾏的 fix ⽅法:
// 1. 使⽤函数传递
func main1() {
	var m = [...]int{1, 2, 3}

	for i, v := range m {
		go func(i, v int) {
			fmt.Println(i, v)
		}(i, v)
	}
}

// 2. 使⽤临时变量保留当前值
func main2() {
	var m = [...]int{1, 2, 3}

	for i, v := range m {
		i := i // 这⾥的 := 会重新声明变量，⽽不是重⽤
		v := v
		go func() {
			fmt.Println(i, v)
		}()
	}
}
