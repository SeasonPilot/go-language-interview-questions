package main

import "fmt"

func main() {
	s := [3]int{1, 2, 3}
	fmt.Printf("长度:%v, 容量: %v\n", len(s), cap(s))

	s2 := [3]int{1}
	fmt.Printf("长度:%v, 容量: %v\n", len(s2), cap(s2))

	var s3 [4]int
	fmt.Printf("长度:%v, 容量: %v\n", len(s3), cap(s3))
	fmt.Println(s3)

	// An array's length is part of its type
	// var s3 [4]int  方括号内为数组的 长度
}
