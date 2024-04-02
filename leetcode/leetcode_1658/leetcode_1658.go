package leetcode_1658

// https://leetcode.cn/problems/minimum-operations-to-reduce-x-to-zero/description/
func minOperations(nums []int, x int) int {
	// param check
	if len(nums) == 0 {
		return -1
	}
	//题目可以转化为  求解   nums 中和为 total(sum nums) - x 的最长子数组的长度 len
	// maxLen = len(nums) - len
	var (
		target = -x
		l, r   int
		maxLen = -1
	)
	for _, num := range nums {
		target += num
	}
	if target < 0 {
		return -1
	}

	for r < len(nums) {
		target -= nums[r]
		for l < r && target < 0 {
			target += nums[l]
			l++
		}
		// cal maxLen
		if target == 0 {
			maxLen = max(maxLen, r-l+1)
		}
		r++
	}
	if maxLen == -1 {
		return maxLen
	}
	return len(nums) - maxLen
}
