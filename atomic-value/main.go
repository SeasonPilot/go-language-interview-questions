package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var v atomic.Value
	v.Store(1)            // store an int value
	fmt.Println(v.Load()) // load the int value
	v.Store("hello")      // store a string value, this will cause panic
	fmt.Println(v.Load()) // this will not be executed
}
