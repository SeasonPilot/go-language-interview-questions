package main

import "fmt"

var p *int

func foo() (*int, error) {
	var i int = 5
	return &i, nil
}

func bar() {
	//use p
	fmt.Println(*p)
}

func main() {
	p, err := foo() // 对于本例来说，main() 函数⾥的 p 是新定义的变量，会遮住全局变量 p，导致执⾏到 bar() 时程序，全局变量 p 依然还是 nil，程序随即 Crash。
	if err != nil {
		fmt.Println(err)
		return
	}

	bar()
	fmt.Println(*p)
}

/*
	知识点：变量作⽤域。
	问题出在操作符 := ，对于使⽤ := 定义的变量，如果新变量与同名已定义的变量不在同⼀个作⽤域中，那么 Go 会新定义这个变量。
	对于本例来说，main() 函数⾥的 p 是新定义的变量，会遮住全局变量 p，导致执⾏到 bar() 时程序，全局变量 p 依然还是 nil，程序随即 Crash。
*/

