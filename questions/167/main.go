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
	defer func() {
		s.Add(1).Add(2)
	}()
	s.Add(3)
}

/*
	对⽐昨天的第 166 题，本题的 s.Add(1).Add(2) 作为⼀个整体包在⼀个匿名函数中，会延迟执⾏。
*/
