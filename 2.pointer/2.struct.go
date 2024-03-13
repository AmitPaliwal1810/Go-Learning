package main

import (
	"fmt"
)

/*
similar c language , go also has struct which is the collection of the multiple values.
*/

type Coordinate struct {
	x int
	y int
}

func main() {
	fmt.Printf("%v \n", Coordinate{1, 12})
	fmt.Println(Coordinate{15, 12})
}
