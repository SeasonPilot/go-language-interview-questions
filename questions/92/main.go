package main

import "fmt"

type User struct{}
type User1 User
type User2 = User // 别名 完全等价于 User

func (i User1) m1() {
	fmt.Println("m1")
}

func (i User) m2() {
	fmt.Println("m2")
}

func main() {
	var i1 User1
	var i2 User2
	i1.m1()
	i2.m2()

	i1.m2() // 不能执⾏的，因为 User1 没有定义该⽅法。
}
