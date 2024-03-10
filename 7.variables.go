package main

import (
	"fmt"
)

// The var statement declares a list of variables; as in function argument lists, the type is last.
// A var statement can be at package or function level. We see both in this example.

var c, pyhton, java bool

func main() {
	fmt.Println("hello")
	fmt.Println(c, java, pyhton)
}
