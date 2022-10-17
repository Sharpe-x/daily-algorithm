package main

func main() {

}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	return sumOfLeft(root)
}

func sumOfLeft(node *TreeNode) int {
	if node == nil {
		return 0
	}

	if node.Left == nil && node.Right == nil {
		return 0
	}

	var leftNum, rightNum int
	leftNum = sumOfLeft(node.Left)
	if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
		leftNum = node.Left.Val
	}
	rightNum = sumOfLeft(node.Right)
	return leftNum + rightNum
}
