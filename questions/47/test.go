package main

import "fmt"

// 47. What is the output of the following code snippet?

type Person struct {
	age int
}

func main() {
	person := &Person{28}
	fmt.Printf("%T\n", person)
	fmt.Printf("%T\n", person.age)
	fmt.Println(person.age)
}
