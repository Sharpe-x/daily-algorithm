package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)

	s := make(map[rune]struct{})
	for _, v := range str {
		s[v] = struct{}{}
	}
	fmt.Println(len(s))
}
