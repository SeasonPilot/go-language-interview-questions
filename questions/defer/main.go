package main

import "fmt"

// https://caohuan.tech/2019/08/24/golang/defer%E7%9A%84%E6%89%A7%E8%A1%8C%E9%A1%BA%E5%BA%8F/
// defer的执行顺序
func main() {
	print(1)
	defer print(print(11) + print(12))
	defer print(3)
	print(4)
}

func print(i int) int {
	fmt.Println(i)
	return i
}
