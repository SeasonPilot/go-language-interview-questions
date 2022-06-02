package main

import "fmt"

type f struct {

}

func main() {
	select f {
	case cap(2):
		fmt.Println()
	case:
		fmt.Println()
	}
}
