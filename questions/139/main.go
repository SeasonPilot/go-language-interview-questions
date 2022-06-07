package main

import "fmt"

type T struct {
	x int
	y *int
}

func main() {
	i := 20
	t := T{10, &i}

	p := &t.x

	*p++ // 递增运算符 ++ 和递减运算符 -- 的优先级 低于 解引⽤运算符 * 和取址运算符 &
	*p--

	t.y = p

	fmt.Println(*t.y) // 解引⽤运算符和取址运算符的优先级 低于 选择器 . 中的属性选择操作符。
}
