package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	stack := []*TreeNode{root}
	curr := &TreeNode{0, nil, nil}
	for len(stack) != 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// DFS from left to right, push right to the stack first
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		// push left to the stack
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
		// put top to curr's right and clear curr's left
		curr.Left = nil
		curr.Right = top
		// move on to the next
		curr = curr.Right
	}
}

func printResult(head *TreeNode) {
	if head != nil {
		fmt.Printf("%d", head.Val)
		for head.Right != nil {
			head = head.Right
			fmt.Printf(" -> %d", head.Val)
		}
	}
	fmt.Println()
}

func main() {
	a5 := &TreeNode{5, nil, nil}
	a4 := &TreeNode{4, a5, nil}
	a3 := &TreeNode{3, a4, nil}
	a2 := &TreeNode{2, nil, nil}
	a1 := &TreeNode{1, a3, a2}
	flatten(a1)
	printResult(a1)
}
