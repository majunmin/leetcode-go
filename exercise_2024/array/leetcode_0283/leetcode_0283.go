package leetcode_0283

// https://leetcode.cn/problems/move-zeroes/?envType=study-plan-v2&envId=top-100-liked
func moveZeroes(nums []int) {
	si := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			nums[i], nums[si] = nums[si], nums[i]
			si++
		}
	}
	for si < len(nums) {
		nums[si] = 0
		si++
	}
}
