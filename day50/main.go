package main

import "strconv"

func main() {

}

func restoreIpAddresses(s string) []string {
	var res, path []string
	backtracking(s, 0, path, &res)
	return res
}

func backtracking(s string, startIndex int, path []string, res *[]string) {
	if startIndex == len(s) && len(path) == 4 {
		tmpIpString := path[0] + "." + path[1] + "." + path[2] + "." + path[3]
		*res = append(*res, tmpIpString)
	}

	for i := startIndex; i < len(s); i++ {
		path := append(path, s[startIndex:i+1])
		if i-startIndex+1 <= 3 && len(path) <= 4 && isNormalIp(s, startIndex, i) {
			backtracking(s, i+1, path, res)
		} else { // 首尾超过了3, 或者路径多于4个，或者大于255 直接回退
			return
		}
		path = path[:len(path)-1]
	}
}

func isNormalIp(s string, startIndex, end int) bool {
	checkInt, _ := strconv.Atoi(s[startIndex : end+1])
	if end-startIndex+1 > 1 && s[startIndex] == '0' { // 前导0
		return false
	}
	if checkInt > 255 {
		return false
	}
	return true
}
