package main

import (
	"fmt"
)

/*
for continued

the init and post statement are optional

*/

func main() {
	sum := 1

	for sum < 100 {
		sum += sum
	}

	fmt.Printf("the sum is %v", sum)
}

/*
the answer is 128
1. sum = 1
2. sum = 1 + 1 = 2
3. sum = 2 + 2 = 4
4. sum = 4 + 4 = 8
5. sum = 8 + 8 = 16
6. sum = 16 + 16 = 32
7. sum = 32 + 32 = 64
8. sum = 64 + 64 = 128

*/
