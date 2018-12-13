package main

import "fmt"

func findK(garden []bool, k int) bool {
	last := -1
    for i := 0; i < len(garden); i++ {
        blooming := garden[i]
		if blooming {
			if last != -1 && i - last - 1 == k {
				return true
			}
			last = i
		}
	}
	return false
}

func kEmptySlots(flowers []int, k int) int {
    garden := make([]bool, len(flowers))
    for day, pos := range flowers {
    	garden[pos - 1] = true
    	if findK(garden, k) {
    		return day + 1
    	}
    }
    return -1
}

func kEmptySlots2(flowers []int, k int) int {
    garden := make([]int, len(flowers))
    for day, pos := range flowers {
    	garden[pos - 1] = day
    }
    left := 0
    right := left + k + 1
    res := -1
    i := 0
    for ; right < len(garden); i++ {
    	if garden[i] < garden[left] || garden[i] <= garden[right] {
			if i == right {
				if garden[left] > garden[right] {
				    res = right
				} else {
					res = left
				}
			}
			left = i
			right = left + k + 1
    	}
    }
    return res
}

func main() {
	fmt.Println(kEmptySlots2([]int{1,3,2}, 1))
	fmt.Println(kEmptySlots2([]int{6,10,7,1,9,8,4,3,5,2}, 3))
	fmt.Println(kEmptySlots2([]int{3,9,2,8,1,6,10,5,4,7}, 1))
}