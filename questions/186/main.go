package main

import "fmt"

type TimesMatcher struct {
	base int
}

func NewTimesMatcher(base int) *TimesMatcher {
	return &TimesMatcher{base: base}
}

func main() {
	p := NewTimesMatcher(3)
	fmt.Println(p)
}

/*
	Go语⾔的内存回收机制规定，只要有⼀个指针指向引⽤⼀个变量，那么这个变量就不会被释放（内存逃逸），因此在 Go 语⾔中返回函数参数或临时变量是安全的。
*/
