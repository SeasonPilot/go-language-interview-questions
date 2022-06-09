package main

import "fmt"

func main() {
	defer func() { // 每次执行“defer”语句时，调用的函数值和参数都会 像往常一样评估 并重新保存，但不会调用实际函数。
		defer fmt.Println(recover()) // 计算 recover() 值为 2 // 在调⽤ defer() 时，便会计算函数的参数并压⼊栈中，所以执⾏第 6 ⾏代码时，此时便会捕获 panic(2)
		panic(1)
	}()

	panic(2)
}
