package main

import "math"

// https://leetcode.cn/problems/minimum-size-subarray-sum/submissions/
// https://www.bilibili.com/video/BV1tZ4y1q7XE?vd_source=4cbc8ea75af98bd8e8f573b7b783e54b

func main() {

}

// minSubArrayLen 暴力解法
func minSubArrayLen1(target int, nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	ans := math.MaxInt
	for i := 0; i < len(nums); i++ {
		total := 0
		for j := i; j < len(nums); j++ {
			total = total + nums[j]
			if total >= target {
				ans = min(ans, j-i+1)
				break
			}
		}
	}

	if ans == math.MaxInt {
		return 0
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 滑动窗口
func minSubArrayLen(target int, nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	res := math.MaxInt
	total, start := 0, 0
	for end := 0; end < len(nums); end++ { // end 表示滑动窗口数组的右边界
		total += nums[end]
		for total >= target { // 知道不满足条件才移动end
			res = min(res, end-start+1)
			total = total - nums[start]
			start++
		}
	}

	if res == math.MaxInt { // 全部加起来都不大于target
		return 0
	}

	return res
}
