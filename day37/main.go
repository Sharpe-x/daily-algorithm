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

// 利用完全二叉树的特性
func getNodeNum2(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := node.Left
	right := node.Right

	leftHight, rightHight := 0, 0
	for left != nil {
		left = left.Left
		leftHight++
	}

	for right != nil {
		right = right.Right
		rightHight++
	}

	if leftHight == rightHight {
		return 2<<leftHight - 1
	}

	leftNum := getNodeNum2(node.Left)
	righNum := getNodeNum2(node.Right)
	return leftNum + righNum + 1
}
