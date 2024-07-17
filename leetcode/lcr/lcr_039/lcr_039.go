package lcr_039

// https://leetcode.cn/problems/0ynMMM/
func largestRectangleArea(heights []int) int {
	newHeights := make([]int, len(heights)+1)
	newHeights = append([]int{0}, heights...)
	newHeights = append(newHeights, 0)
	var (
		stack  = make([]int, 0, len(newHeights))
		result int
	)
	stack = append(stack, 0)
	for i := 1; i < len(newHeights)-1; i++ {
		for len(stack) > 0 && newHeights[i] < newHeights[stack[len(stack)-1]] {
			height := newHeights[stack[len(stack)-1]]
			// pop stack
			stack = stack[:len(stack)-1]
			width := i - stack[len(stack)-1] - 1
			result = max(result, height*width)
		}
		stack = append(stack, i)
	}
	return result
}
