package main

import (
	"fmt"
)

func main() {

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var a [2]string
	a[0] = "hello"
	a[1] = "this is me"

	fmt.Println("Array.", a)
}
