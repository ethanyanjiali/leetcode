package main

import "fmt"

func totalFruit2(tree []int) int {
	counters := map[int]int{}
	left := 0
	max := 0
	for right, treeType := range tree {
		counters[treeType] += 1
		for len(counters) > 2 {
			counters[tree[left]] -= 1
			if counters[tree[left]] == 0 {
				delete(counters, tree[left])
			}
			left++
		}
		if right - left + 1 > max {
			max = right - left + 1
		}
	}
	return max
}

func totalFruit(tree []int) int {
	baskets := [2]int{-1, -1} // baskets[1] is the last type, baskets[0] is the 2nd last type
	amount := [2]int{0, 0} // amount[1] is the amount of fruit from the last type, amount[0] is the for 2nd last type
	consecutive := 0 // how many consecutive trees does the baskets[1] has
	max := 0
	for _, treeType := range tree {
		if treeType == baskets[0] {
			// if current type equls to 2nd last type, we need to swap last type and second last type
			baskets[0], baskets[1] = baskets[1], treeType
			amount[0], amount[1] = amount[1], amount[0] + 1
			consecutive = 1
		} else if treeType == baskets[1] {
			// if current type equls to last type, then just need to increment the amount
			amount[1] += 1
			consecutive += 1 // since current type == last type, we have one more consecutive element of last type
		} else {
			// otherwise, we run into a new type, let's ditch the 2nd last type and push in this new type
			baskets[0], baskets[1] = baskets[1], treeType
			amount[0], amount[1] = consecutive, 1
			consecutive = 1
		}
		// max is just the total of last type and 2nd last type
		if max < amount[0] + amount[1] {
			max = amount[0] + amount[1]
		}
	}
	return max
}

func main() {
	tree1 := []int{1, 2, 1}
	fmt.Println(totalFruit2(tree1))
	tree2 := []int{1,2,3,2,2}
	fmt.Println(totalFruit2(tree2))
	tree3 := []int{3,3,3,1,2,1,1,2,3,3,4}
	fmt.Println(totalFruit2(tree3))
	tree4 := []int{0, 1, 1}
	fmt.Println(totalFruit2(tree4))
	tree5 := []int{1,0,1,4,1,4,1,2,3}
	fmt.Println(totalFruit2(tree5))
}