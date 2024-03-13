package main

import (
	"fmt"
)

type vertex struct {
	x int
	y int
}

func main() {

	v := vertex{14, 25}

	fmt.Printf("%v", v.x)
}
