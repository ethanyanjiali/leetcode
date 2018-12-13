package main

import (
	"fmt"
)

func computeFinal(s string) []byte {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
    	if s[i] == byte('#') {
    		if len(stack) > 0 {
    			stack = stack[:len(stack) - 1]
    		}
    	} else {
    		stack = append(stack, s[i])
    	}
    }
    return stack
}

func backspaceCompare(S string, T string) bool {
    stackS := computeFinal(S)
    stackT := computeFinal(T)
    if len(stackS) == len(stackT) {
    	return string(stackS) == string(stackT)
    }
    return false
}

func main() {
	fmt.Println(backspaceCompare("y#fo##f", "y#f#o##f"))
	fmt.Println(backspaceCompare("a##c", "#a#c"))
}