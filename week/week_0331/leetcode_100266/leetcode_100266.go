package leetcode_100266

// https://leetcode.cn/contest/weekly-contest-391/problems/count-alternating-subarrays/
func countAlternatingSubarrays(nums []int) int64 {
	// 动态规划
	// param  check
	n := len(nums)
	if n < 2 {
		return int64(n)
	}
	var (
		total = 1
		add   = 1
	)
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			add++
			total += add
		} else {
			add = 1
			total += 1
		}
	}
	return int64(total)
}
