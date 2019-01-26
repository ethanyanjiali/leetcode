package main

import (
	"fmt"
	"math"
)

func buildGraph(flights [][]int) map[int]map[int]int {
	graph := map[int]map[int]int{}
	for _, flight := range flights {
		src, dst, cost := flight[0], flight[1], flight[2]
		cmap, ok := graph[src]
		if !ok {
			cmap = map[int]int{}
		}
		cmap[dst] = cost
		graph[src] = cmap
	}
	return graph
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	graph := buildGraph(flights)
	queue := [][]int{{src, 0}}
	stops := -1
	cheapest := math.MaxInt64
	visited := map[int]int{}
	for len(queue) > 0 && stops <= K {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[i]
			from, sum := curr[0], curr[1]
			if from == dst && sum < cheapest {
				cheapest = sum
				continue
			}
			toMap := graph[from]
			for to, cost := range toMap {
				if min, ok := visited[to]; ok && sum+cost > min {
					continue
				}
				queue = append(queue, []int{to, sum + cost})
				visited[to] = sum + cost
			}
		}
		queue = queue[size:len(queue)]
		stops++
	}
	if cheapest == math.MaxInt64 {
		return -1
	}
	return cheapest
}

func main() {
	flights := [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}
	src := 0
	dst := 2
	K := 0
	fmt.Println(findCheapestPrice(3, flights, src, dst, K))
}
