package main

import "fmt"

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := 0
	for left < right {
		min := height[left]
		width := right - left
		if height[left] > height[right] {
			min = height[right]
			right--
		} else {
			left++
		}
		currArea := min * width
		if currArea >= maxArea {
			maxArea = currArea
		}
	}
	return maxArea
}

func main() {
	fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))
}
