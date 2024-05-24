package leetcode_0150

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	// check invalid tokens
	var (
		numStack = make([]int, 0)
	)
	for _, item := range tokens {

		switch item {
		case "+", "-", "*", "/":
			num1 := numStack[len(numStack)-1]
			num2 := numStack[len(numStack)-1]
			numStack = append(numStack, eval(num1, num2, item))
		default:
			numStack = append(numStack, strToInt(item))
		}
	}
	return numStack[0]
}

func eval(n1, n2 int, item string) int {
	var res int
	switch item {
	case "+":
		res = n1 + n2
	case "-":
		res = n1 - n2
	case "*":
		res = n1 * n2
	case "/":
		res = n1 / n2
	}
	return res
}

func strToInt(val string) int {
	num, _ := strconv.Atoi(val)
	return num
}
