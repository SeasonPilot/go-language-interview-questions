package main

import "fmt"

type T struct {
	n int
}

func (t *T) Set(n int) {
	t.n = n
}

func getT() T {
	return T{}
}

func main() {
	getT().Set(1)
}

/*
	# awesomeProject2/go-language-interview-questions/130
	./main.go:14:8: cannot call pointer method on getT()
	./main.go:14:8: cannot take the address of getT()
*/

// 修复代码：
func main() {
	t := getT()
	fmt.Printf("%#v\n", t) // main.T{n:0}

	t.Set(2)
	fmt.Println(t.n)
}
