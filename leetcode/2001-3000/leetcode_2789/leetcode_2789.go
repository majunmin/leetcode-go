package leetcode_2789

// https://leetcode.cn/problems/largest-element-in-an-array-after-merge-operations/
func maxArrayValue(nums []int) int64 {
	// param check
	if len(nums) == 0 {
		return -1
	}
	var maxVal int64
	latest := int64(nums[len(nums)-1])
	for i := len(nums) - 2; i >= 0; i-- {
		item := int64(nums[i])
		if latest < item {
			maxVal = max(maxVal, latest)
			latest = item
			continue
		}
		latest += item
	}
	return max(maxVal, latest)
}
