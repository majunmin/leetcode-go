package leetcode

func rotate(nums []int, k int) {
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k)
	reverse(nums, k+1, len(nums)-1)
}

func reverse(nums []int, l, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
