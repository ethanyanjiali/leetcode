package main

import (
	"fmt"
)

func isValid(s string) bool {
	var stack []rune
	var tuple = map[rune]rune{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	for _, rune := range s {
		size := len(stack)
		if size > 0 {
			top := stack[size-1]
			if tuple[top] == rune {
				stack = stack[:size-1]
				continue
			}
		}
		stack = append(stack, rune)
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("{(}"))
}
