package main

import (
	"fmt"
)

var gvar int

func main() {
	var one int
	_ = one

	two := 2
	fmt.Println(two)
	var three int
	three = 3
	one = three

	var four int
	four = four

	func(unused string) {
		fmt.Println("Unused arg. No compile error")
	}("what?")

}
