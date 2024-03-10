package main

import (
	"fmt"
	"math/cmplx"
)

/*
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
*/

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type of ToBe is %T and %v\n", ToBe, ToBe)
	fmt.Printf("Type of MaxInt is %T and %v\n", MaxInt, MaxInt)
	fmt.Printf("Type of z is %T and %v\n", z, z)
}

/*
 fmt.Printf provides more control over the formatting of variables using format specifiers, 
 while println simply prints the raw values of the arguments.
*/
