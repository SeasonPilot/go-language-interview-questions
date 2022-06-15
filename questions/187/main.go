package main

import "fmt"

func main() {
	count := 0
	for i := range [256]struct{}{} {
		m, n := byte(i), int8(i)
		if n == -n {
			count++
			fmt.Printf("i: %d , n: %d , -n: %d\n", i, n, -n)
		}
		if m == -m {
			count++
			fmt.Printf("m: %d\n", m)
		}
	}
	fmt.Println(count)
}

/*
	变量：
	无符号整型将会丢弃溢出的位
	对于带符号整数值来说，可以合法地溢出
	结果是 int8 类型，值是 128。超出了 int8 的范围，因为结果不是常量，允许溢出，128 的二进制表达式是 [1000 0000]，正好是 -128 的补码，因此答案是 -128。
*/
