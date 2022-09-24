package main

import "container/list"

type Node struct {
	Val      int
	Children []*Node
}

func preorderRec(root *Node) []int {
	var ans []int
	if root == nil {
		return ans
	}
	ans = append(ans, root.Val)
	for _, v := range root.Children {
		tmp := preorderRec(v)
		ans = append(ans, tmp...)
	}

	return ans
}

func preorder(root *Node) []int {
	var ans []int
	if root == nil {
		return ans
	}

	stack := list.New()
	stack.PushBack(root)
	for stack.Len() != 0 {
		node := stack.Remove(stack.Back()).(*Node)
		ans = append(ans, node.Val)

		for i := len(node.Children) - 1; i >= 0; i-- {
			stack.PushBack(node.Children[i])
		}

	}

	return ans
}

func postorderRec(root *Node) []int {
	var ans []int
	if root == nil {
		return ans
	}

	for _, v := range root.Children {
		ans = append(ans, postorderRec(v)...)
	}
	ans = append(ans, root.Val)

	return ans
}

// 前序是中左右
// 后序是左右中 就是 中右左做一次翻转
func postorder(root *Node) []int {
	var ans []int
	if root == nil {
		return ans
	}

	stack := list.New()
	stack.PushBack(root)

	for stack.Len() != 0 {
		node := stack.Remove(stack.Back()).(*Node)
		ans = append(ans, node.Val)

		for _, v := range node.Children {
			stack.PushBack(v)
		}
	}

	reverse(ans)
	return ans
}

func reverse(nums []int) {
	l, r := 0, len(nums)-1
	for l <= r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
