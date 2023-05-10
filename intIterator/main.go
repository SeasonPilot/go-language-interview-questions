package main

import "fmt"

type IntIterator struct {
	start   int
	end     int
	current int
}

func (it *IntIterator) Next() int {
	if it.HasNext() {
		result := it.current
		it.current++
		return result
	}
	return 0
}

func (it *IntIterator) HasNext() bool {
	return it.current < it.end
}

func main() {
	it := &IntIterator{1, 10, 1}
	for it.HasNext() {
		fmt.Println(it.Next())
	}
}
