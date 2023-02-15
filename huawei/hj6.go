package main

import "fmt"

func main() {
	var num int
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		panic(err)
	}

	if num < 2 {
		fmt.Println(num)
		return
	}

	for i := 2; i*i <= num; {
		if num%i == 0 {
			num /= i
			fmt.Printf("%d ", i)
		} else {
			i++
		}
	}

	if num > 1 {
		fmt.Println(num)
	}
}
