package main

import (
	"fmt"
	"runtime"
)

func main() {
	var x int64 = 1           // x 分配在栈上
	var y *int64 = new(int64) // y 分配在堆上
	*y = 1
	var z *int64 = &x // z 分配在栈上，但 x 会发生逃逸
	fmt.Println(x, y, z)
	runtime.GC() // 手动触发垃圾回收

}
