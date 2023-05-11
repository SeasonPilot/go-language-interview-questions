package main

/*
Go语言的内存模型保证基于happens-before关系，这是一种偏序关系，用于描述两个事件之间的因果关系。如果事件A happens-before事件B，
那么事件A对内存的修改对事件B可见，而且事件A和事件B不能被重排序

在这个示例中，f1函数对x变量的写入happens-before f2函数对x变量的读取，因为f1函数向通道发送信号happens-before f2函数从通道接收信号。
因此，f2函数可以看到f1函数对x变量的修改，打印出42。
*/

// 定义一个全局变量x
var x int

// main函数启动了两个goroutine：f1和f2
func main() {
	// 创建一个无缓冲的通道
	// 用作通知
	c := make(chan int)
	// 启动f1 goroutine
	go f1(c)
	// 启动f2 goroutine
	go f2(c)
	// 等待两个goroutine结束
	for {
	}
}

// f1函数向x变量写入42，然后向通道发送信号
func f1(c chan int) {
	// 写入x变量
	x = 42
	// 向通道发送信号
	c <- 0
}

// f2函数从通道接收信号，然后读取x变量
func f2(c chan int) {
	// 从通道接收信号
	<-c
	// 读取x变量
	print(x)
}
