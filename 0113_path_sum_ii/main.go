package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findPath(node *TreeNode, res [][]int, path []int, sum int) {
	if node == nil {
		return
	}
	path = append(path, node.Val)
	if node.Left == nil && node.Right == nil && node.Val == sum {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
	findPath(node.Left, res, path, sum-node.Val)
	findPath(node.Right, res, path, sum-node.Val)
	path = path[0 : len(path)-1]
}

func pathSum(root *TreeNode, sum int) [][]int {
	res := [][]int{}
	path := []int{}
	findPath(root, res, path, sum)
	return res
}

func main() {

}
