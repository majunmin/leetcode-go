package leetcode_0031

// https://leetcode.cn/problems/next-permutation/?envType=study-plan-v2&envId=top-100-liked
func nextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 {
		// 找到第一个 nums[i] < nums[i+1] 的位置(升序)
		if nums[i] < nums[i+1] {
			break
		}
		i--
	}
	if i > -1 {
		// 从后往前找,找到第一个 > nums[i] 的值
		for j := len(nums) - 1; j >= 0; j-- {
			if nums[j] > nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
				break
			}
		}
	}
	// reverse [i+1:]
	l, r := i+1, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
	}
}
