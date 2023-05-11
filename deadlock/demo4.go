package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

// 不要在协程中使用 defer 语句来释放锁，因为这会导致协程在函数返回之前一直持有锁。如果其他协程需要获取同一个锁，就会发生死锁
// 两个 goroutine 抢同一把锁
func main() {
	go func() {
		mu.Lock()
		defer mu.Unlock()
		fmt.Println("goroutine 1")
	}()

	go func() {
		mu.Lock()
		defer mu.Unlock()
		fmt.Println("goroutine 2")
	}()

	fmt.Println("main")
	select {}
}
