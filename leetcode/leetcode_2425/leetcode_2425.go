package leetcode_2425

// https://leetcode.cn/problems/bitwise-xor-of-all-pairings/
func xorAllNums(nums1 []int, nums2 []int) int {
	var result int
	if len(nums1)&1 == 1 {
		for _, num := range nums2 {
			result ^= num
		}
	}
	if len(nums2)&1 == 1 {
		for _, num := range nums1 {
			result ^= num
		}
	}
	return result
}
