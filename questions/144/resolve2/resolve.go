package main

import (
	"fmt"
	"sync"
	"time"
)

// 通过嵌⼊ *Mutex 来避免复制的问题，但需要初始化。
type data struct {
	*sync.Mutex // *Mutex
}

func (d data) test(s string) { // 值⽅法
	d.Lock()
	defer d.Unlock()

	for i := 0; i < 5; i++ {
		fmt.Println(s, i)
		time.Sleep(time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	d := data{new(sync.Mutex)} // 初始化

	go func() {
		defer wg.Done()
		d.test("read")
	}()

	go func() {
		defer wg.Done()
		d.test("write")
	}()

	wg.Wait()
}
