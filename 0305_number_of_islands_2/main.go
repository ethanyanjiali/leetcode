package main

import "fmt"

type Point struct {
	X, Y int
}

type UnionFind struct {
	Size []int
	Parents []int
	IsLand []bool
	M, N, Count int
}

func (uf *UnionFind) Union(idx1, idx2 int) {
	if idx1 == idx2 {
		return
	}
	parent1 := uf.Find(idx1)
	parent2 := uf.Find(idx2)
	if uf.Size[parent1] > uf.Size[parent2] {
		uf.Parents[parent1] = parent2
	} else if uf.Size[parent2] > uf.Size[parent1] {
		uf.Parents[parent2] = parent1
	} else {
		uf.Parents[parent2] = parent1
		uf.Size[parent2] += 1
	}
	uf.Count--
}

func (uf *UnionFind) Find(idx int) int {
	parent := uf.Parents[idx]
	if parent == idx {
		return idx
	}
	return uf.Find(parent)
}

func (uf *UnionFind) GetIndex(x, y int) int {
	return x * uf.N + y
}

func (uf *UnionFind) AddLand(x, y int) {
	idx := uf.GetIndex(x, y)
	uf.IsLand[idx] = true
	uf.Count++
}

func (uf *UnionFind) GetIsLand(x, y int) bool {
	idx := uf.GetIndex(x, y)
	return uf.IsLand[idx]
}

func (uf *UnionFind) Connect(x1, y1, x2, y2 int) {
	uf.Union(uf.GetIndex(x1, y1), uf.GetIndex(x2, y2))
}

func NewUnionFind(n int, m int) *UnionFind {
	parents := make([]int, n*m)
	for i := range parents {
		parents[i] = i
	}
	return &UnionFind{
		Size: make([]int, n*m),
		Parents: parents,
		IsLand: make([]bool, n*m),
		M: m,
		N: n,
		Count: 0,
	}
}

func numIslands2 (n int, m int, operators []*Point) []int {
	uf := NewUnionFind(m, n)
	res := []int{}
	for _, operator := range operators {
		x, y := operator.X, operator.Y
		uf.AddLand(x, y)
		if x > 0 && uf.GetIsLand(x-1, y) {
			uf.Connect(x, y, x - 1, y)
		}
		if x < n - 1 && uf.GetIsLand(x+1, y) {
			uf.Connect(x, y, x + 1, y)
		}
		if y > 0 && uf.GetIsLand(x, y-1) {
			uf.Connect(x, y, x, y - 1)
		}
		if y < m - 1 && uf.GetIsLand(x, y+1) {
			uf.Connect(x, y, x, y + 1)
		}
		res = append(res, uf.Count)
	}
	return res
}

func main() {
	n :=4
	m :=5
	points := []*Point{
		{1,1},{0,1},{3,3},{3,4},
	}
	//fmt.Println(numIslands2(n, m, points))
	n =4
	m =5
	points = []*Point{
		{1,1},{0,0},{1,0},{1,1},
	}
	fmt.Println(numIslands2(n, m, points))
	}
