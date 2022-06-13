package main

import "fmt"

/*
	https://fcxnw9u8pt.feishu.cn/docs/doccnsTD5a8ugC5uRRsUSmcYj6u#ewbih8
	6. 长度为 0 的数组，进行[1:]切片后会报数组越界吗？
	invalid slice index 1 (out of bounds for 0-element array)
*/
func main() {
	a := [0]int{}
	fmt.Printf("%#v\n", a[1:])
}
