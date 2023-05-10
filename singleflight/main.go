package main

import (
	"fmt"
	"time"

	"golang.org/x/sync/singleflight"
)

var sf singleflight.Group

func main() {
	// 定义一个获取数据的函数
	getData := func() (interface{}, error) {
		fmt.Println("Fetching data...")
		time.Sleep(time.Second * 3)
		return "data", nil
	}

	// 同时发起三个请求
	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Println(i)
			time.Sleep(time.Second * 3)

			// 使用 singleflight.Do 方法获取数据
			data, _, _ := sf.Do("key", getData)
			fmt.Println("Got data:", data)
		}(i)
	}

	// 等待所有请求完成
	time.Sleep(7 * time.Second)
}
