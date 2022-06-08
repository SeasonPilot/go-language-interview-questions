package main

import "fmt"

func main() {
	var k = 9
	for k = range []int{} {
		fmt.Println("AAAAAAAAAAA") // 不会进入循环
	}
	fmt.Println(k) // 9

	for k = 0; k < 3; k++ {
		fmt.Println("BBBBBBBBBB")
	}
	fmt.Println(k) // 3

	for k = range (*[3]int)(nil) {
		fmt.Printf("%#v\n", k)
		fmt.Println("CCCCCCCCC")
	}
	fmt.Println(k) // 2
}
