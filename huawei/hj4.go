package main

import (
	"fmt"
)

func main() {
	var str string
	_, _ = fmt.Scan(&str)
	splitStr(str)
}

func splitStr(str string) {

	strLen := len(str)
	if strLen == 0 {
		return
	}

	n := strLen / 8
	m := strLen % 8

	if m != 0 {
		for i := 0; i < 8-m; i++ {
			str += "0"
		}
		n++
	}

	s := 0
	for n > 0 {
		word := str[s*8 : (s*8 + 8)]
		s++
		n--
		fmt.Println(word)
	}
}
