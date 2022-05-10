package leetcode_0084

func maxInt(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

//https://leetcode-cn.com/problems/largest-rectangle-in-histogram/
func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	if len(heights) == 1 {
		return heights[0]
	}

	newHeights := make([]int, len(heights)+2)
	for i, h := range heights {
		newHeights[i+1] = h
	}

	maxArea := 0
	stack := make([]int, 0, len(newHeights))
	stack = append(stack, 0)
	for i := 1; i < len(newHeights); i++ {
		for newHeights[i] < newHeights[stack[len(stack)-1]] {
			//出栈
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			height := newHeights[top]
			width := i - stack[len(stack)-1] - 1
			maxArea = maxInt(maxArea, height*width)

		}
		stack = append(stack, i)
	}
	return maxArea
}

func stackSolution1(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	if len(heights) == 1 {
		return heights[0]
	}
	maxArea := 0
	//单调栈解法

	stack := make([]int, 0, len(heights))
	for i, h := range heights {

		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			width := 0
			height := heights[top]
			if len(stack) == 0 {
				width = i
			} else {
				width = i - stack[len(stack)-1] - 1
			}
			// 出栈
			maxArea = maxInt(maxArea, height*width)

		}
		stack = append(stack, i)
	}
	// 遍历 剩余栈顶元素
	for len(stack) > 0 {
		//出栈
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		width := 0
		height := heights[top]
		if len(stack) == 0 {
			width = len(heights)
		} else {
			width = len(heights) - stack[len(stack)-1] - 1
		}
		maxArea = maxInt(maxArea, height*width)
	}

	return maxArea
}

// 暴力枚举, 枚举柱状图的高
// 时间复杂度: O(N^2)
// 空间复杂度: O(1)
func iterSolution(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	if len(heights) == 1 {
		return heights[0]
	}

	maxArea := 0
	// 枚举柱状图的高度
	for i := 0; i < len(heights); i++ {
		h := heights[i]
		left, right := i, i
		for left > 0 && heights[left-1] >= heights[i] {
			left--
		}
		for right < len(heights)-1 && heights[right+1] >= heights[i] {
			right++
		}
		maxArea = maxInt(maxArea, (right-left+1)*h)
	}
	return maxArea
}
