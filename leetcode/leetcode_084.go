package leetcode

func largestRectangleArea(heights []int) int {

	if len(heights) == 0 {
		return 0
	}
	if len(heights) == 1 {
		return heights[0]
	}

	newHeights := make([]int, len(heights)+2)
	copy(newHeights[1:], heights)
	res := 0
	stack := make([]int, 1)
	for i, h := range newHeights {
		for len(stack) > 0 && h < newHeights[stack[len(stack)-1]] {
			// 找到右边界  i
			// pop
			curHeight := newHeights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			for curHeight == newHeights[stack[len(stack)-1]] {
				stack = stack[:len(stack)-1]
			}
			res = maxInt(res, curHeight*(i-stack[len(stack)-1]-1))
		}
		// push
		stack = append(stack, i)
	}
	return res
}
