package main

var f = func(i int) {
	print("x")
}

func main() {
	f := func(i int) {
		print(i)
		if i > 0 {
			f(i - 1)
		}
	}
	f(10)
}

/*
	这道题⼀眼看上去会输出 109876543210，其实这是错误的答案，这⾥不是递归。假设 main() 函数⾥为
	f2()，外⾯的为 f1()，当声明 f2() 时，调⽤的是已经完成声明的 f1()。
*/

// 看下⾯这段代码你应该会更容易理解⼀点：
var x = 23

func main() {
	x := 2*x - 4
	println(x) // 输出:42
}
