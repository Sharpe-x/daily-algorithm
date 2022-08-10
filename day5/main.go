package main

import "fmt"

//https://leetcode.cn/problems/spiral-matrix-ii/
// 给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。

func main() {
	fmt.Println(generateMatrix(1))
	fmt.Println(generateMatrix(3))
	fmt.Println(generateMatrix(4))
}

// 模拟过程
// 坚持循环不变量原则，左闭右开
// 从左到右，从上到下，从右到左，从下到上
func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for k := 0; k < n; k++ {
		ans[k] = make([]int, n)
	}

	// 装圈数
	circleNum := n / 2
	startx, starty := 0, 0
	count := 1
	i, j := 0, 0
	offset := 1 // 偏移量

	for circleNum > 0 {
		// 从左到右 最后一个元素不处理
		for j = starty; j < n-offset; j++ {
			ans[startx][j] = count
			count++
		}

		// 从上到下  最后一个元素不处理
		for i = startx; i < n-offset; i++ {
			ans[i][j] = count
			count++
		}

		// 从右到左  最后一个元素不处理
		for ; j > starty; j-- {
			ans[i][j] = count
			count++
		}

		// 从下到上  最后一个元素不处理
		for ; i > startx; i-- {
			ans[i][j] = count
			count++
		}

		startx++
		starty++
		offset++
		circleNum--

	}

	if (n % 2) == 1 {
		ans[n/2][n/2] = count
	}

	return ans
}
