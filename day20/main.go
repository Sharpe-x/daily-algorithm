package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello  world   a b "
	fmt.Println(reverseWords(s))
	//fmt.Println(string(removeSpace([]rune(s))))
}

func replaceSpaceWitKit(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

func replaceSpace(s string) string {

	var newStr []rune
	for _, v := range s {
		if v == ' ' {
			newStr = append(newStr, []rune("%20")...)
		} else {
			newStr = append(newStr, v)
		}
	}
	return string(newStr)
}

func replaceSpaceInSuti(s string) string {
	bytes := []byte(s)
	var spaceNum int
	orLen := len(bytes)
	for i := 0; i < orLen; i++ {
		if bytes[i] == ' ' {
			spaceNum++
		}
	}

	tmp := make([]byte, spaceNum*2)
	bytes = append(bytes, tmp...)
	j := len(bytes) - 1
	i := orLen - 1
	for i >= 0 {
		if bytes[i] == ' ' {
			bytes[j] = '0'
			bytes[j-1] = '2'
			bytes[j-2] = '%'
			j = j - 3
		} else {
			bytes[j] = bytes[i]
			j--
		}
		i--
	}

	return string(bytes)
}

//https://leetcode.cn/problems/reverse-words-in-a-string/

func reverseWords(s string) string {
	strs := removeSpace([]rune(s))
	reverse(strs)
	l, r := 0, 0
	for ; r < len(strs); r++ {
		if strs[r] == ' ' {
			reverse(strs[l:r])
			l = r + 1
		}
	}
	reverse(strs[l:])
	return string(strs)
}

func removeSpace(strs []rune) []rune {
	slow := 0
	for fast := 0; fast < len(strs); fast++ {
		if strs[fast] != ' ' {

			if slow != 0 {
				strs[slow] = ' '
				slow++
			}

			for fast < len(strs) && strs[fast] != ' ' {
				strs[slow] = strs[fast]
				fast++
				slow++
			}
		}
	}

	return strs[:slow]
}

func reverse(s []rune) {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
