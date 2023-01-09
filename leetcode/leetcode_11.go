package leetcode

// 1,8,6,2,5,4,8,3,7
// https://leetcode.cn/problems/container-with-most-water/
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}

	left, right := 0, len(height)-1
	maxValue := 0
	for left < right {
		if height[left] <= height[right] {
			maxValue = maxInt(maxValue, (right-left)*height[left])
			left++
			for left < right && height[left] < height[left-1] {
				left++
			}
		} else if height[left] > height[right] {
			maxValue = maxInt(maxValue, (right-left)*height[right])
			right--
			for left < right && height[right] < height[right+1] {
				right--
			}
		}
	}
	return maxValue
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
