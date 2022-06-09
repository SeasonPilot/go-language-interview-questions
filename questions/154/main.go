package main

import "fmt"

func main() {
	defer func() {
		fmt.Print(recover()) // 2
	}()

	defer func() {

		defer func() {
			fmt.Print(recover()) // 1
		}()

		panic(1)

	}()

	defer recover()

	panic(2)
}

// panic不但可以在函数正常流程中抛出，在defer逻辑里也可以再次调用panic或抛出panic。defer里面的panic能够被后续执行的defer捕获。
