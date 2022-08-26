package main

import "fmt"

func main() {
	fmt.Println(reverseLeftWords("helloworld", 2))
}

func reverseLeftWords1(s string, n int) string {

	bytes := []byte(s)
	var newBytes []byte
	newBytes = append(newBytes, bytes[n:]...)
	newBytes = append(newBytes, bytes[:n]...)

	return string(newBytes)
}

func reverse(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

func reverseLeftWords(s string, n int) string {

	bytes := []byte(s)
	reverse(bytes[:n])
	reverse(bytes[n:])
	reverse(bytes)

	return string(bytes)
}
