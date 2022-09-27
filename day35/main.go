package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return getDepth(root)
}

func getDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := getDepth(node.Left)
	right := getDepth(node.Right)

	return 1 + max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
