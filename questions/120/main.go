package main

import "fmt"

//func main() {
//	fmt.Println(~2)
//}
/*
	解析：
	很多语⾔都是采⽤ ~ 作为按位取反运算符，Go ⾥⾯采⽤的是 ^ 。
*/

//func main() {
//	var a int8 = 3
//	var b uint8 = 3
//	var c int8 = -3
//	fmt.Printf("^%b=%b %d\n", a, ^a, ^a) // ^11=-100 -4
//	fmt.Printf("^%b=%b %d\n", b, ^b, ^b) // ^11=11111100 252
//	fmt.Printf("^%b=%b %d\n", c, ^c, ^c) // ^-11=10 2
//}

/*
	按位取反之后返回⼀个每个 bit 位都取反的数，
	对于有符号的整数来说，是按照补码进⾏取反操作的（快速计算⽅法：对数 a 取反，结果为 -(a+1) ），对于⽆符号整数来说就是按位取反。
*/
func main() {
	var x uint8 = 214
	var y uint8 = 92
	fmt.Printf("x: %08b\n", x)
	fmt.Printf("y: %08b\n", y)
	fmt.Printf("x | y: %08b\n", x|y)
	fmt.Printf("x &^ y: %08b\n", x&^y)
}

/*
  5. 给⼤家重点介绍下这个操作符 &^，按位置零，例如：z = x &^ y，表示如果 y 中的 bit 位为 1，则 z 对应bit 位为 0，否则 z 对应 bit 位等于 x 中相应的 bit 位的值。
  6. 不知道⼤家发现没有，我们还可以这样理解或操作符 | ，表达式 z = x | y，如果 y 中的 bit 位为 1，则 z 对应 bit 位为 1，否则 z 对应 bit 位等于 x 中相应的 bit 位的值，与 &^ 完全相反。
*/
