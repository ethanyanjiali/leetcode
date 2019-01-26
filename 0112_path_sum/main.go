package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(node *TreeNode, sum int) bool {
	if node == nil {
		return false
	}
	diff := sum - node.Val
	if diff == 0 && node.Left == nil && node.Right == nil {
		return true
	}
	return hasPathSum(node.Left, diff) || hasPathSum(node.Right, diff)
}

func main() {

}
