package main

import "fmt"

// 58. What is the output of the following code snippet?

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b)) // 1 2
	a = 0
	defer calc("2", a, calc("20", a, b)) // 0 2
	b = 1
}
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

/*
	解析：
	程序执⾏到 main() 函数三⾏代码的时候，会先执⾏ calc() 函数的 b 参数，即： calc("10",a,b) ，输
	出：10 1 2 3，得到值 3，
	因为 defer 定义的函数是延迟函数，故 calc("1",1,3) 会被延迟执⾏；
	程序执⾏到第五⾏的时候，同样先执⾏ calc("20",a,b) 输出：20 0 2 2 得到值 2，同样将 calc("2",0,2) 延
	迟执⾏；
	程序执⾏到末尾的时候，按照栈先进后出的⽅式依次执⾏：calc("2",0,2)，calc("1",1,3)，则就依次输
	出：2 0 2 2，1 1 3 4。
*/

/*
我的答案：  错误。
20 0 2 2

2 0 2 2

10 1 2 3

1 1 3 4
*/
