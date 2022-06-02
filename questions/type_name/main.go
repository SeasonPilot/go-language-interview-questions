package main

import "fmt"

type User struct {
}

func (user *User) Pfunc() {
	fmt.Println("pfunc")
}

func (user User) Vfunc() {
	fmt.Println("vfunc")
}

func main() {
	fmt.Printf("%T\n", &User)
}
