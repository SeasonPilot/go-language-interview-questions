package main

import "fmt"

func f(n int) (r int) {
	var f2 func()

	fmt.Printf("%#v\n", f2)
	fmt.Printf("%T\n", f2)
	f2() // nil

	defer f2() // 调用的 函数值 和参数都会 拷贝，拷贝的时候 f2 的值为 nil,所以引发异常
	f2 = func() {
		r += 2
	}
	return n + 1
}

func main() {
	fmt.Println(f(3))
}
