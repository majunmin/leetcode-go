package _024_03_30

// https://leetcode.cn/contest/biweekly-contest-127/problems/shortest-subarray-with-or-at-least-k-i/
func minimumSubarrayLength(nums []int, k int) int {
	// param check
	if len(nums) == 0 {
		return -1
	}
	//  枚举 长度
	for j := 1; j < len(nums); j++ {
		for i := 0; i <= len(nums)-j; i++ {
			var total int
			for m := i; m < i+j; m++ {
				total |= nums[m]
			}
			if total >= k {
				return j
			}
		}
	}
	return -1
}
