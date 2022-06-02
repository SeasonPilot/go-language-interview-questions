package main

import "fmt"

type Math struct {
	x, y int
}

var m = map[string]*Math{ // value 类型为指针类型
	"foo": &Math{2, 3},
}

func main() {
	m["foo"].x = 4
	fmt.Println(m["foo"].x)
	fmt.Printf("%#v", m["foo"]) // %#v 格式化输出详细信息
}
