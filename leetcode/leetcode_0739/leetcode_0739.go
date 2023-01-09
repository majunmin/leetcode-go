package leetcode_0739

// https://leetcode.cn/problems/daily-temperatures/
// 用一个单点递减栈, 栈中存储   index 数据, 从后往前遍历
func dailyTemperatures(temperatures []int) []int {
	// param check
	n := len(temperatures)
	if n == 0 {
		return nil
	}

	//
	stack := make([]int, 0, n)
	result := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && temperatures[i] >= temperatures[stack[len(stack)-1]] {
			// pop 出栈, 维护单调递递减栈
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 { //如果不存在 比当前温度大的, 填充 0
			result[i] = 0
		} else {
			result[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	return result
}
