package main

import "sort"

func main() {

}

func canConstruct(ransomNote string, magazine string) bool {
	nums := [26]int{}

	for _, v := range magazine {
		nums[v-'a']++
	}

	for _, v := range ransomNote {
		nums[v-'a']--
		if nums[v-'a'] < 0 {
			return false
		}
	}

	return true
}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}

		// a 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			leftNum, rightNum := nums[left], nums[right]
			if nums[i]+leftNum+rightNum > 0 {
				right--
			} else if nums[i]+leftNum+rightNum < 0 {
				left++
			} else {
				res = append(res, []int{nums[i], leftNum, rightNum})
				for right > left && nums[right] == rightNum {
					right--
				}

				for right > left && nums[left] == leftNum {
					left++
				}
			}
		}

	}

	return res
}
