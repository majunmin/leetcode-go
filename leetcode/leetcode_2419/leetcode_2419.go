package leetcode_2419

// https://leetcode.cn/problems/longest-subarray-with-maximum-bitwise-and/
func longestSubarray(nums []int) int {
	maxVal := max(nums[0], nums[1:]...)
	var result int
	var cnt int
	for _, num := range nums {
		if num == maxVal {
			cnt++
			result = max(result, cnt)
		} else {
			cnt = 0
		}
	}
	return result
}
