package main

import "fmt"

/*

给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
https://leetcode.cn/problems/binary-search/solution/dai-ma-sui-xiang-lu-dai-ni-xue-tou-er-fe-ajox/

*/

func main() {
	nums := []int{1, 2, 4, 5, 6, 6, 7, 8, 9, 10, 11}
	fmt.Print(search(nums, 7))
}

func search(nums []int, target int) int {

	var (
		l = 0
		r = len(nums) - 1
	)

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
