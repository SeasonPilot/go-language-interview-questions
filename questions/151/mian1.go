package main

import "fmt"

func main() {
	defer fmt.Println(recover()) // <nil>

	panic(1)
}

// recover() 必须在 defer() 函数中直接调⽤才有效。
