package main

import "fmt"

// 59. 下⾯这段代码输出什么？为什么？

func (i int) PrintInt() {
	fmt.Println(i)
}

func main() {
	var i int = 1
	i.PrintInt()
}

/*
	解析：
	基于类型创建的⽅法必须定义在同⼀个包内，上⾯的代码基于 int 类型创建了 PrintInt() ⽅法，由于 int
	类型和⽅法 PrintInt() 定义在不同的包内，所以编译出错。
*/
