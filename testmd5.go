package main

import (
	"fmt"
	_ "opms/initial"
	_ "opms/routers"
	_ "opms/utils"
)

func main() {
	f := fibonacci(0, 1)
	for i := 0; i < 300; i++ {
		fmt.Println("next value: ", f())
	}
}
func add(x, y int) int {
	return x + y
}

func fibonacci(s1, s2 int) func() int {
	fmt.Println("fibonacci start with: ", s1, " and ", s2)
	left := s1
	right := s2

	return func() int {
		next := add(left, right)
		left = right
		right = next
		return next
	}
}
