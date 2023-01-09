package leetcode

import "math"

// https://leetcode.cn/problems/daily-temperatures/
func dailyTemperatures(temperatures []int) []int {
	if len(temperatures) == 0 {
		return nil
	}
	if len(temperatures) == 1 {
		return []int{0}
	}

	res := make([]int, len(temperatures))
	newTemperatures := make([]int, len(temperatures)+1)
	copy(newTemperatures[1:], temperatures)
	newTemperatures[0] = math.MaxInt

	// 存储下标， 用于计算  间隔
	stack := make([]int, 0)
	for i := 1; i < len(newTemperatures); i++ {
		t := newTemperatures[i]
		for len(stack) > 0 && t > newTemperatures[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[idx] = i - idx
		}
		//
		stack = append(stack, i)
	}
	return res
}
