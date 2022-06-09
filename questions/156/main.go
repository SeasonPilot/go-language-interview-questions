package main

import "fmt"

type T struct {
	n int
}

func main() {
	ts := [2]T{}

	for i, t := range &ts { // 循环的是原数组的指针，
		switch i {
		case 0:
			t.n = 3 // 不能修改，可以访问

			fmt.Printf("%T\n", t)

			ts[1].n = 9 // 数组操作，修改原数组
			fmt.Printf("%#v\n", ts)
		case 1:
			fmt.Println(t.n, " ") // 9  访问原数组
			fmt.Printf("%#v\n", ts)
		}
	}

	fmt.Println(ts) // [{0} {9}]
}

/*
	知识点：for-range 数组指针。
	for-range 循环中的循环变量 t 是原数组元素的副本。如果数组元素是 结构体值，则副本的字段和原数组
	字段是两个不同的值。
*/
