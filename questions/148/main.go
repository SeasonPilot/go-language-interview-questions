package main

import "fmt"

func main() {
	var x int8 = -128
	var y = x / -1
	fmt.Println(y) // -128  溢出
}

// 结果是 int8 类型，值是 128。超出了 int8 的范围，因为结果不是常量，允许溢出，128 的二进制表达式是 [1000 0000]，正好是 -128 的补码，因此答案是 -128。
