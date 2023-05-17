package main

// 如果一个函数返回对一个变量的引用，例如：
func main() {
	var y *int64 = foo() // y 分配在堆上，接收了对 x 的引用
}

func foo() *int64 {
	var x int64 = 1 // x 分配在堆上，发生逃逸
	return &x       // 返回对 x 的引用
}
