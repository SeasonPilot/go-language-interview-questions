package main

import "fmt"

func main() {
	s := make([]int)
	m := make(map[int]int)
	c := make(chan int)

	// 第一种用法，即缺少长度的参数，只传类型，这种用法只能用在类型为 map 或 chan 的场景，例如 make([]int) 是会报错的。
	// 这样返回的空间长度都是默认为 0 的。

	fmt.Println(s, m, c)
}
