package leetcode_0215

// solution2
// 超时啦...
func findKthLargest(nums []int, k int) int {
	// param check
	if k < 0 || k > len(nums) {
		return -1
	}
	pivot := nums[len(nums)-1]
	var cnt int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] <= pivot {
			nums[i], nums[cnt] = nums[cnt], nums[i]
			cnt++
		}
	}
	nums[cnt], nums[len(nums)-1] = nums[len(nums)-1], nums[cnt]
	if k == cnt+1 {
		return nums[cnt] // pivot
	}
	if k < cnt+1 {
		return findKthLargest(nums[:cnt], k)
	}
	// k > cnt + 1
	return findKthLargest(nums[cnt+1:], k-cnt-1)
}
