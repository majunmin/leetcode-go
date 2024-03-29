package leetcode_0055

// https://leetcode.cn/problems/jump-game/
func canJump(nums []int) bool {
	// param check
	if len(nums) == 0 {
		panic("invalid param")
	}
	if len(nums) == 1 {
		return true
	}
	//
	var maxRight int
	for i := 0; i <= maxRight; i++ {
		maxRight = max(maxRight, i+nums[i])
		if maxRight >= len(nums)-1 {
			return true
		}
	}
	return false
}
