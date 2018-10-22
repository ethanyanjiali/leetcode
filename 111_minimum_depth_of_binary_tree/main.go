package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var queue []*TreeNode
	height := 1
	queue = append(queue, root)
	// BFS
	for len(queue) != 0 {
		var nextLevel []*TreeNode
		// Iterate through current level
		for _, node := range queue {
			if node.Left == nil && node.Right == nil {
				return height
			}
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		height++
		queue = nextLevel
	}
	return height - 1
}

func main() {
	a5 := &TreeNode{2, nil, nil}
	a4 := &TreeNode{3, a5, nil}
	a3 := &TreeNode{3, a4, nil}
	a2 := &TreeNode{2, nil, nil}
	a1 := &TreeNode{1, a3, a2}

	fmt.Println(minDepth(a1))
	fmt.Println()
}
