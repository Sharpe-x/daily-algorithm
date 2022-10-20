package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func main() {

}

func buildTree(inorder []int, postOrder []int) *TreeNode {
	if len(inorder) == 0 || len(postOrder) == 0 || len(postOrder) != len(inorder) {
		return nil
	}

	return build(inorder, postOrder)
}

func build(inorder, postOrder []int) *TreeNode {
	if len(postOrder) == 0 {
		return nil
	}

	value := postOrder[len(postOrder)-1]
	node := &TreeNode{
		Val: value,
	}

	if len(postOrder) == 1 {
		return node
	}

	i := 0
	for i = 0; i < len(inorder); i++ {
		if inorder[i] == value {
			break
		}
	}

	leftInorder := inorder[:i]
	rightOrder := inorder[i+1:]

	node.Left = build(leftInorder, postOrder[:len(leftInorder)])
	node.Right = build(rightOrder, postOrder[len(leftInorder):len(postOrder)-1])
	return node
}
