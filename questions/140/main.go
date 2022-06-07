package main

import "fmt"

func main() {
	x := make([]int, 2, 10)
	_ = x[6:10]

	s := x[1:]
	fmt.Printf("%#v\n", len(s)) // 1  截取符号 [i:j]，如果 j 省略，默认是原切⽚或者数组的⻓度
	fmt.Printf("%#v\n", cap(s)) // 9  如果 j 省略，容量为 cap(baseSlice) - i

	s1 := x[2:]
	fmt.Printf("%#v\n", s1)      // []int{}
	fmt.Printf("%#v\n", len(s1)) // 0
	fmt.Printf("%#v\n", cap(s1)) // 8

	_ = x[6:] // 截取符号 [i:j]，如果 j 省略，默认是原切⽚或者数组的⻓度，x 的⻓度是 2，⼩于起始下标 6 ，所以 panic。 // slice bounds out of range [6:2]
	_ = x[2:]
}
