package main

import "fmt"

type Orange struct {
	Quantity int
}

func (o *Orange) Increase(n int) {
	o.Quantity += n
}

func (o *Orange) Decrease(n int) {
	o.Quantity -= n
}

func (o *Orange) String() string {
	return fmt.Sprintf("%#v", o.Quantity)
}

func main() {
	var orange Orange
	orange.Increase(10)
	orange.Decrease(5)
	fmt.Println(orange) // {5}
}

/*
	这道题容易忽视的点是，String() 是指针⽅法，⽽不是值⽅法，所以使⽤ Println() 输出时不会调⽤到 String() ⽅法。
*/

// 修复代码：
func main() {
	orange := &Orange{}
	orange.Increase(10)
	orange.Decrease(5)
	fmt.Println(orange) // 5
}
