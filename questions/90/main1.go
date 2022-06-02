package main

import "fmt"

func main() {
	var i *int
	i = new(int)
	*i = 1
	fmt.Println(i, &i)

	a := 6
	i = &a
	fmt.Println(i, &i)
}
