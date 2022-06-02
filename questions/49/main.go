package main

import "fmt"

// 49. What is the output of the following code snippet?

type Person struct {
	age int
}

func main() {
	person := &Person{28} // 给变量赋值了一个地址  0xc00006fdd8
	// 1.
	defer fmt.Println(person.age) // 28
	// 2.
	defer func(p *Person) { // 28
		fmt.Println(p.age)
	}(person)
	// 3.
	defer func() {
		fmt.Println(person.age)
	}()
	person = &Person{29} // 给变量赋值了一个新地址。变成了一个新地址了。   0xc00006fde0
}

/*
	解析：
	知识点： defer 函数的执⾏顺序。

	这道题在第 47 题⽬的基础上做了⼀点点⼩改动，前⼀题最后⼀⾏代码 person.age = 29 是修改引⽤对象的成员 age，
	这题最后⼀⾏代码 person = &Person{29} 是修改引⽤对象本身，来看看有什么区别。

	1处.person.age 这⼀⾏代码跟之前含义是⼀样的，此时是将 28 当做 defer 函数的参数，会把 28 缓存在栈中，等到最后执⾏该 defer 语句的时候取出，即输出 28；
	2处.defer 缓存的是结构体 Person{28} 的地址，这个地址指向的结构体没有被改变，最后 defer 语句后⾯的函数执⾏的时候取出仍是 28；
	3处.闭包引⽤，person 的值已经被改变，指向结构体 Person{29} ，所以输出 29.
	由于 defer 的执⾏顺序为先进后出，即 3 2 1，所以输出 29 28 28。
*/
