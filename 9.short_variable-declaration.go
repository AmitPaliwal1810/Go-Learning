package main

import (
	"fmt"
)

/*
Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
*/

var a int = 4

func main() {
	b := 1.25
	c, python := true, 25

	fmt.Println("hello")
	fmt.Println(a, b, c, python)
}
