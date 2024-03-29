package leetcode_0084

type pair struct {
	idx int
	h   int
}

// https://leetcode.cn/problems/largest-rectangle-in-histogram/?envType=study-plan-v2&envId=top-100-liked
func largestRectangleArea(heights []int) int {
	// 单调递增栈
	// 使用哨兵模式  优化代码?
	var (
		result int
		stack  = make([]int, 0)
		i      int
	)
	for i < len(heights) {
		if len(stack) == 0 || heights[i] > heights[stack[len(stack)-1]] {
			stack = append(stack, i)
			i++
			continue
		}
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			// pop
			item := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := -1
			if len(stack) > 0 {
				left = stack[len(stack)-1]
			}
			result = max(result, heights[item]*(i-left-1))
		}
		stack = append(stack, i)
		i++
	}
	// 计算 栈中剩余元素的结果
	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		left := -1
		if len(stack) > 0 {
			left = stack[len(stack)-1]
		}
		result = max(result, heights[item]*(i-left-1))
	}
	return result
}
