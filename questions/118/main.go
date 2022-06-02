package main

import "fmt"

func main() {
	x := 1
	fmt.Println(x)

	{
		fmt.Println(x)
		i, x := 2, 2
		fmt.Println(i, x)
	}

	fmt.Println(x) // print ?
}

/*
	知识点：变量隐藏。
	使⽤变量简短声明符号 := 时，如果符号左边有多个变量，只需要保证⾄少有⼀个变量是新声明的，并对
	已定义的变量尽进⾏赋值操作。
	但如果出现作⽤域之后，就会导致变量隐藏的问题，就像这个例⼦⼀样。
*/
