package main

import "fmt"

func main() {
	var m map[int]bool // nil
	//m[456] = false     // nil map 不能赋值，可以访问。 // panic: assignment to entry in nil map  // nil map 不能用来存放键值对
	fmt.Printf("%v\n", m[456]) // false
	v, ok := m[456]            // 查看元素在 map 中是否存在
	fmt.Println(v, ok)         // false false
	_ = m[123]                 // 查看 key 为 123 的值
	fmt.Println(m[123])        // false

	var p *[5]string // nil
	for range p {
		_ = len(p)
		fmt.Printf("%v\n", len(p)) // 5
	}

	var s []int // nil
	_ = s[:]
	fmt.Printf("%#v\n", s[:])                      // []int(nil)
	fmt.Printf("%#v, %#v\n", len(s[:]), cap(s[:])) // 0, 0
	s, s[0] = []int{1, 2}, 9                       //  index out of range [0] with length 0
}
