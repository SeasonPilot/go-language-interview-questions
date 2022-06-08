package main

import "fmt"

func main() {
	nil := 123
	fmt.Println(nil)
	var _ map[string]int = nil // 当前作⽤域中，预定义的 nil 被覆盖，此时 nil 是 int 类型值，不能赋值给 map 类型。
}
