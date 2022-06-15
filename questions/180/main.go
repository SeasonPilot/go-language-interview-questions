package main

import "fmt"

func alwaysFalse() bool {
	return false
}

func main() {
	//fmt.Printf("%#v\n", alwaysFalse())

	switch alwaysFalse(); { // 有 ;
	case true:
		println(true)
	case false:
		println(false)
	}

	switch a := alwaysFalse(); a {
	case true:
		println(true)
	case false:
		println(false)
	}

	switch alwaysFalse() { // 没有 ;  多数情况使用的是这个
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	}

	if x := alwaysFalse(); !x {
		// ...
	}

}

// https://gfw.go101.org/article/line-break-rules.html#:~:text=%E6%B3%A8%E6%84%8F%EF%BC%8C%E4%B8%8A%E4%BE%8B%E4%B8%AD%E7%9A%84switch%2Dcase%E4%BB%A3%E7%A0%81%E5%9D%97%E5%B0%86%E8%BE%93%E5%87%BAtrue%EF%BC%8C%E8%80%8C%E4%B8%8D%E6%98%AFfalse%E3%80%82%20%E6%AD%A4%E4%BB%A3%E7%A0%81%E5%9D%97%E5%92%8C%E4%B8%8B%E9%9D%A2%E8%BF%99%E4%B8%AA%E6%98%AF%E4%B8%8D%E5%90%8C%E7%9A%84%EF%BC%9A
