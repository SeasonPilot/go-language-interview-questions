package main

import "fmt"

func main() {
	//以下捕获失败
	defer recover()
	defer fmt.Println(recover())
	defer func() {
		func() {
			recover() //无效，嵌套两层
		}()
	}()

	//以下捕获有效
	defer func() {
		recover()
	}()

}

//以下捕获有效
func except() {
	recover()
}

func test() {
	defer except()
	panic("runtime error")
}

// https://www.runoob.com/note/38704
