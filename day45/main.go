package main

import "strconv"

func main() {

}

var num2char = map[int][]byte{
	0: []byte(""),
	1: []byte(""),
	2: []byte("abc"),
	3: []byte("def"),
	4: []byte("ghi"),
	5: []byte("jkl"),
	6: []byte("mno"),
	7: []byte("pqrs"),
	8: []byte("tuv"),
	9: []byte("wxyz"),
}

var result []string

func letterCombinations(digits string) []string {
	result = []string{}
	if digits == "" {
		return result
	}
	backtracking(digits, "", 0)
	return result
}

func backtracking(digits, str string, index int) {

	if len(str) == len(digits) {
		result = append(result, str)
		return
	}

	num, _ := strconv.Atoi(string(digits[index]))
	baseStr := num2char[num]
	tmp := []byte(str)
	for i := 0; i < len(baseStr); i++ {
		tmp = append(tmp, baseStr[i])
		backtracking(digits, string(tmp), index+1)
		tmp = tmp[:len(tmp)-1]
	}
}

var digitsMap = [10]string{
	"",     // 0
	"",     // 1
	"abc",  // 2
	"def",  // 3
	"ghi",  // 4
	"jkl",  // 5
	"mno",  // 6
	"pqrs", // 7
	"tuv",  // 8
	"wxyz", // 9
}

func test(digits string, n int) {
	letter := digitsMap[digits[n]-'0']
}
