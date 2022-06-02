package main

import "fmt"

// 50. 下⾯的两个切⽚声明中有什么区别？哪个更可取？
func main() {
	var a []int
	b := []int{}

	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", b)
	// A 声明的是 nil 切⽚；B 声明的是⻓度和容量都为 0 的空切⽚。
	// 第⼀种切⽚声明不会分配内存，优先选择。
}
