package main

import "fmt"

func main() {
	var b bool
	var a int

	b = true
	b = 1
	b = bool(1)
	b = bool(a)
	b = 1 == 2

	fmt.Printf("%T\n%#v", b, b)
}
