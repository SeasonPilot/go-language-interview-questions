package main

import "fmt"

type X struct{}

func (x *X) test() {
	println(x)
}

func main() {
	var a *X
	a.test()
	X{}.test() // X{} 是不可寻址的，不能直接调⽤⽅法   // 字面量通通不能寻址
}

// 修复代码
func main() {
	var a *X
	a.test() // 相当于 test(nil)

	var x = X{}
	fmt.Printf("%p\n", x)
	x.test()
}

func main() {
	var a *X
	fmt.Printf("%#v\n", a)
	a.test()

	fmt.Printf("%#v\n", X{})
	(&X{}).test()
}
