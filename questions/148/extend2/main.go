package main

func main() {
	const a int8 = -1
	var b int8 = -128 / a //

	println(b)
}

/*
	对于 var b int8 = -128 / a，因为 a 是 int8 类型常量，所以 -128 / a 是常量表达式，在编译器计算，结果必然也是常量。
	因为 a 的类型是 int8，因此 -128 也会隐式转为 int8 类型，128 这个结果超过了 int8 的范围，但常量不允许溢出，因此编译报错。
*/
