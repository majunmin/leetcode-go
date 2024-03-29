package leetcode_0189

// https://leetcode.cn/problems/rotate-array/?envType=study-plan-v2&envId=top-100-liked
func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums, 0, k)
	reverse(nums, k+1, len(nums)-1)
	reverse(nums, 0, len(nums)-1)
}

func reverse(nums []int, l int, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
