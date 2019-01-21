package main

import (
	"sort"
	"fmt"
)

func threeSum(nums []int) [][]int {
    sort.Slice(nums, func(i, j int) bool {
    	return nums[i] < nums[j]
    })
    res := [][]int{}
    for i := 0; i < len(nums); i++ {
    	if i != 0 && nums[i] == nums[i - 1] {
    		continue
    	}
    	res = twoSum(nums, i + 1, nums[i], res)
    }
    return res
}

func twoSum(nums []int, start, first int, res [][]int) [][]int {
    left := start
    right := len(nums) - 1
    for left < right {
    	sum := nums[left] + nums[right]
    	if sum == 0 - first {
    		res = append(res, []int{first, nums[left], nums[right]})
    		left++
    		right--
    		for left < len(nums) && left > 0 && nums[left] == nums[left - 1] {
    			left++
    		}
            for right >= 0 && right < len(nums) - 1 && nums[right] == nums[right + 1] {
    			right--
    		}
    	} else if sum > 0 - first {
    		right--
    	} else if sum < 0 - first {
    		left++
    	}
    }
    return res
}

func main() {
	arr := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(arr))
}