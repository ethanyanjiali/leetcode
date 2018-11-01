package main

import (
	"fmt"
	"strings"
)

func buildGraphAndIndegree(words []string) (map[byte][]byte, map[byte]int) {
	graph := make(map[byte][]byte)
	indegree := make(map[byte]int)
	for i := 1; i < len(words); i++ {
		prev := words[i-1]
		curr := words[i]
		for idx := range prev {
			if idx >= len(prev) {
				break
			}
			prevChar := prev[idx]
			if _, ok := indegree[prevChar]; !ok {
				indegree[prevChar] = 0
			}
			if idx >= len(curr) {
				break
			}
			currChar := curr[idx]
			if prevChar == currChar {
				continue
			}
			targets := graph[prevChar]
			found := false
			for _, el := range targets {
				if el == currChar {
					found = true
				}
			}
			if !found {
				graph[prevChar] = append(targets, currChar)
				indegree[currChar]++
			}
		}
	}
	return graph, indegree
}

func alienOrder(words []string) string {
	graph, indegree := buildGraphAndIndegree(words)
	// find the first batch of roots for topological sort
	var roots []byte
	sb := strings.Builder{}
	for key, value := range indegree {
		if value == 0 {
			roots = append(roots, key)
			sb.WriteByte(key)
		}
	}
	// keep sorting
	for len(roots) != 0 {
		newRoots := []byte{}
		for _, root := range roots {
			targets := graph[root]
			for _, target := range targets {
				if indegree[target] > 0 {
					indegree[target]--
				}
				if indegree[target] == 0 {
					newRoots = append(newRoots, target)
				}
			}
		}
		for _, root := range newRoots {
			sb.WriteByte(root)
		}
		roots = newRoots
	}
	if sb.Len() == len(indegree) {
		return sb.String()
	}
	return ""
}

func main() {
	words := []string{
		"wrt",
		"wrf",
		"er",
		"ett",
		"rftt",
	}
	fmt.Println(alienOrder(words))
	words = []string{
		"wrt",
		"wrft",
		"wrff",
		"er",
		"ett",
		"rftt",
	}
	fmt.Println(alienOrder(words))
	words = []string{
		"wrt",
		"wrtf",
		"wrft",
	}
	fmt.Println(alienOrder(words))
}
