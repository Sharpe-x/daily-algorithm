package main

import (
	"fmt"
)

func main() {
	a := "我爱你中国"
	fmt.Println(reverseStringByRune(a))
	bytes := []byte(a)
	reverseString(bytes)
	fmt.Println(string(bytes))
}

func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func reverseStr(s string, k int) string {

	ss := []byte(s)

	for i := 0; i < len(ss); i = i + 2*k {
		if i+k <= len(ss) {
			reverseString(ss[i : i+k])
		} else {
			reverseString(ss[i:])
		}
	}

	return string(ss)
}

// 支持中文,英文
func reverseStringByRune(s string) string {
	ss := []rune(s)
	l, r := 0, len(ss)-1
	for l < r {
		ss[l], ss[r] = ss[r], ss[l]
		l++
		r--
	}
	return string(ss)
}
