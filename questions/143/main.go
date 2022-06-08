package main

import "fmt"

type T struct{}

func (*T) foo() {
}

func (T) bar() {
}

type S struct {
	*T
}

func main() {
	s := S{}
	fmt.Printf("%#v\n", s)   // main.S{T:(*main.T)(nil)}
	fmt.Printf("%#v\n", s.T) // (*main.T)(nil)
	//fmt.Printf("%#v\n", *s.T) // panic: runtime error: invalid memory address or nil pointer dereference
	_ = s.foo
	s.foo()
	_ = s.bar      // panic: runtime error: invalid memory address or nil pointer dereference
	_ = (*s.T).bar // 因为 s.bar 将被展开为 (*s.T).bar，⽽ s.T 是个空指针，解引⽤会 panic。
}

// 空指针无法解引用，会 panic
