package main

import (
	"fmt"
)

func main() {

	sum := 0

	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Printf("sum is %v", sum)
}
