package main

import (
	"math"
	"strings"
)

func repeatedStringMatch(A string, B string) int {
	times := int(math.Ceil(float64(len(B)) / float64(len(A))))
	match := strings.Repeat(A, times)
	if strings.Contains(match, B) {
		return times
	}
	match = match + A
	if strings.Contains(match, B) {
		return times + 1
	}
	return -1
}

func main() {

}
