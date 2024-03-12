package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now().Weekday()

	fmt.Printf("%v \n", today)
	fmt.Printf("%v \n", today+3) // you can add any number it will provide you the next and other day according to your number and that is depend upon your day ex- today is wednesday so next day +1 will be thrurday.

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.\n")
	case today + 1:
		fmt.Println("Tomorrow.\n")
	case today + 2:
		fmt.Println("In two days.\n")
	default:
		fmt.Println("Too far away.\n")
	}
}
