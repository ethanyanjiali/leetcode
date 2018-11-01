package main

import (
	"fmt"
)

func isAdjacent(dst string, src string) bool {
	if len(dst) != len(src) {
		return false
	}
	diff := 0
	for i := range src {
		if dst[i] != src[i] {
			diff++
		}
		if diff > 2 {
			break
		}
	}
	return diff == 1
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	visited := map[string]bool{}
	queue := []string{beginWord}
	steps := 1
	for len(queue) != 0 {
		newQueue := []string{}
		for _, curr := range queue {
			if curr == endWord {
				return steps
			}
			if _, ok := visited[curr]; ok {
				continue
			}
			visited[curr] = true
			for _, word := range wordList {
				if isAdjacent(word, curr) {
					newQueue = append(newQueue, word)
				}
			}
		}
		queue = newQueue
		steps++
	}
	return 0
}

func main() {
	wordList := []string{
		"dog",
		"dot",
		"bot",
		"bat",
	}
	fmt.Println(ladderLength("dog", "bat", wordList))
	wordList = []string{
		"hot", "dot", "dog", "lot", "log", "cog",
	}
	fmt.Println(ladderLength("hit", "cog", wordList))
}
