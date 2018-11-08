package main

import "fmt"

func searchPattern(pattern string, str string, dict map[byte]string, used map[byte]bool, pStart, sStart int) bool {
	if pStart == len(pattern) - 1 && sStart == len(str) - 1 {
		return true
	}
	if pStart >= len(pattern) || sStart >= len(str) {
		return false
	}
	char := pattern[pStart]
	for i := sStart + 1; i <= len(str); i++ {
		target := str[sStart:i]
		shouldUpdate := false
		if _, ok := dict[char]; ok {
			if dict[char] != target {
				return false
			}
		} else {
			if _, ok := used[char]; ok {
				return false
			}
			shouldUpdate = true
		}
		if shouldUpdate {
			used[char] = true
			dict[char] = target
		}
		res := searchPattern(pattern, str, dict, used, pStart+1, i)
		if res {
			return true
		}
		if shouldUpdate {
			delete(used, char)
			delete(dict, char)
		}
	}
	return false
}

func wordPatternMatch(pattern string, str string) bool {
	dict := map[byte]string{}
	used := map[byte]bool{}
	return searchPattern(pattern, str, dict, used, 0, 0)
}

func main() {
	//fmt.Println(wordPatternMatch("aba", "redbluered"))
	fmt.Println(wordPatternMatch("abab", "redblueredblue"))
	//fmt.Println(wordPatternMatch("aaaa", "asdasdasdasd"))
	//fmt.Println(wordPatternMatch("aabb", "xyzabcxzyabc"))
}