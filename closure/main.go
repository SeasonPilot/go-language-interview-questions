package main

import "fmt"

func main() {
	x := 10
	f := func() {
		fmt.Println(x) // 使用外部变量x
	}
	g := func() {
		fmt.Println("Hello") // 不使用外部变量
	}
	fmt.Printf("%T\n", f) // 打印f的类型
	fmt.Printf("%T\n", g) // 打印g的类型
}
