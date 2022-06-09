package main

func main() {
	x := []int{0, 1, 2}
	y := [3]*int{}
	for i, v := range x {
		defer func() {
			println(v)
		}()
		y[i] = &v
	}
	println(*y[0], *y[1], *y[2])
}

// for-range 虽然使⽤的是 :=，但是 v 不会重新声明，可以打印 v 的地址验证下。
