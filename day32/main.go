package main

import "container/list"

func main() {

}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func invertTreeRec(root *TreeNode) *TreeNode {

	if root == nil {
		return root
	}

	root.Left, root.Right = root.Right, root.Left
	invertTreeRec(root.Left)
	invertTreeRec(root.Right)

	return root
}

func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return root
	}

	stack := list.New()
	stack.PushBack(root)

	for stack.Len() != 0 {
		top := stack.Remove(stack.Back()).(*TreeNode)
		top.Left, top.Right = top.Right, top.Left

		if top.Left != nil {
			stack.PushBack(top.Left)
		}

		if top.Right != nil {
			stack.PushBack(top.Right)
		}

	}

	return root
}
