package main

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return traversal(nums)
}

func traversal(nums []int) *TreeNode {
	if len(nums) == 1 {
		return &TreeNode{
			Val: nums[0],
		}
	}

	maxNum := -1
	maxIndex := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxNum {
			maxIndex = i
			maxNum = nums[i]
		}
	}

	node := &TreeNode{
		Val: maxNum,
	}

	if maxIndex > 0 {
		node.Left = traversal(nums[:maxIndex])
	}

	if maxIndex < len(nums)-1 {
		node.Right = traversal(nums[maxIndex+1:])
	}

	return node
}
