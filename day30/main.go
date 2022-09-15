package main

import (
	"container/list"
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	reverse(nums)
	fmt.Println(nums)
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}

	stack := list.New()
	stack.PushBack(root)

	for stack.Len() != 0 {
		back := stack.Remove(stack.Back()).(*TreeNode)
		ans = append(ans, back.Val)

		if back.Right != nil {
			stack.PushBack(back.Right)
		}
		if back.Left != nil {
			stack.PushBack(back.Left)
		}
	}

	return ans
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 前序 中左右
// 后序 中右左 再做一次翻转 左右中 就是
func postorderTraversal(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}

	stack := list.New()
	stack.PushBack(root)
	for stack.Len() != 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		ans = append(ans, node.Val)

		if node.Left != nil {
			stack.PushBack(node.Left)
		}

		if node.Right != nil {
			stack.PushBack(node.Right)
		}
	}

	reverse(ans)

	return ans
}

func reverse(nums []int) {

	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

// 为什么前序遍历的代码，不能和中序遍历通用呢，
//  因为前序遍历的顺序是中左右，先访问的元素是中间节点，
// 要处理的元素也是中间节点，所以才能写出相对简洁的代码，
// 因为要访问的元素和要处理的元素顺序是一致的，都是中间节点。

// 再看看中序遍历，中序遍历是左中右，
// 先访问的是二叉树顶部的节点，
//然后一层一层向下访问，直到到达树左面的最底部，
// 再开始处理节点（也就是在把节点的数值放进result数组中），这就造成了处理顺序和访问顺序是不一致的。
//那么在使用迭代法写中序遍历，就需要借用指针的遍历来帮助访问节点，栈则用来处理节点上的元素。
func inorderTraversal(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}

	stack := list.New()
	cur := root

	for cur != nil || stack.Len() != 0 {
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		} else {
			cur = stack.Remove(stack.Back()).(*TreeNode)
			ans = append(ans, cur.Val)
			cur = cur.Right
		}
	}

	return ans
}
