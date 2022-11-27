package main

import "fmt"

func main() {

	fmt.Println(partition("aab"))
	fmt.Println(partition("cbbbcc"))

}

func partition(s string) [][]string {
	var path []string
	var ans [][]string
	backtracking(s, 0, path, &ans)
	return ans
}

func backtracking(s string, startIndex int, path []string, res *[][]string) {
	if startIndex == len(s) {
		t := make([]string, len(path))
		copy(t, path)
		*res = append(*res, t)
		return
	}

	for i := startIndex; i < len(s); i++ {
		str := s[startIndex : i+1]
		if isPalindrome(str) {
			path = append(path, str)
		} else {
			continue
		}
		backtracking(s, i+1, path, res)
		path = path[:len(path)-1]
	}
}

func isPalindrome(str string) bool {
	l, r := 0, len(str)-1
	for l < r {
		if str[l] != str[r] {
			return false
		}
		l++
		r--
	}
	return true
}
