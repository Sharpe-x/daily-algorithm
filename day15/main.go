package main

import "fmt"

func main() {
	fmt.Println(isHappy(19))
}

func isHappy1(n int) bool {

	nums := make(map[int]struct{})

	for n > 0 {
		tmp := n
		sum := 0

		for tmp > 0 {
			a := tmp % 10
			tmp = tmp / 10
			sum += a * a
		}

		if sum == 1 {
			return true
		}

		if _, ok := nums[sum]; ok {
			return false
		}

		nums[sum] = struct{}{}
		n = sum
	}

	return true
}

func isHappy(n int) bool {

	nums := make(map[int]bool)

	for n != 1 && !nums[n] {
		n, nums[n] = getAns(n), true
	}

	return n == 1
}

func getAns(n int) int {
	sum := 0
	for n > 0 {
		tmp := n % 10
		sum += tmp * tmp
		n = n / 10
	}
	return sum
}
