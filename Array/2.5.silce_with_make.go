package main

import (
	"fmt"
)

func main() {
	s := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%v\n", s)

	// and this array will fill with 0
	b := make([]int, 5, 5) // [] int represent the array , 5 represent the length and another 5 represent capacity.
	fmt.Println(b)
}
