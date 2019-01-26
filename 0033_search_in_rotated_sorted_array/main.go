package main

func search(nums []int, target int) int {
	min, max := 0, len(nums) - 1
	// when min == max, there could still be a solution where min=max=mid
	for min <= max {
		mid := (min + max) / 2
		// if the target is found, return mid
		if nums[mid] == target {
			return mid
		}
		// if nums[min] <= nums[mid], the left side of the sequence is ordered
		if nums[min] <= nums[mid] {
			// if the target falls between the ordered part, we can
			// safely infer that max would the the mid - 1
			if nums[mid] > target && nums[min] <= target {
				max = mid - 1
			// otherwise, target must on the other side
			} else {
				min = mid + 1
			}
		// if the left side is not ordered, the the right side must be ordered
		} else {
			if nums[mid] < target && target <= nums[max] {
				min = mid + 1
			} else {
				max = mid - 1
			}
		}
	}
	return -1
}

func main() {

}