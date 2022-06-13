package main

import "fmt"

type Slice []int

func NewSlice() Slice {
	return make(Slice, 0)
}

func (s *Slice) Add(elem int) *Slice {
	*s = append(*s, elem)
	fmt.Print(elem)
	return s
}

func main() {
	s := NewSlice()
	defer s.Add(1).Add(2)
	s.Add(3)
}

/*
	这⼀题有两点需要注意：
	1. Add() ⽅法的返回值依然是指针类型 *Slice，所以可以循环调⽤⽅法 Add()；
	2. defer 函数的参数（包括接收者）是在 defer 语句出现的位置做计算的，⽽不是在函数执⾏的时候
	计算的，所以 s.Add(1) 会先于 s.Add(3) 执⾏。
*/
