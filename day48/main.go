package main

import "sort"

func main() {

}

func combinationSum2(candidates []int, target int) [][]int {
	var path []int
	var res [][]int
	used := make(map[int]bool)
	sort.Ints(candidates)
	backtracking(candidates, path, target, 0, 0, used, &res)
	return res
}

func backtracking(candidates, path []int, target, sum, startIndex int, used map[int]bool, res *[][]int) {
	if sum == target {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
	}

	if sum > target {
		return
	}

	for i := startIndex; i < len(candidates) && sum+candidates[i] <= target; i++ {

		// used[i-1] == true 说明同一树枝candidates[i-1] 使用过
		// used[i-1] == flase 说明同一树层 使用过
		if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == false {
			continue
		}
		path = append(path, candidates[i])
		sum += candidates[i]
		used[i] = true
		backtracking(candidates, path, target, sum, i+1, used, res)
		path = path[:len(path)-1]
		sum -= candidates[i]
		used[i] = false
	}
}
