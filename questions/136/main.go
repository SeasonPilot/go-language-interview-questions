package main

import "fmt"

// 关联 141 题

type N int

func (n N) test() {
	fmt.Println(n)
}

func main() {
	var n N = 10
	p := &n

	n++
	f1 := n.test

	n++
	a := *p
	f2 := p.test // 当 指针值 赋值给 变量 或者作为函数 参数 传递时，会⽴即计算并复制该⽅法执⾏所需的接收者对象，与其绑定，以便在稍后执⾏时，能隐式第传⼊接收者参数。
	N.test(*p)   // 12

	n++
	fmt.Println(n) // 13

	f1() // 11
	f2() // 12

	fmt.Println(a) // 12
}
