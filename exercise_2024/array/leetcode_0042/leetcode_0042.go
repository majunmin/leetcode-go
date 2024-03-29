package leetcode_0042

func trap(height []int) int {

	return stackSolution(height)
}

// 单调递减栈解法
// 在每次出战的时候 进行计算水量
func stackSolution(height []int) int {
	stack := make([]int, 0)
	var result int
	for i := 0; i < len(height); i++ {
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			if height[i] < height[top] {
				break
			}
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			topIdx := stack[len(stack)-1]
			result += (i - topIdx - 1) * (min(height[i], height[topIdx]) - height[top])
		}
		stack = append(stack, i)
	}
	return result
}

// 动态规划解法
// 当前柱子 可以接水量 取决于其左右两侧的最高的柱子有多高
func dpSolution(height []int) int {
	// param check
	if len(height) == 0 {
		return 0
	}
	// process
	leftMax, rightMax := make([]int, len(height)), make([]int, len(height))
	lm, rm := 0, 0
	for i := 0; i < len(height); i++ {
		lm, rm = max(lm, height[i]), max(rm, height[len(height)-1-i])
		leftMax[i], rightMax[len(height)-i-1] = lm, rm
	}
	//
	var result int
	for i := 0; i < len(height); i++ {
		result += min(leftMax[i], rightMax[i]) - height[i]
	}
	return result
}
