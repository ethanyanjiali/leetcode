package main

import "fmt"

func searchRoute(graph map[int]map[int]bool, src int, dst int) int {
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
			for next := range nexts {
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

func buildGraph(routes [][]int) map[int]map[int]bool {
	graph := map[int]map[int]bool{}
	for _, route := range routes {
		for _, start := range route {
			for j := 0; j < len(route); j++ {
				stop := route[j]
				if start == stop {
					continue
				}
				if _, ok := graph[start]; ok {
					dsts := graph[start]
					if _, ok := dsts[stop]; !ok {
						graph[start][stop] = true
					}
				} else {
					graph[start] = map[int]bool{
						stop: true,
					}
				}
			}
		}
	}
	return graph
}

func searchRoute2(routes [][]int, graph map[int]map[int]bool, src, dst int) int {
	queue := []int{}
	for routeNum := range graph[src] {
		queue = append(queue, routeNum)
	}
	visited := map[int]bool{}
	dstRoutes := map[int]bool{}
	// once one of the route in this map get hit, we find the solution
	for routeNum := range graph[dst] {
		dstRoutes[routeNum] = true
	}
	times := 1
	for len(queue) != 0 {
		newQueue := []int{}
		for _, routeNum := range queue {
			if _, ok := dstRoutes[routeNum]; ok {
				return times
			}
			for _, stop := range routes[routeNum] {
				nextRouteNums := graph[stop]
				for nextRouteNum := range nextRouteNums {
					if _, ok := visited[nextRouteNum]; !ok {
						newQueue = append(newQueue, nextRouteNum)
						visited[nextRouteNum] = true
					}
				}
			}
		}
		queue = newQueue
		times++
	}
	return -1
}

func buildGraph2(routes [][]int) map[int]map[int]bool {
	graph := map[int]map[int]bool{}
	for i, route := range routes {
		for _, stop := range route {
			if _, ok := graph[stop]; ok {
				graph[stop][i] = true
			} else {
				graph[stop] = map[int]bool{
					i: true,
				}
			}
		}
	}
	return graph
}

func numBusesToDestination(routes [][]int, S int, T int) int {
	if S == T {
		return 0
	}
	graph := buildGraph2(routes)
	return searchRoute2(routes, graph, S, T)
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
	routes = [][]int{
		{1, 7}, {3, 5},
	}
	fmt.Println(numBusesToDestination(routes, 5, 5))
	routes = [][]int{
		{10, 13, 22, 28, 32, 35, 43},
		{2, 11, 15, 25, 27},
		{6, 13, 18, 25, 42},
		{5, 6, 20, 27, 37, 47},
		{7, 11, 19, 23, 35},
		{7, 11, 17, 25, 31, 43, 46, 48},
		{1, 4, 10, 16, 25, 26, 46},
		{7, 11},
		{3, 9, 19, 20, 21, 24, 32, 45, 46, 49},
		{11, 41},
	}
	fmt.Println(numBusesToDestination(routes, 37, 43))
}
