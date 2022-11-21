package main

import "sort"

func main() {

}

func combinationSum(candidates []int, target int) [][]int {
	var (
		path []int
		ans  [][]int
	)
	sort.Ints(candidates)
	backtracking(0, 0, target, candidates, path, &ans)
	return ans
}

func backtracking(startIndex, sum, target int, nums, path []int, ans *[][]int) {
	if sum == target {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*ans = append(*ans, tmp)
		return
	}

	if sum > target {
		return
	}

	for i := startIndex; i < len(nums) && sum+nums[i] <= target; i++ {
		sum += nums[i]
		path = append(path, nums[i])
		backtracking(i, sum, target, nums, path, ans)
		sum -= nums[i]
		path = path[:len(path)-1]
	}
}
