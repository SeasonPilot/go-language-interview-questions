package main

import "fmt"

func f(n int) (r int) {
	defer func() {
		r += n
		recover()
	}()

	var f2 func()

	defer f2() // 调用的 函数值 和参数都会 拷贝，拷贝的时候 f2 的值为 nil,所以引发异常
	f2 = func() {
		r += 2
	}
	return n + 1
}

func main() {
	fmt.Println(f(3))
}

/*
	解析：
	第⼀步执⾏ r = n +1 ，接着执⾏第⼆个 defer，由于此时 f2() 未初始化，引发异常panic ，随即执⾏第⼀个
	defer，异常被 recover()，程序正常执⾏，最后 return。
*/
