package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil {
		return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return false
}

func main() {
	a3 := &TreeNode{3, nil, nil}
	a2 := &TreeNode{2, nil, nil}
	a1 := &TreeNode{1, a3, a2}

	b3 := &TreeNode{3, nil, nil}
	b2 := &TreeNode{2, nil, nil}
	b1 := &TreeNode{1, b3, b2}

	fmt.Println(isSameTree(a1, b1))
}
