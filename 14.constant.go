package main

import (
	"fmt"
)

/*
Constant

Constants are declared like variables, but with the const keyword.
Constants can be character, string, boolean, or numeric values.
Constants cannot be declared using the := syntax.

*/

const pi = 22 / 7

func main() {
	fmt.Printf("pi value is %v\n", pi)
	const Truth = true
	fmt.Println("Go rules?", Truth)
}
