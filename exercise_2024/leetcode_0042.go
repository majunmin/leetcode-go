package exercise_2024

func trap(height []int) int {
	// 单调递减栈
	stack := make([]int, 0, len(height))
	var result int
	for i := 0; i < len(height); i++ {
		for len(stack) > 0 && height[i] >= height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			// pop
			stack = stack[:len(stack)-1]
			//calculate
			if len(stack) > 0 {
				wid := i - stack[len(stack)-1] - 1
				hei := min(height[i], height[stack[len(stack)-1]]) - height[top]
				result += wid * hei
			}
		}
		stack = append(stack, i)
	}
	return result
}
