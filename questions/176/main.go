package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		fmt.Println("1")
		wg.Done()
		wg.Add(1)
	}()

	wg.Wait() // panic: sync: WaitGroup is reused before previous Wait has returned
}

// 协程⾥⾯，使⽤ wg.Add(1) 但是没有 wg.Done()，导致 panic()。
