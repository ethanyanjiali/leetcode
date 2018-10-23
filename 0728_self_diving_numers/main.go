package main

import (
	"fmt"
)

func selfDividingNumbers(left int, right int) []int {
	var res []int
	for ; left <= right; left++ {
		curr := left
		for curr > 0 {
			rem := curr % 10
			// the digit can't be zero and should be divisible
			if rem == 0 || left%rem != 0 {
				break
			}
			curr = curr / 10
		}
		if curr == 0 {
			res = append(res, left)
		}
	}
	return res
}

func main() {
	fmt.Println(selfDividingNumbers(1, 22))
}
