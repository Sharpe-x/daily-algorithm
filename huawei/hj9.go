package main

import "fmt"

func main() {

	var num int64
	_, err := fmt.Scan(&num)
	if err != nil {
		panic(err)
	}

	nums := make(map[int64]bool, 0)
	var target int64
	for num > 0 {
		tmp := num % 10
		if _, ok := nums[tmp]; !ok {
			target = tmp + target*10
			nums[tmp] = true
		}
		num = num / 10
	}

	fmt.Println(target)
}
