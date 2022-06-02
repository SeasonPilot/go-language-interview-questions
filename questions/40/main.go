package main

import "fmt"

// 40. 切⽚ a、b、c 的⻓度和容量分别是多少？
func main() {
	s := [3]int{1, 2, 3}
	a := s[:0]
	b := s[:2]
	c := s[1:2:cap(s)]

	fmt.Printf("%v,%v\n", len(a), cap(a))  // 0,3
	fmt.Printf("%v,%v\n", len(b), cap(b))  // 2,3
	fmt.Printf("%v,%v\n", len(c), cap(c))  // 1,2
}

/*
	解析：
	知识点：数组或切⽚的截取操作。

	截取操作有带 2 个或者 3 个参数，形如：[i:j] 和 [i:j:k]，假设截取对象的底层数组⻓度为 l。
	在操作符 [i:j]中，如果 i 省略，默认 0，如果 j 省略，默认底层数组的⻓度，截取得到的切⽚⻓度和容量计算⽅法是 j-i、l-i。
	操作符 [i:j:k]，k 主要是⽤来限制切⽚的容量，但是不能⼤于数组的⻓度 l，截取得到的切⽚⻓度和容量计算⽅法是 j-i、k-i。
*/
