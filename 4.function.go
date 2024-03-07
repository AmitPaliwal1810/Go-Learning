package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

// if the argument type is same you can write a single type also.
func multiply(a, b int) int {
	return a * b
}

func main() {
	fmt.Println("The sum is ", add(5, 7))
	fmt.Println("The multiply is", multiply(48, 2))
}
