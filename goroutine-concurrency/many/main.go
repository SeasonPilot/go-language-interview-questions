package main

import (
	"fmt"
	"sync"
	"time"
)

type Glimit struct {
	n int
	c chan struct{}
}

// initialization Glimit struct
func New(n int) *Glimit {
	return &Glimit{
		n: n,
		c: make(chan struct{}, n),
	}
}

// Run f in a new goroutine but with limit.
func (g *Glimit) Run(f func()) {
	g.c <- struct{}{}
	go func() {
		f()
		<-g.c
	}()
}

var wg = sync.WaitGroup{}

func main() {
	number := 10 // 一共起多少 goroutine
	g := New(2)  // 并发启动多少 goroutine

	for i := 0; i < number; i++ {
		wg.Add(2)

		value := i

		goFunc := func() {
			// 做一些业务逻辑处理
			fmt.Printf("go func 1: %d\n", value)
			time.Sleep(time.Second)
			wg.Done()
		}

		goFunc2 := func() {
			// 做一些业务逻辑处理
			fmt.Printf("go func 2: %d\n", value)
			time.Sleep(time.Second)
			wg.Done()
		}

		g.Run(goFunc)
		g.Run(goFunc2)
	}

	wg.Wait()
}
