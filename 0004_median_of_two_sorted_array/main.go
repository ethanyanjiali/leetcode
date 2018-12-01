package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    len1 := len(nums1)
    len2 := len(nums2)
    total := len1 + len2
    if total % 2 == 0 {
    	return float64(findKth(total / 2, nums1, nums2) + findKth((total / 2) + 1, nums1, nums2))/ 2.0
    } else {
    	return float64(findKth(total / 2 + 1, nums1, nums2))
    }
}

func findKth(k int, nums1, nums2 []int) int {
	p1, p2 := 0, 0
	for i := 0; i < k - 1; i++ {
		if p1 >= len(nums1) {
			p2++
		} else if p2 >= len(nums2) {
			p1++
		} else if nums1[p1] > nums2[p2] {
			p2++
		} else {
			p1++
		}
	}
	if p1 >= len(nums1) {
		return nums2[p2]
	}
	if p2 >= len(nums2) {
		return nums1[p1]
	}
	if nums1[p1] > nums2[p2] {
		return nums2[p2]
	}
	return nums1[p1]
}

func main() {
	// nums1 := []int{1, 3, 5, 7}
	// nums2 := []int{2, 4, 6, 8}
	// fmt.Println(findMedianSortedArrays(nums1, nums2))
	nums3 := []int{1, 3}
	nums4 := []int{2}
	fmt.Println(findMedianSortedArrays(nums3, nums4))
}
