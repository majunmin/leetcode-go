package leetcode_3038

// https://leetcode.cn/problems/maximum-number-of-operations-with-the-same-score-i/?envType=daily-question&envId=2024-06-07
func maxOperations(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	var (
		target = nums[0] + nums[1]
		cnt    = 1
	)
	for i := 3; i < len(nums); i += 2 {
		if nums[i]+nums[i-1] == target {
			cnt++
		}
	}
	return cnt
}
