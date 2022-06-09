package main

import (
	"fmt"
	"time"
)

func mayPanic2() {
	defer func() {
		if r := recover(); r != nil { //（2）导致panic异常的函数不会继续运行，但是能正常返回，这意味着调用recover()的函数的上一级调用者可以继续往下执行代码，而不是中断程序运行。

			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	panic("a problem")

	fmt.Println("不可到达的代码 ")
}

func main() {

	mayPanic2()

	time.Sleep(time.Second * 2)
	fmt.Println("After mayPanic2()")
}
