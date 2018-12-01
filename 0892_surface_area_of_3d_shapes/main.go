package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func surfaceArea(grid [][]int) int {
	area := 0
    for i := range grid {
    	for j := range grid[i] {
    		height := grid[i][j]
    		if height > 0 {
    			area += height * 4 + 2
    		}
    		if i > 0 {
    			area -= min(grid[i-1][j], height) * 2
    		}
    		if j > 0 {
    			area -= min(grid[i][j-1], height) * 2
    		}
    	}
    }
    return area
}

func main() {
	fmt.Println(surfaceArea([][]int{
		{2},
	}))
	fmt.Println(surfaceArea([][]int{
		{1, 2}, {3, 4},
	}))
}