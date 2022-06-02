package main

import "fmt"

type Direction int

const (
	North Direction = iota // 0
	East                   // 1
	South                  // 2
	West                   // 3
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South 1111", "West"}[d]
}
func main() {
	fmt.Println(South) // 2
}
