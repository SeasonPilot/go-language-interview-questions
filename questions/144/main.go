package main

import (
	"fmt"
	"sync"
	"time"
)

type data struct {
	sync.Mutex
}

// 将 Mutex 作为匿名字段时，相关的⽅法必须使⽤指针接收者，否则会导致锁机制失效。
func (d data) test(s string) {
	d.Lock()
	defer d.Unlock()

	for i := 0; i < 5; i++ {
		fmt.Println(s, i)
		time.Sleep(time.Second)
	}
}

func main() {
	t := time.Now()
	//defer fmt.Printf("%#v\n", time.Since(t).Seconds()) // 错误
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

/*
	write 0
	read 0


	read 1
	write 1


	read 2
	write 2


	write 3
	read 3


	read 4
	write 4
*/
