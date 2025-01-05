package leetcode_3101

// https://leetcode.cn/problems/count-alternating-subarrays/
func countAlternatingSubarrays(nums []int) int64 {
	if len(nums) == 0 {
		return 0
	}
	nums = append([]int{nums[0]}, nums...)
	var (
		result int
		l      int

		r = 1
		n = len(nums)
	)
	for ; r < n; r++ {
		if nums[r] != nums[r-1] {
			continue
		}
		// n(n+1)/2
		size := r - l
		result += (size * (size + 1)) / 2
		l = r
	}
	size := r - l
	result += (size * (size + 1)) / 2
	return int64(result) - 1
}
