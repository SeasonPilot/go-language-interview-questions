package main

import "fmt"

const (
	a = iota
	b = iota
)
const (
	name = "name"
	c    = iota
	d    = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

/*
	知识点：iota 的⽤法。

	iota 是 golang 语⾔的常量计数器，只能在常量的表达式中使⽤。
	iota 在 const 关键字出现时将被重置为0，const中每新增⼀⾏常量声明将使 iota 计数⼀次。
*/
