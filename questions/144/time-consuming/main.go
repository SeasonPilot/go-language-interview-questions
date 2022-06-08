package main

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
)

// 关联 第 150 题
// 在函数调用前就已经完成求值。
func main() {
	//cost1()
	//cost2()
	//cost3()
	cost4()
}

func cost1() { // 错误
	logsParam := "[test]"
	startTime := time.Now()

	// 函数参数 在函数调用前就已经完成求值。
	defer logs.Warning("%s end-cost2 Seconds:%vs ", logsParam, time.Since(startTime).Seconds()) // defer 调用的函数 参数的值 defer 被定义时就确定了。

	time.Sleep(2 * time.Second)

	logs.Warning("%s end-cost2-2 Seconds:%vs ", logsParam, time.Since(startTime).Seconds())
}

func cost2() {
	logsParam := "[test]"
	startTime := time.Now()

	defer func(time1 time.Time) {
		logs.Warning("%s end-cost2 Seconds:%v\n", logsParam, time.Since(time1))
	}(startTime)

	time.Sleep(2 * time.Second)

	logs.Warning("%s end-cost2-2 Seconds:%vs ", logsParam, time.Since(startTime).Seconds())
}

func cost3() { // 错误
	t := time.Now()

	defer fmt.Printf("end-cost %#v\n", time.Since(t).Seconds()) // 错误  // defer 调用的函数Printf 参数的值 在 defer 被定义时就确定了。

	fmt.Println(time.Since(t).Seconds())

	time.Sleep(2 * time.Second)

}

func cost4() {
	t := time.Now()

	defer func() {
		fmt.Printf("end-cost Seconds: %#v\n", time.Since(t).Seconds())
	}()

	fmt.Println(time.Since(t).Seconds())

	time.Sleep(2 * time.Second)
}
