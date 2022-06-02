package main

import "fmt"

type info struct {
	result int
}

func work() (int, error) {
	return 13, nil
}

func main() {
	var data info
	data.result, err := work() // 解析： 不能使⽤短变量声明设置结构体字段值
	fmt.Printf("info: %+v\n", data)
}

// 修复代码：
func main1() {
	var data info

	var err error
	data.result, err = work() //ok
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data)
}
