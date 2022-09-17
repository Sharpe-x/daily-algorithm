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

// 二叉树层序遍历
func levelOrder(root *TreeNode) [][]int {

	res := make([][]int, 0, 0)

	queue := list.New()
	if root != nil {
		queue.PushBack(root)

		for queue.Len() != 0 {

			levelNum := queue.Len()
			var levelNums []int
			for levelNum > 0 {

				node := queue.Remove(queue.Front()).(*TreeNode)
				levelNums = append(levelNums, node.Val)
				levelNum--

				if node.Left != nil {
					queue.PushBack(node.Left)
				}

				if node.Right != nil {
					queue.PushBack(node.Right)
				}
			}

			res = append(res, levelNums)
		}

	}

	return res
}

func levelOrderBottom(root *TreeNode) [][]int {

	var ans [][]int

	queue := list.New()
	if root != nil {
		queue.PushBack(root)
	}

	for queue.Len() != 0 {
		levelNum := queue.Len()
		var levelNums []int
		for levelNum > 0 {
			node := queue.Remove(queue.Front()).(*TreeNode)
			levelNums = append(levelNums, node.Val)
			levelNum--

			if node.Left != nil {
				queue.PushBack(node.Left)
			}

			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		ans = append(ans, levelNums)
	}

	reverse(ans)
	return ans
}

func reverse(nums [][]int) {
	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

// 199 二叉树的右视图
func rightSideView(root *TreeNode) []int {

	var ans []int
	queue := list.New()
	if root != nil {
		queue.PushBack(root)
	}

	var tmp int
	for queue.Len() != 0 {
		levelNum := queue.Len()
		for levelNum > 0 {
			node := queue.Remove(queue.Front()).(*TreeNode)
			levelNum--
			tmp = node.Val

			if node.Left != nil {
				queue.PushBack(node.Left)
			}

			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		ans = append(ans, tmp)
	}
	return ans
}

// 637 二叉树的层平均值
func averageOfLevels(root *TreeNode) []float64 {
	var ans []float64
	queue := list.New()
	if root != nil {
		queue.PushBack(root)
	}
	for queue.Len() != 0 {
		levelNum := queue.Len()
		sum, len := 0, levelNum

		for levelNum > 0 {
			node := queue.Remove(queue.Front()).(*TreeNode)
			sum += node.Val
			levelNum--

			if node.Left != nil {
				queue.PushBack(node.Left)
			}

			if node.Right != nil {
				queue.PushBack(node.Right)
			}

		}

		ans = append(ans, float64(sum)/float64(len))
	}

	return ans
}

type Node struct {
	Val      int
	Children []*Node
}

//429. N 叉树的层序遍历
func levelOrder429(root *Node) [][]int {
	var ans [][]int

	queue := list.New()
	if root != nil {
		queue.PushBack(root)
	}

	for queue.Len() != 0 {
		levelNum := queue.Len()
		var levelAns []int
		for levelNum > 0 {
			node := queue.Remove(queue.Front()).(*Node)
			levelNum--
			levelAns = append(levelAns, node.Val)

			for _, v := range node.Children {
				if v != nil {
					queue.PushBack(v)
				}
			}
		}
		ans = append(ans, levelAns)

	}

	return ans
}
