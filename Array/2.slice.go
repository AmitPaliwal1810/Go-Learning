package main

import (
	"fmt"
)

/*

An array has fixed size. A slice, on the other hand, is a dynamically sized, flexible view into the elements of an array.

syntax: a[low : high]

*/

func main() {
	a := [6]int{1,2,3,4,5,6}
	fmt.Printf("%v", a[1:5])
}
