package main

import "fmt"

// 截取切片
func main() {
	s := make([]int, 3, 9)
	fmt.Println(len(s))

	//s2 := s[4:] // slice bounds out of range [4:3]

	//s2 := s[4:8:10] // slice bounds out of range [::10] with capacity 9
	s2 := s[4:8]
	fmt.Println(len(s2)) // 4
	fmt.Println(cap(s2)) // 5
}

/*
	从⼀个基础切⽚派⽣出的⼦切⽚的⻓度可能⼤于基础切⽚的⻓度。假设基础切⽚是 baseSlice，使⽤操作
	符 [low,high]，有如下规则：0 <= low <= high <= cap(baseSlice)，只要上述满⾜这个关系，下标 low
	和 high 都可以⼤于 len(baseSlice)。
*/
