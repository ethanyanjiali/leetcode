package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, str string) bool {
	words := strings.Split(str, " ")
	if len(words) != len(pattern) {
		return false
	}
	used := map[rune]bool{}
	sb := strings.Builder{}
	wordToChar := map[string]rune{}
	for i, char := range pattern {
		if _, ok := wordToChar[words[i]]; !ok {
			if _, ok := used[char]; ok {
				return false
			}
			wordToChar[words[i]] = char
			used[char] = true
		}
		sb.WriteRune(wordToChar[words[i]])
	}
	return sb.String() == pattern
}

func main() {
	fmt.Println(wordPattern("abba", "dog dog dog dog"))
	fmt.Println(wordPattern("aaaa", "dog cat cat dog"))
}
