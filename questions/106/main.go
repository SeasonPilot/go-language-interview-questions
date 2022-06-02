package main

import "fmt"

func main() {
	x := interface{}(nil)

	y := (*int)(nil)
	a := y == x
	b := y == nil

	fmt.Printf("%#v\n", x)
	fmt.Printf("%T\n", x)
	_, c := x.(interface{})

	println(a, b, c)

}

/*
	类型断⾔语法：i.(Type)，其中 i 是接⼝，Type 是类型或接⼝。编译时会⾃动检测 i 的动态类型与 Type
	是否⼀致。但是，如果动态类型不存在，则断⾔总是失败。
*/
