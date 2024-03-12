package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i) // while applying defer and creating a stack , means your's last print will print first and first will print in last.
	}

	fmt.Println("done")
}
