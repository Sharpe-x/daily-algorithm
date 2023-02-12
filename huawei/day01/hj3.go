package main

import (
	"fmt"
	"sort"
)

func main() {
	dupRemoveTwo()
}

func dupRemoveOne() {
	var (
		num, tmp int
		nums     []int
	)
	fmt.Scanf("%d", &num)
	for i := 0; i < num; i++ {
		fmt.Scanf("%d", &tmp)
		nums = append(nums, tmp)
	}

	sort.Ints(nums)
	nums = dupRemove(nums)
	for _, v := range nums {
		fmt.Println(v)
	}
}

func dupRemove(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	var ans []int
	ans = append(ans, nums[0])
	for i := 1; i < len(nums); i++ {
		if nums[i] != ans[len(ans)-1] {
			ans = append(ans, nums[i])
		}
	}

	return ans
}

func dupRemoveTwo() {
	intSlice := make([]bool, 1000)

	var num, tmp int
	fmt.Scanf("%d", &num)
	for i := 0; i < num; i++ {
		fmt.Scanf("%d", &tmp)
		intSlice[tmp] = true
	}

	for k, v := range intSlice {
		if v {
			fmt.Println(k)
		}
	}
}
