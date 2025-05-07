package leetcode_3392

// https://leetcode.cn/problems/count-subarrays-of-length-three-with-a-condition/?envType=daily-question&envId=2025-04-27
func countSubarrays(nums []int) int {
	var result int
	for i := 1; i < len(nums)-1; i++ {
		if nums[i]&1 == 1 {
			continue
		}
		if nums[i-1]+nums[i+1] == nums[i]>>1 {
			result++
		}
	}
	return result
}
