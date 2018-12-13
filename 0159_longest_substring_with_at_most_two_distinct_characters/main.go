package main

import (
	"fmt"
)

func lengthOfLongestSubstringTwoDistinct(s string) int {
    chars := map[byte]int{}
    start := 0
    maxLength := 0
    for index := 0; index < len(s); index++ {
    	char := s[index]
    	if _, ok := chars[char]; !ok && len(chars) >= 2 {
    		minIndex := len(s)
    		for _, value := range chars {
    			if value < minIndex {
    				minIndex = value
    			}
    		}
    		start = minIndex + 1
    		delete(chars, s[minIndex])
    	}
    	length := index - start + 1
		if maxLength < length {
			maxLength = length
		}
    	chars[char] = index
    }
    return maxLength
}

func main() {
	fmt.Println(lengthOfLongestSubstringTwoDistinct("eceba"))
	fmt.Println(lengthOfLongestSubstringTwoDistinct("ccaabbb"))
}