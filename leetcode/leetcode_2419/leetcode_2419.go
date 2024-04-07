package leetcode_2419

// https://leetcode.cn/problems/longest-subarray-with-maximum-bitwise-and/
func longestSubarray(nums []int) int {
	// 一次遍历的解法
	var (
		cnt    int
		maxVal int
		result int
	)
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxVal {
			maxVal, cnt, result = nums[i], 1, 1
		} else if nums[i] == maxVal {
			cnt++
			result = max(result, cnt)
		} else {
			cnt = 0
		}
	}
	return result
}

// 两次遍历的解法
func solution1(nums []int) int {
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
