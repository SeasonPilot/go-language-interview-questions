package main

import "fmt"

type T struct {
	n int
}

func main() {
	m := make(map[int]T)

	m[0].n = 1

	fmt.Println(m[0].n)
}

/*
	map[key]struct 中 struct 是不可寻址的，所以⽆法直接赋值。
*/

// 修复代码：
func main() {
	m := make(map[int]T)

	t := T{1}
	m[0] = t

	fmt.Println(m[0].n)
}
