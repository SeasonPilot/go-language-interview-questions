package main

import "fmt"

func main() {
	defer func() {
		fmt.Print(recover())
	}()

	defer func() { // 每次执行“defer”语句时，调用的函数值和参数都会 像往常一样评估 并重新保存，但不会调用实际函数。
		defer fmt.Print(recover()) // 在调⽤ defer() 时，便会计算函数的参数并压⼊栈中，所以执⾏第 6 ⾏代码时，此时便会捕获 panic(2)
		panic(1)
	}()

	defer recover() // recover() 必须在 defer() 函数中调⽤才有效，所以这⾏代码捕获是⽆效的。

	panic(2)
}

/*
	解析：
	recover() 必须在 defer() 函数中调⽤才有效，所以第 9 ⾏代码捕获是⽆效的。
	在调⽤ defer() 时，便会计算函数的参数并压⼊栈中，所以执⾏第 6 ⾏代码时，此时便会捕获 panic(2)；
	此后的 panic(1)，会被上⼀层的 recover() 捕获。所以输出 21。
*/
