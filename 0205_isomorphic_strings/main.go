package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {
	usedT := map[byte]bool{}
	s2t := map[rune]byte{}
	if len(s) != len(t) {
		return false
	}
	for i, sChar := range s {
		tChar := t[i]
		if expect, ok := s2t[sChar]; ok {
			if expect != tChar {
				return false
			}
		} else {
			if _, ok := usedT[tChar]; ok {
				return false
			}
			s2t[sChar] = tChar
			usedT[tChar] = true
		}
	}
	return true
}

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("dog", "add"))
}
