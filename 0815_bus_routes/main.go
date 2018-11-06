package main

import (
	"fmt"
)

func searchRoute(graph map[int][]int, src int, dst int) int {
	queue := []int{src}
	visited := map[int]map[int]bool{}
	times := 0
	for len(queue) != 0 {
		newQueue := []int{}
		currVisited := map[int]bool{}
		for _, start := range queue {
			if start == dst {
				return times
			}
			nexts := graph[start]
			for _, next := range nexts {
				if _, ok := visited[next]; !ok {
					visited[next] = map[int]bool{}
				}
				if _, ok := visited[next][start]; !ok {
					if _, ok = currVisited[next]; !ok {
						newQueue = append(newQueue, next)
						visited[next][start] = true
						currVisited[next] = true
					}
				}
			}
		}
		queue = newQueue
		times++
	}
	return -1
}

func buildGraph(routes [][]int) map[int][]int {
	graph := map[int][]int{}
	for _, route := range routes {
		for _, start := range route {
			for j := 0; j < len(route); j++ {
				stop := route[j]
				if start == stop {
					continue
				}
				if _, ok := graph[start]; ok {
					dsts := graph[start]
					found := false
					for _, dst := range dsts {
						if dst == stop {
							found = true
							break
						}
					}
					if found != true {
						graph[start] = append(graph[start], stop)
					}
				} else {
					graph[start] = []int{stop}
				}
			}
		}
	}
	return graph
}

func numBusesToDestination(routes [][]int, S int, T int) int {
	graph := buildGraph(routes)
	return searchRoute(graph, S, T)
}

func main() {
	routes := [][]int{
		{1, 2, 7},
		{3, 6, 7},
	}
	fmt.Println(numBusesToDestination(routes, 1, 6))
	routes = [][]int{
		{7, 12},
		{4, 5, 15},
		{6},
		{15, 19},
		{9, 12, 13},
	}
	fmt.Println(numBusesToDestination(routes, 15, 12))
	routes = [][]int{
		{24}, {3, 6, 11, 14, 22}, {1, 23, 24}, {0, 6, 14}, {1, 3, 8, 11, 20},
	}
	fmt.Println(numBusesToDestination(routes, 20, 8))
}
