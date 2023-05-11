package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup // 创建一个 WaitGroup 变量

	wg.Add(3) // 向内部计数器添加 3

	go func() {
		defer wg.Done() // 在协程结束时调用 Done()
		time.Sleep(2 * time.Second)
		fmt.Println("Goroutine 1 finished")
	}()

	go func() {
		defer wg.Done() // 在协程结束时调用 Done()
		time.Sleep(1 * time.Second)
		fmt.Println("Goroutine 2 finished")
	}()

	go func() {
		defer wg.Done() // 在协程结束时调用 Done()
		time.Sleep(3 * time.Second)
		fmt.Println("Goroutine 3 finished")
	}()

	wg.Wait() // 阻塞直到计数器为 0
	fmt.Println("All goroutines done")
}
