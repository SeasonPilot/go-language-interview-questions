package main

import "fmt"

type Person struct {
	name   string
	height string
}

func main() {
	p := &Person{"张三", "一米八"}

	p2 := Person{"李四", "一米七"}

	fmt.Println("打印指针类型变量: ", p)
	fmt.Println("打印值类型变量: ", p2)
}
