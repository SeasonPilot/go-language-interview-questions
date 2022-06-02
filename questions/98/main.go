package main

type Param map[string]interface{}

type Show struct {
	*Param
}

func main() {
	s := new(Show)
	s.Param["day"] = 2
}

/*
	解析：
	1. map 需要初始化才能使⽤；
	2. 指针不⽀持索引。修复代码如下：
*/
