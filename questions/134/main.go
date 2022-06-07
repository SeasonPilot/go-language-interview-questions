package main

import "fmt"

type T struct {
	n int
}

func getT() T {
	return T{}
}

func main() {
	t := getT()
	p := &t.n // <=> p = &(t.n)   // 可直接使用 & 操作符取地址的对象，就是可寻址的(Addressable)。
	*p = 1
	fmt.Println(t.n)
}
