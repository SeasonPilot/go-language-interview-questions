package main

import "fmt"

type T struct {
	n int
}

func main() {
	ts := [2]T{}

	for i := range ts[:] {
		switch i {
		case 0:
			ts[1].n = 9
		case 1:
			fmt.Println(ts[i].n, " ") // 9
		}
	}

	fmt.Println(ts)
}

/*
	知识点：for-range 切⽚。
	for-range 切⽚时使⽤的是切⽚的副本，但不会复制底层数组，换句话说，此副本切⽚与原数组 共享底层数组。
*/
