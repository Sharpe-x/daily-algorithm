package main

func main() {

}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func minDepth(root *TreeNode) int {
	return getMinDepth(root)
}

func getMinDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftDepth := getMinDepth(node.Left)
	rightDepth := getMinDepth(node.Right)

	if node.Left != nil && node.Right == nil {
		return 1 + leftDepth
	}

	if node.Left == nil && node.Right != nil {
		return 1 + rightDepth
	}

	return 1 + min(leftDepth, rightDepth)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
