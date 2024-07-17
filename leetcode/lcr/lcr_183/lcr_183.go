package lcr_183

// 滑动窗口最大值
// https://leetcode.cn/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/
func maxAltitude(heights []int, limit int) []int {
	// 利用单调递减栈
	// param check
	if limit > len(heights) {
		return nil
	}
	if limit == 1 {
		return heights
	}
	var (
		stack  = make([]int, 0)
		result = make([]int, len(heights)-limit+1)
	)
	for i := 0; i < limit; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] <= heights[i] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	result = append(result, heights[stack[0]])
	for i := limit; i < len(heights); i++ {
		if stack[0] <= i-limit {
			stack = stack[1:]
		}
		for len(stack) > 0 && heights[stack[len(stack)-1]] <= heights[i] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
		result = append(result, heights[stack[0]])
	}
	return result
}
