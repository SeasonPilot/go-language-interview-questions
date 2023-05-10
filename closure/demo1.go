package main

import (
	"fmt"
	"time"
)

// 定义一个定时器，接受一个函数作为参数，该函数可以访问定时器内的变量和方法
func createTimer(callback func(stop func())) {
	// 定义一个变量，用于存储定时器的状态
	var running = true
	// 定义一个方法，用于停止定时器
	var stop = func() {
		running = false
	}
	// 启动一个协程，每隔一秒执行一次回调函数，并传入停止方法作为参数
	go func() {
		for running {
			callback(stop)
			time.Sleep(time.Second)
		}
	}()
}

func main() {
	// 定义一个变量，用于存储计数器的值
	var count = 0
	// 调用定时器，传入一个闭包作为回调函数，该闭包可以访问计数器的值和停止方法
	createTimer(func(stop func()) {
		// 每次执行回调函数时，增加计数器的值并输出
		count++
		fmt.Println(count)
		// 如果计数器达到十次，则停止定时器
		if count == 10 {
			stop()
		}
	})
	// 等待定时器结束后退出程序
	time.Sleep(11 * time.Second)
}
