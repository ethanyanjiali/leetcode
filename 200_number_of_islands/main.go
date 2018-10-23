package main

import (
	"fmt"
)

func explore(grid *[][]byte, row int, col int) {
	(*grid)[row][col] = '0'
	if row > 0 && (*grid)[row-1][col] == '1' {
		explore(grid, row-1, col)
	}
	if row < len((*grid))-1 && (*grid)[row+1][col] == '1' {
		explore(grid, row+1, col)
	}
	if col > 0 && (*grid)[row][col-1] == '1' {
		explore(grid, row, col-1)
	}
	if col < len((*grid)[0])-1 && (*grid)[row][col+1] == '1' {
		explore(grid, row, col+1)
	}
}

func numIslands(grid [][]byte) int {
	count := 0
	for i, row := range grid {
		for j, val := range row {
			if val == '1' {
				explore(&grid, i, j)
				count++
			}
		}

	}
	return count
}

func main() {
	grid1 := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	grid2 := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '0', '0', '0', '1'},
		{'0', '0', '1', '1', '1'},
	}
	fmt.Println(numIslands(grid1))
	fmt.Println(numIslands(grid2))
}
