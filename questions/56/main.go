package main

import "fmt"

// 56. 下⾯选项正确的是？
func main() {
	if a := 1; false {
	} else if b := 2; false {
	} else {
		println(a, b)
	}

	if a := 1; false {
		fmt.Println(111111111)
	} else if b := 2; true {
		fmt.Println(222222222)
	} else {
		println(a, b)
	}
}
