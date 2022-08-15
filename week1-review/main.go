package main

import "math"

func main() {

}

// 左闭右闭区间
func search(nums []int, target int) int {

	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

func removeElement(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {

		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}

func sortedSquares(nums []int) []int {
	newNums := make([]int, len(nums))

	i, j := 0, len(nums)-1
	k := len(nums) - 1
	for k >= 0 {
		if nums[i]*nums[i] < nums[j]*nums[j] {
			newNums[k] = nums[j] * nums[j]
			j--

		} else {
			newNums[k] = nums[i] * nums[i]
			i++
		}
		k--
	}

	return newNums
}

func minSubArrayLen(target int, nums []int) int {

	l, r := 0, 0
	ans := math.MaxInt
	sum := 0
	for r < len(nums) {
		sum = sum + nums[r]
		for sum >= target {
			ans = min(ans, r-l+1)
			sum = sum - nums[l]
			l++
		}
		r++
	}

	if ans == math.MaxInt {
		return 0
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func generateMatrix(n int) [][]int {

	nums := make([][]int, n)
	for i := 0; i < n; i++ {
		nums[i] = make([]int, n)
	}

	circleNum := n / 2
	target := 1
	startx, starty, offset := 0, 0, 1
	i, j := 0, 0
	for circleNum > 0 {

		for j = starty; j < n-offset; j++ {
			nums[startx][j] = target
			target++
		}

		for i = startx; i < n-offset; i++ {
			nums[i][j] = target
			target++
		}

		for ; j > starty; j-- {
			nums[i][j] = target
			target++
		}

		for ; i > startx; i-- {
			nums[i][j] = target
			target++
		}

		circleNum--
		startx++
		starty++
		offset++
	}

	if n%2 == 1 {
		nums[n/2][n/2] = target
	}

	return nums
}
