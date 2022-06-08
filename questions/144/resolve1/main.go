package main

import (
	"fmt"
	"sync"
	"time"
)

type data struct {
	sync.Mutex
}

func (d *data) test(s string) { // 指针接收者
	d.Lock()
	defer d.Unlock()
	for i := 0; i < 5; i++ {
		fmt.Println(s, i)
		time.Sleep(time.Second)
	}
}

func main() {
	t := time.Now()

	defer func() {
		fmt.Printf("%#v\n", time.Since(t).Seconds())
	}()

	var wg sync.WaitGroup
	wg.Add(2)

	var d data

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
