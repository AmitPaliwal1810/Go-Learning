package main

import (
	"fmt"
)

/*
Go has pointers, which hold the address in the memory.
The type *T is a pointer to a T valu. its zero value is nil.

var p *int

i := 45
value = &i

*/

func main() {

	i := 10
	p := &i

	fmt.Printf("the value of p is %v\n", p) // this will print the address of the value i
	fmt.Printf("the value of p is %v\n", *p) // this will print the value on that address of i
}
