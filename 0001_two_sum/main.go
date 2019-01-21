package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	result := make([]int, 2)

	for i := 0; i < len(nums); i++ {
		var curr = nums[i]
		var diff = target - curr
		result[0] = i
		if _, ok := m[diff]; ok {
			result[1] = m[diff]
			return result
		} else {
			m[curr] = i
		}
	}
	return result
}


func main() {
	var arr = []int{2, 7, 11, 15}
	fmt.Println(twoSum(arr, 9))
}
