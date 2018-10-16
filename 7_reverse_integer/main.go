package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	if x == math.MinInt32 {
		return 0
	}
	sign := 1
	if x < 0 {
		x = -x
		sign = -1
	}
	result := 0
	for x > 0 {
		rem := x % 10
		if math.MaxInt32/10 <= result {
			return 0
		}
		result = result*10 + rem
		x = x / 10
	}
	return result * sign
}

func main() {
	fmt.Println(reverse(math.MinInt32))
	fmt.Println(reverse(math.MaxInt32))
	fmt.Println(reverse(746384741))
}
