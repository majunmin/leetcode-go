package leetcode_0189

// https://leetcode.cn/problems/rotate-array/?envType=study-plan-v2&envId=top-interview-150
func rotate(nums []int, k int) {
	// param check
	if len(nums) == 0 {
		return
	}
	n := len(nums)
	k = k % n
	reverse(nums, 0, n-k-1)
	reverse(nums, n-k, n-1)
	reverse(nums, 0, n-1)
}

func reverse(nums []int, left int, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
