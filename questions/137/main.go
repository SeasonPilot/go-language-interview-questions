package main

import "fmt"

func main() {
	var x interface{}
	var y interface{} = []int{3, 5}
	fmt.Printf("%#v, %T\n", x, x)           // <nil>, <nil>
	fmt.Printf("%#v, %T\n", y, y)           // []int{3, 5}, []int
	fmt.Printf("%#v,%#v\n", x == x, x == y) // true,false
	_ = x == x                              // true
	_ = x == y                              // false
	_ = y == y                              // panic: runtime error: comparing uncomparable type []int  //因为两个⽐较值 的 动态类型 为 同⼀个 不可⽐较类型。
}
