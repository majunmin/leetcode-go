package leetcode_0682

import "strconv"

// https://leetcode.cn/problems/baseball-game/
func calPoints(operations []string) int {
	var result int
	stack := make([]int, 0, len(operations))
	for _, str := range operations {
		if str == "C" {
			result -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		} else if str == "D" {
			stack = append(stack, int(stack[len(stack)-1]*2))
			result += stack[len(stack)-1]
		} else if str == "+" {
			stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
			result += stack[len(stack)-1]
		} else {
			item, _ := strconv.Atoi(str)
			stack = append(stack, item)
			result += stack[len(stack)-1]
		}
	}
	return result
}

func isDigit(x byte) bool {
	return x >= '0' && x <= '9'
}
