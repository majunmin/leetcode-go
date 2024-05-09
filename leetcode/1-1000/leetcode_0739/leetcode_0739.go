package leetcode_0739

// https://leetcode.cn/problems/daily-temperatures/
// 用一个单点递减栈, 栈中存储   index 数据, 从后往前遍历
func dailyTemperatures(temperatures []int) []int {
	return solution2(temperatures)
}

func solution2(temperatures []int) []int {
	if len(temperatures) == 0 {
		panic("len(temperatures) should greater than 0")
	}
	if len(temperatures) == 1 {
		return []int{0}
	}

	//
	result := make([]int, len(temperatures))
	// 单调递减栈
	stack := make([]int, 0)
	n := len(temperatures)
	// 0  1  2  3  4  5  6
	// 73 74 75 71 69 72 76
	for i := 0; i < n; i++ {
		// 有比当前栈顶元素大的数据, 依次出栈.
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[idx] = i - idx
		}
		stack = append(stack, i)
	}
	return result

}

// 从后往前j计算,效率快些
func solution1(temperatures []int) []int {
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
		if len(stack) > 0 { //如果不存在 比当前温度大的, 填充 0
			result[i] = stack[len(stack)-1] - i
		}
		// if len(stack) == 0， result[i] = 0, j即是默认值
		stack = append(stack, i)
	}
	return result
}
