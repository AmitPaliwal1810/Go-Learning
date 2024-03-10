package main

import (
	"fmt"
)

// in the below function sum is the integer that will be send as argument and x and y both are the value that will be return as the return value

func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	return
}

func main() {
	fmt.Println("Hello this our function")
	fmt.Println(Split(17))
}
