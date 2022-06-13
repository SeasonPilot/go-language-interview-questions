package main

func test() []func() {
	var funs []func()
	for i := 0; i < 2; i++ {
		funs = append(funs, func() {
			println(&i, i)
		})
	}
	return funs
}

func main() {
	funs := test()
	for _, f := range funs {
		f()
	}
}

/*
	知识点：闭包延迟求值。for 循环局部变量 i，匿名函数每⼀次使⽤的都是同⼀个变量。
*/
