package main

import (
	"fmt"
	"sort"
	"strconv"
)

func nextClosestTime(time string) string {
    original := []int{
    	int(time[0] - '0'),
    	int(time[1] - '0'),
    	int(time[3] - '0'),
    	int(time[4] - '0'),
    }
    digits := make([]int, 4)
    copy(digits, original)
    sort.Ints(digits)
    smallest := strconv.Itoa(digits[0])
    for _, digit := range digits {
    	if digit > original[3] {
    		return time[:4] + strconv.Itoa(digit)
    	}
    }
    for i := 0; i < len(digits) && digits[i] < 6; i++ {
    	if digits[i] > original[2] {
    		return time[:3] + strconv.Itoa(digits[i]) + smallest
    	}
    }
    for i := 0; i < len(digits); i++ {
    	if ((original[0] == 2 && digits[i] < 4) || original[0] != 2) && digits[i] > original[1] {
    		return time[:1] + strconv.Itoa(digits[i]) + ":" + smallest + smallest
    	}
    }
    for i := 0; i < len(digits) && digits[i] < 3; i++ {
    	if digits[i] > original[0] {
    		return strconv.Itoa(digits[i]) + smallest + ":" + smallest + smallest
    	}
    }
    return smallest + smallest + ":" + smallest + smallest
}

func main() {
	fmt.Println(nextClosestTime("10:54"))
	fmt.Println(nextClosestTime("10:45"))
	fmt.Println(nextClosestTime("10:59"))
	fmt.Println(nextClosestTime("12:22"))
	fmt.Println(nextClosestTime("23:59"))
}