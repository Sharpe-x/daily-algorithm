package main

/*
https://leetcode.cn/problems/squares-of-a-sorted-array/submissions/
*/

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, -1, -9, 0}
	fmt.Println(sortedSquaresOn(nums))
	fmt.Println(sortedSquares(nums))
}

// 暴力
func sortedSquaresOn(nums []int) []int {
	for i, v := range nums {
		nums[i] = v * v
	}

	sort.Ints(nums)

	return nums
}

// 数组其实是有序的， 只不过负数平方之后可能成为最大数了。

// 那么数组平方的最大值就在数组的两端，不是最左边就是最右边，不可能是中间。

// 此时可以考虑双指针法了，i指向起始位置，j指向终止位置。

// 定义一个新数组result，和A数组一样的大小，让k指向result数组终止位置。

// 如果A[i] * A[i] < A[j] * A[j] 那么result[k--] = A[j] * A[j]; 。

// 如果A[i] * A[i] >= A[j] * A[j] 那么result[k--] = A[i] * A[i]; 。
func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))

	var (
		i    = 0
		j, k = len(nums) - 1, len(nums) - 1
	)

	for i <= j {
		if nums[i]*nums[i] > nums[j]*nums[j] {
			result[k] = nums[i] * nums[i]
			i++
		} else {
			result[k] = nums[j] * nums[j]
			j--
		}
		k--
	}

	return result
}
