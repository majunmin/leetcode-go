package leetcode_100344

// https://leetcode.cn/contest/biweekly-contest-133/problems/minimum-operations-to-make-binary-array-elements-equal-to-one-i/
func minOperations(nums []int) int {
	// param check
	var (
		n      = len(nums)
		result int
	)
	for i := 0; i < n-2; i++ {
		if nums[i] == 1 {
			continue
		}
		// revert nums[i:i+3]
		for j := i; j < i+3; j++ {
			nums[j] = 1 - nums[j]
		}
		result++
	}
	if nums[n-2] == 1 && nums[n-1] == 1 {
		return result
	}
	return -1
}
