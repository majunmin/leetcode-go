package leetcode_0042
 visua

// 双指针解法
func doublePointer(height []int) int {
	// param check
	if len(height) < 3 {
		return 0
	}

	// process

	left, right := 0, len(height)-1
	var leftMax, rightMax int
	res := 0
	for left < right {
		leftMax = maxInt(leftMax, height[left])
		rightMax = maxInt(rightMax, height[right])
		if leftMax < rightMax {
			res += leftMax - height[left]
			left++
		} else {
			res += rightMax - height[right]
			right--
		}
	}

	return res
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b

}
