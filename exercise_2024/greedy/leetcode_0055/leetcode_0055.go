package leetcode_0055

// https://leetcode.cn/problems/jump-game/?envType=study-plan-v2&envId=top-100-liked
func canJump(nums []int) bool {
	// param check
	n := len(nums)
	if n == 1 {
		return true
	}
	// len(nums) > 1
	end := nums[0]
	for i := 0; i < end; i++ {
		end = max(end, i+nums[i])
		if end >= n-1 {
			return true
		}
	}
	return false
}
