package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()
	for {
	}
}

/*
	答：for {} 独占 CPU 资源导致其他 Goroutine 饿死
*/

// 可以通过阻塞的⽅式避免 CPU 占⽤，修复代码：
func main() {
	runtime.GOMAXPROCS(1)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		os.Exit(0)
	}()
	select {}
}
