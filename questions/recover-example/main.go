package main

import (
	"fmt"
)

// https://www.cnblogs.com/wujuntian/p/11747113.html
func main() {
	defer func() { fmt.Println(1) }()

	fmt.Println(2)

	defer func() { fmt.Println(3) }()

	defer_call1()

	fmt.Println(4)
}

func defer_call1() {
	defer func() { fmt.Println("A") }()

	fmt.Println("B")

	defer func() {
		fmt.Println("C")
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("D")
	}()

	defer func() { fmt.Println("E") }()

	defer_call2()
	fmt.Println("F")
}

func defer_call2() {
	defer func() { fmt.Println("a") }()
	defer func() { fmt.Println("b") }()
	defer func() { fmt.Println("c") }()

	fmt.Println("d")

	panic("触发异常1")

	panic("触发异常2")
	fmt.Println("e")
}
