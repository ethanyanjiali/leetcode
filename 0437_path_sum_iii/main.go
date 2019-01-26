package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findPath(node *TreeNode, target int, path []int) int {
	if node == nil {
		return 0
	}
	path = append(path, node.Val)
	sum := 0
	count := 0
	for i := len(path) - 1; i >= 0; i-- {
		sum += path[i]
		if sum == target {
			count++
		}
	}
	left := findPath(node.Left, target, path)
	right := findPath(node.Right, target, path)
	path = path[0 : len(path)-1]
	return left + right + count
}

func pathSum(root *TreeNode, sum int) int {
	path := []int{}
	return findPath(root, sum, path)
}
