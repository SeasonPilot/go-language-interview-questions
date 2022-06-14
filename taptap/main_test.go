package main

import (
	"fmt"
	"testing"
)

func TestGetSteps(t *testing.T) {
	nums := []int{-1, 0, 3, 17, 19, 23, 45, 67, 7879, 676}
	for _, i := range nums {
		fmt.Printf("%#v\n", GetSteps(i))
	}
}
