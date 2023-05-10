package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	m.Store("hello", "world")
	m.Store(42, "answer")
	value, ok := m.Load("hello")
	if ok {
		fmt.Println(value) // world
	}
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
	// hello world
	// 42 answer
	m.Delete(42)
	value, ok = m.Load(42)
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("not found") // not found
	}
}
