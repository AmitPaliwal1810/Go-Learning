package main

import "fmt"

// if while slicing there is no value present bydefault it is 0

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4] // [3,5,7]
	fmt.Println(s)

	s = s[:2] // [3,5]
	fmt.Println(s)

	s = s[1:] // [5]
	fmt.Println(s)

	s = s[:] // [5] because slice will be [0 : 0]
	fmt.Println(s)
}
