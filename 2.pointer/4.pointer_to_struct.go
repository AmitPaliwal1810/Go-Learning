package main

import (
	"fmt"
)

type vertex struct {
	x int
	y int
}

func main() {
	v := vertex{12, 25}
	p := &v
	p.x = 10
	p.y = 2563

	fmt.Printf("%v", v)
}
