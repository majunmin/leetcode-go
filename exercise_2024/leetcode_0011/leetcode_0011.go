package leetcode_0011

// https://leetcode.cn/problems/container-with-most-water/?envType=study-plan-v2&envId=top-interview-150
func maxArea(height []int) int {
	// param check
	if len(height) < 2 {
		return 0
	}
	var (
		result      int
		n           = len(height)
		left, right = 0, n - 1
	)
	for left < right {
		width := right - left
		h := min(height[left], height[right])
		result = max(result, h*width)
		// 指针移动
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return result
}
