package main

import "fmt"

func singleNumber(nums []int) int {
    res := 0
    for _, num := range nums {
    	res ^= num
    }
    return res
}

func main() {
	fmt.Println(singleNumber([]int{2,2,1}))
	fmt.Println(singleNumber([]int{2,2,3,5,3,4,4,1,1}))
}

