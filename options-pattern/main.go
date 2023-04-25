package main

import "fmt"

type User struct {
	Id   int32
	Name string

	Addr string
}

type Options func(user *User)

func NewUser(id int32, name string, opts ...Options) *User {
	user := &User{
		Id:   id,
		Name: name,
	}

	for _, o := range opts {
		o(user)
	}
	return user
}

func WithAddr(addr string) Options {
	return func(user *User) {
		user.Addr = addr
	}
}

func main() {
	opt1 := func(user *User) {
		user.Addr = "jilishi"
	}

	u := NewUser(3, "season", opt1)

	opt2 := WithAddr("jianshui")
	u2 := NewUser(4, "pilot", opt2)

	fmt.Println(u)
	fmt.Println(u2)
}
