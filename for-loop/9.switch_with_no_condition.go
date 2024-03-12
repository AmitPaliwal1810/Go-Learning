package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	// there is no condition.
	switch {
	case t.Hour() < 12:
		fmt.Printf("Good Morning")
	case t.Hour() < 17:
		fmt.Printf("Good Afternoon")
	default:
		fmt.Printf("Good Evening")
	}
}
