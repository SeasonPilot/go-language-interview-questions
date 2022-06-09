package main

//func main() {
//	recover()
//	panic(1)
//}
//
//func main() {
//	defer recover()
//	panic(1)
//}

func main() { // 正确
	defer func() {
		recover()
	}()
	panic(1)
}

//func main() {
//	defer func() {
//		defer func() {
//			recover()
//		}()
//	}()
//	panic(1)
//}

/*
	recover() 必须在 defer() 函数中直接调⽤才有效。
	上⾯其他⼏种情况调⽤都是⽆效的：直接调⽤recover()、在 defer() 中直接调⽤ recover() 和 defer() 调⽤时多层嵌套。
*/
