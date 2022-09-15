package main

func main() {

}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {

	// 确定递归终止的条件
	if root == nil {
		return nil
	}

	var ans []int
	ans = append(ans, root.Val)
	// 确定递归函数的参数和返回值
	// 确定单层逻辑
	tmp := preorderTraversal(root.Left)
	for _, v := range tmp {
		ans = append(ans, v)
	}

	tmp = preorderTraversal(root.Right)
	for _, v := range tmp {
		ans = append(ans, v)
	}

	return ans
}

func postorderTraversal(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	var ans []int
	tmp := postorderTraversal(root.Left)
	for _, v := range tmp {
		ans = append(ans, v)
	}
	tmp = postorderTraversal(root.Right)
	for _, v := range tmp {
		ans = append(ans, v)
	}
	ans = append(ans, root.Val)

	return ans
}

func inorderTraversal(root *TreeNode) []int {

	if root == nil {
		return nil
	}
	var ans []int
	tmp := inorderTraversal(root.Left)
	for _, v := range tmp {
		ans = append(ans, v)
	}
	ans = append(ans, root.Val)
	tmp = inorderTraversal(root.Right)
	for _, v := range tmp {
		ans = append(ans, v)
	}
	return ans
}
