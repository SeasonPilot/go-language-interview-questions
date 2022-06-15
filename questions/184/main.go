package main

import (
	"fmt"
	"sync"
)

// 错误
func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		go func(wg sync.WaitGroup, i int) { // 使⽤了 sync.WaitGroup 副本；
			wg.Add(1) // 在协程中使⽤ wg.Add()；
			fmt.Printf("i:%d\n", i)
			wg.Done()
		}(wg, i)
	}

	wg.Wait()
	fmt.Println("exit")
}

/*
	解析：
	知识点：WaitGroup 的使⽤。存在两个问题：
	在协程中使⽤ wg.Add()；
	使⽤了 sync.WaitGroup 副本；
*/

// 修复代码：
func main() { // 使用全局变量
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Printf("i:%d\n", i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("exit")
}

//或者：
func main() { // 传递指针
	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			fmt.Printf("i:%d\n", i)
			wg.Done()
		}(wg, i)
	}

	wg.Wait()
	fmt.Println("exit")
}
