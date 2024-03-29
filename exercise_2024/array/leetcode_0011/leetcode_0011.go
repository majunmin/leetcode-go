package leetcode_0011

// https://leetcode.cn/problems/container-with-most-water/?envType=study-plan-v2&envId=top-100-liked
func maxArea(height []int) int {
	// param check
	if len(height) < 2 {
		return 0
	}
	left, right := 0, len(height)-1
	var result int
	for left < right {
		wid := right - left
		hei := min(height[left], height[right])
		// cal water =  width * height
		result = max(result, wid*hei)
		// 移动左右
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return result
}
