package main

func main() {
	for i := 0; i < 10; i++ {
	loop:
		println(i)
	}
	goto loop
}

// goto 不能跳转到 其他函数 或者 内层代码。
