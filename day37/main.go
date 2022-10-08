package main

func main() {

}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func countNodes(root *TreeNode) int {
	return getNodeNum(root)
}

func getNodeNum(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := getNodeNum(node.Left)
	right := getNodeNum(node.Right)

	return left + right + 1
}
