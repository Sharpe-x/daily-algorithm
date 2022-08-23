package main

import "sort"

func main() {

}

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)

	sort.Ints(nums)
	for k := 0; k < len(nums)-1; k++ {

		// 剪枝
		if nums[k] > target && target >= 0 && nums[k] >= 0 {
			break
		}

		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		for i := k + 1; i < len(nums)-1; i++ {
			twoSum := nums[i] + nums[k]
			// 剪枝
			if twoSum > target && target >= 0 && twoSum >= 0 {
				break
			}

			if i > k+1 && nums[i] == nums[i-1] {
				continue
			}
			left, right := i+1, len(nums)-1
			for left < right {
				leftNum, rightNum := nums[left], nums[right]
				sum := twoSum + leftNum + rightNum
				if sum < target {
					left++
				} else if sum > target {
					right--
				} else {
					res = append(res, []int{nums[k], nums[i], leftNum, rightNum})
					for left < right && nums[left] == leftNum {
						left++
					}
					for left < right && nums[right] == rightNum {
						right--
					}
				}
			}

		}
	}

	return res
}
