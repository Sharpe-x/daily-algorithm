package main

func main() {
}

var ans [][]int

func combine(n int, k int) [][]int {
	ans = [][]int{}
	if n <= 0 || k <= 0 || k > n {
		return ans
	}
	backtracking(n, k, 1, []int{})
	return ans
}

func backtracking(n, k, startIndex int, track []int) {
	if len(track) == k {
		temp := make([]int, k)
		copy(temp, track)
		ans = append(ans, temp)
	}

	if len(track)+n-startIndex+1 < k {
		return
	}

	for i := startIndex; i <= n; i++ {
		track = append(track, i)
		backtracking(n, k, i+1, track)
		track = track[:len(track)-1]
	}
}
