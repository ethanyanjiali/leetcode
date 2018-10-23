package main

import (
	"fmt"
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)
	indegree := make([]int, numCourses)
	// build the graph, count the indegree for each course
	for _, prerequisite := range prerequisites {
		course1 := prerequisite[0]
		course2 := prerequisite[1]
		if graph[course2] != nil {
			graph[course2] = append(graph[course2], course1)
		} else {
			graph[course2] = []int{course1}
		}
		indegree[course1]++
	}
	// find out those with 0 indegree as roots of the graph
	var queue []int
	for i, degree := range indegree {
		if degree == 0 {
			queue = append(queue, i)
		}
	}
	// traverse the graph from each root, decrement the indegree for each leaf
	count := 0
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		count++
		for _, course := range graph[curr] {
			indegree[course]--
			// once a node's indegree is 0, it becomes a root too
			if indegree[course] == 0 {
				queue = append(queue, course)
			}
		}
	}
	// if all nodes have been the root before, we traversed the entire graph
	return count == numCourses
}

func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))
	fmt.Println(canFinish(3, [][]int{{1, 0}, {1, 2}, {0, 1}}))
}
