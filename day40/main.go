/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

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

func findBottomLeftValue(root *TreeNode) int {
	queue := list.New()
	queue.PushBack(root)

	result := 0
	for queue.Len() != 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			elem := queue.Remove(queue.Front()).(TreeNode)
			if i == 0 {
				result = elem.Val
			}

			if elem.Left != nil {
				queue.PushBack(elem.Left)
			}

			if elem.Right != nil {
				queue.PushBack(elem.Right)
			}
		}
	}
	return result
}

func hasPathSum(root *TreeNode, targetSum int) bool {

}

func hasPath(node *TreeNode, sum int) bool {
	if node.Left == nil && node.Right == nil && sum == 0 {
		return true
	}

	if node.Left != nil {
		sum -= node.Left.Val
		if hasPath(node.Left, sum) {
			return true
		}
		sum += node.Left.Val
	}

	if node.Right != nil {
		sum -= node.Right.Val
		if hasPath(node.Right, sum) {
			return true
		}
		sum += node.Right.Val
	}

	return false
}
