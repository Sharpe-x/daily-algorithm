package main

import "strconv"

func main() {

}

func evalRPN(tokens []string) int {

	var stack []int
	for _, v := range tokens {
		num := 0
		switch v {
		case "+":
			num = stack[len(stack)-2] + stack[len(stack)-1]
			stack = stack[:len(stack)-2]
		case "-":
			num = stack[len(stack)-2] - stack[len(stack)-1]
			stack = stack[:len(stack)-2]
		case "*":
			num = stack[len(stack)-2] * stack[len(stack)-1]
			stack = stack[:len(stack)-2]
		case "/":
			num = stack[len(stack)-2] / stack[len(stack)-1]
			stack = stack[:len(stack)-2]
		default:
			num, _ = strconv.Atoi(v)
		}
		stack = append(stack, num)
	}
	return stack[len(stack)-1]
}
