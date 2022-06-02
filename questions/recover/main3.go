package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获异常:", fmt.Errorf("%v", err).Error())
		}
	}()
	panic("a")
}

/*
	6、recover返回的是interface{}类型而不是go中的 error 类型，如果外层函数需要调用err.Error()，会编译错误，也可能会在执行时panic
*/
