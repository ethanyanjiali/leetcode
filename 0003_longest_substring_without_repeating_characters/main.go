package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
    chars := map[byte]int{}
    start := 0
    maxLength := 0
    for i := 0; i < len(s); i++ {
    	char := s[i]
    	if _, ok := chars[char]; ok {
    		// if last index >= start, it means we have a duplicate
    		if chars[char] >= start {
    			start = chars[char] + 1
    			delete(chars, char)
    		}
    	}
    	chars[char] = i
    	length := i - start + 1
    	if maxLength < length {
    		maxLength = length
    	}
    }
    return maxLength
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}