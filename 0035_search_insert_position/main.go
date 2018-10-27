package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	min := 0
	max := len(nums) - 1
	for min <= max {
		mid := min + (max-min)/2
		fmt.Printf("min: %d, max: %d, mid: %d\n", min, max, mid)
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			max = mid - 1
		} else if nums[mid] < target {
			min = mid + 1
		}
	}
	fmt.Printf("min: %d, max: %d\n", min, max)
	return min // nums[min] will always be larger/equal than target
}

func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 2))
	// 	min: 0, max: 3, mid: 1
	//  min: 0, max: 0, mid: 0
	//  min: 1, max: 0
	//  1
}
