package main

import (
	"fmt"
)

/*
Slice literals
A slice literal is like an array literal without the length.
*/

func main() {
	a := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("%v\n", a)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	// array of struct

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
