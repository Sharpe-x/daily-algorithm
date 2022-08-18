// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

// 注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。

package main

func main() {

}

func isAnagramBySlice(s string, t string) bool {
	var nums [26]int
	for _, v := range s {
		nums[v-rune('a')] += 1
	}

	for _, v := range t {
		nums[v-rune('a')] -= 1
	}

	// for _, v := range nums {
	// 	if v != 0 {
	// 		return false
	// 	}
	// }

	return nums == [26]int{}
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	exists := make(map[byte]int)
	for _, v := range s {
		if _, ok := exists[byte(v)]; ok {
			exists[byte(v)]++
		} else {
			exists[byte(v)] = 1
		}
	}

	for _, v := range t {
		if n, ok := exists[byte(v)]; ok && n > 0 {
			exists[byte(v)]--
		} else {
			return false
		}
	}

	return true
}
