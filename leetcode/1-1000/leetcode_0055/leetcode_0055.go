package leetcode_0055

// https://leetcode.cn/problems/jump-game/
func canJump(nums []int) bool {
	rightMost := 0
	for i := 0; i < len(nums); i++ {
		if i <= rightMost {
			rightMost = maxInt(rightMost, i+nums[i])
			if rightMost >= len(nums)-1 {
				return true
			}
		} else {
			return false
		}

	}
	return false
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
