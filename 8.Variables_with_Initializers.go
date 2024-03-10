package main

import (
	"fmt"
)

// A var declaration can include initializers, one per variable.

// If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

var i, j = 1, 2

func main() {
	var c, paython, java = true, false, false
	fmt.Println("hello")
	fmt.Println(c, paython, java, i, j)
}
