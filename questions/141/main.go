package main

import "fmt"

// 关联 136 题

type N int

func (n *N) test() { // 当⽬标⽅法 的 接收者 是 指针类型时，那么被复制的就是指针。
	fmt.Println(*n)
}

func main() {
	var n N = 10
	p := &n

	n++
	f1 := n.test  // 语法糖， Go 语言会在编译阶段会将它转换成 (&n).test()
	(*N).test(&n) // 11 // &n 就是方法接收者

	n++
	f2 := p.test
	(*N).test(p) // 12

	n++
	fmt.Println(n) // 13

	f1() // 13
	f2() // 13
}
