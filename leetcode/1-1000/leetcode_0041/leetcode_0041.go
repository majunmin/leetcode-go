package leetcode_0041

// https://leetcode.cn/problems/first-missing-positive/?envType=study-plan-v2&envId=top-100-liked
// 题解: https://leetcode.cn/problems/first-missing-positive/solutions/7703/tong-pai-xu-python-dai-ma-by-liweiwei1419/?envType=study-plan-v2&envId=top-100-liked
func firstMissingPositive(nums []int) int {
	// param check
	n := len(nums)
	for i := 0; i < len(nums); i++ {
		//nums[nums[i]-1] != nums[i] 作用?
		//将数组下标视为 hash表的key
		// 为什么不是 nums[nums[i]] != nums[i]
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
