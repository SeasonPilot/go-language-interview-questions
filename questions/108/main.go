package main

import "fmt"

func main() {
	var mm = make(map[int]string) //这个是用make申明，make 的作用是初始化内置的数据结构，可以对切片、哈希表和 Channel进行初始化；
	fmt.Println(mm == nil)        //输出map[]，这里的mm2 不是nil值

	var mm2 map[int]string  //这个是像普通类型一样申明,这里的mm2 是nil值
	fmt.Println(mm2 == nil) //输出map[]

	mm = map[int]string{9: "nine", 7: "seven"}
	mm2 = map[int]string{8: "eight", 4: "four"}
	fmt.Println(mm)
	fmt.Println(mm2)
}
