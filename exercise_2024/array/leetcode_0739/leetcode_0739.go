package leetcode_0739

// https://leetcode.cn/problems/daily-temperatures/?envType=study-plan-v2&envId=top-100-liked
func dailyTemperatures(temperatures []int) []int {

	// 维护一个单调递减栈
	stack := make([]int, 0)
	result := make([]int, len(temperatures))
	stack = append(stack, len(temperatures)-1)
	for i := len(temperatures) - 2; i >= 0; i++ {
		for len(stack) > 0 && stack[len(stack)-1] < temperatures[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			result[i] = stack[len(stack)-1]
		}
		stack = append(stack, temperatures[i])
	}
	return result
}
