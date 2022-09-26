package main

import (
	"container/list"
)

func main() {

}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isSymmetricRec(root *TreeNode) bool {

	if root == nil {
		return false
	}

	return innerIsSymmetric(root.Left, root.Right)
}

func innerIsSymmetric(left, right *TreeNode) bool {
	if left == nil && right != nil {
		return false
	} else if right == nil && left != nil {
		return false
	} else if left == nil && right == nil {
		return true
	} else if left.Val != right.Val {
		return false
	}
	outSide := innerIsSymmetric(left.Left, right.Right)
	inSide := innerIsSymmetric(left.Right, right.Left)

	return outSide && inSide
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}

	queue := list.New()
	queue.PushBack(root.Left)
	queue.PushBack(root.Right)

	for queue.Len() != 0 {
		left := queue.Remove(queue.Front()).(*TreeNode)
		right := queue.Remove(queue.Front()).(*TreeNode)

		if left == nil && right == nil {
			continue
		}

		if left == nil && right != nil {
			return false
		} else if left != nil && right == nil {
			return false
		} else if left.Val != right.Val {
			return false
		}

		queue.PushBack(left.Left)
		queue.PushBack(right.Right)
		queue.PushBack(left.Right)
		queue.PushBack(right.Left)
	}
	return true
}
