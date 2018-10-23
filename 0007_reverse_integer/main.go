package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	// The reserse of MinInt32 will overflow
	if x == math.MinInt32 {
		return 0
	}
	sign := 1
	// Handle negative numbers
	if x < 0 {
		x = -x
		sign = -1
	}
	result := 0
	for x > 0 {
		rem := x % 10
		// When the result == MaxInt32/10, the reminder has to be smaller than 7
		// so that result*10 + rem won't overflow (MaxInt32 is 2147483647)
		if math.MaxInt32/10 < result || (math.MaxInt32/10 == result && rem > 7) {
			return 0
		}
		result = result*10 + rem
		x = x / 10
	}
	return result * sign
}

func main() {
	// fmt.Println(reverse(math.MinInt32))
	// fmt.Println(reverse(math.MaxInt32))
	fmt.Println(reverse(1463847412))
	fmt.Println(math.MaxInt32)
}
