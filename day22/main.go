package main

import "fmt"

func main() {
	//fmt.Println(getNext("aabaaf"))
	fmt.Println(strStr("aabaabaaf", "aabaaf"))
}

func getNext(s string) []int {
	newSlice := []rune(s)
	next := make([]int, len(newSlice))

	j := 0
	for i := 1; i < len(newSlice); i++ {

		for j > 0 && s[i] != s[j] { // j 要保证大于0 因为下面取得是j-1
			j = next[j-1] //前一个字符串对应的回退位置
		}

		if s[i] == s[j] {
			j++
		}

		next[i] = j
	}
	return next
}

func strStr(haystack string, needle string) int {

	h, n := []rune(haystack), []rune(needle)
	next := getNext(needle)

	j := 0
	for i := 0; i < len(h); i++ {

		for j > 0 && h[i] != n[j] {
			j = next[j-1]
		}

		if h[i] == n[j] {
			j++
		}

		if j == len(n) {
			return i - len(n) + 1
		}
	}

	return -1
}
