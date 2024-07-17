package lcr_038

// https://leetcode.cn/problems/iIQa4I/
func dailyTemperatures(temperatures []int) []int {
	var (
		n      = len(temperatures)
		result = make([]int, n)
		stack  = make([]int, 0, n)
	)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			// pop stack
			stack = stack[:len(stack)-1]
		}
		// add to result
		if len(stack) == 0 {
			result[i] = 0
		} else {
			result[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	return result
}
