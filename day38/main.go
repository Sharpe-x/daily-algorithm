package main

import (
	"container/list"
	"strconv"
)

func main() {

}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return getHight(root) != -1
}

func getHight(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftHight := getHight(node.Left)
	if leftHight == -1 {
		return -1
	}

	rightHight := getHight(node.Right)
	if rightHight == -1 {
		return -1
	}

	if absSub(leftHight, rightHight) > 1 {
		return -1
	}

	return max(leftHight, rightHight) + 1
}

func absSub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	res := []string{}
	path := list.New()
	traversal(root, path, &res)
	return res
}

func traversal(node *TreeNode, path *list.List, result *[]string) {

	path.PushBack(node.Val)
	if node.Left == nil && node.Right == nil {
		var sPath string
		cur := path.Front()
		for i := 0; i < path.Len()-1; i++ {
			sPath += strconv.Itoa(cur.Value.(int))
			sPath += "->"
			cur = cur.Next()
		}
		sPath += strconv.Itoa(cur.Value.(int))
		*result = append(*result, sPath)
		return
	}

	if node.Left != nil {
		traversal(node.Left, path, result)
		path.Remove(path.Back())
	}

	if node.Right != nil {
		traversal(node.Right, path, result)
		path.Remove(path.Back())
	}
}

func binaryTreePaths2(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	res := []string{}
	path := []int{}
	traversal1(root, path, &res)
	return res
}

func traversal1(node *TreeNode, path []int, result *[]string) {

	path = append(path, node.Val)
	if node.Left == nil && node.Right == nil {
		var sPath string
		for i := 0; i < len(path)-1; i++ {
			sPath += strconv.Itoa(path[i])
			sPath += "->"
		}
		sPath += strconv.Itoa(path[len(path)-1])
		*result = append(*result, sPath)
		return
	}

	if node.Left != nil {
		traversal1(node.Left, path, result)
	}

	if node.Right != nil {
		traversal1(node.Right, path, result)
	}
}

func binaryTreePath(root *TreeNode) []string {
	res := make([]string, 0)
	var travel func(node *TreeNode, s string)
	travel = func(node *TreeNode, s string) {
		if node.Left == nil && node.Right == nil {
			v := s + strconv.Itoa(node.Val)
			res = append(res, v)
			return
		}

		s = s + strconv.Itoa(node.Val) + "->"
		if node.Left != nil {
			travel(node.Left, s)
		}
		if node.Right != nil {
			travel(node.Right, s)
		}
	}

	travel(root, "")
	return res
}
