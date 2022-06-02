package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	defer fmt.Println("defer main")
	var user = os.Getenv("USER_")

	go func() {
		defer func() {
			fmt.Println("defer caller")
			if err := recover(); err != nil {
				fmt.Println("recover success. err: ", err)
			}
		}()

		func() {
			defer func() {
				fmt.Println("defer here")
			}()

			if user == "" {
				panic("should set user env.")
			}

			// 此处不会执行
			fmt.Println("after panic")
		}()
	}()

	time.Sleep(100)
	fmt.Println("end of main function")
}

/*
	panic会停掉当前正在执行的程序，不只是当前协程。在这之前，它会有序地执行完当前 协程 defer列表里的语句，其它协程里挂的defer语句不作保证。
	因此，我们经常在defer里挂一个recover语句，防止程序直接挂掉，这起到了try...catch的效果。
*/
