package main

var result [][]int

func combinationSum3(k int, n int) [][]int {
	result = [][]int{}
	backtracking(k, n, 0, 1, []int{})
	return result
}

func backtracking(k, n, sum, start int, path []int) {

	if sum > n {
		return
	}
	if len(path) == k {
		if sum == n {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
		}
	}

	for i := start; i <= 9; i++ {
		sum += i
		path = append(path, i)
		backtracking(k, n, sum, i+1, path)
		sum -= i
		path = path[:len(path)-1]
	}
}
