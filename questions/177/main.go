package main

import "fmt"

func main() {
	arr := [2]int{1, 2}
	fmt.Printf("%#v\n", cap(arr))
}

/*
	cap() 函数的作⽤是：
		array 返回数组的元素个数
		slice 返回slice的最⼤容量
		channel 返回 channel的容量
*/
