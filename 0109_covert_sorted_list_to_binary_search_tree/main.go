package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var curr *ListNode

func buildTree(start int, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := start + (end-start)/2
	root := &TreeNode{}
	root.Left = buildTree(start, mid-1)
	root.Val = curr.Val
	curr = curr.Next
	root.Right = buildTree(mid+1, end)
	return root
}

func sortedListToBST(head *ListNode) *TreeNode {
	length := 0
	curr = head
	for curr != nil {
		curr = curr.Next
		length++
	}
	curr = head
	return buildTree(0, length-1)
}

func main() {
	a5 := ListNode{9, nil}
	a4 := ListNode{5, &a5}
	a3 := ListNode{0, &a4}
	a2 := ListNode{-3, &a3}
	a1 := ListNode{-10, &a2}
	sortedListToBST(&a1)
}
