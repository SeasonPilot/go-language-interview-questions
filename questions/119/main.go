package main

type foo struct {
	bar int
}

func main() {
	var f foo
	f.bar, tmp := 1, 2
}

/*
	解析：
	:= 操作符不能⽤于结构体字段赋值。
*/
