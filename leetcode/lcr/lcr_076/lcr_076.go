package lcr_076

import "math/rand/v2"

// https://leetcode.cn/problems/xx4gT2/
func findKthLargest(nums []int, k int) int {
	// param check
	if k <= 0 || k > len(nums) {
		return 0
	}
	return quickFind(nums, k)
}

func quickFind(nums []int, k int) int {
	idx := rand.IntN(len(nums))
	// swap nums idx & len(nums)-1
	nums[idx], nums[len(nums)-1] = nums[idx], nums[len(nums)-1]
	var cnt int
	pivotIdx := len(nums) - 1
	for i := 0; i < pivotIdx; i++ {
		if nums[i] > nums[pivotIdx] {
			nums[i], nums[cnt] = nums[cnt], nums[i]
			cnt++
		}
	}
	nums[cnt], nums[pivotIdx] = nums[pivotIdx], nums[cnt]
	if cnt == k-1 {
		return nums[cnt]
	} else if cnt > k-1 {
		return quickFind(nums[:cnt], k)
	} else {
		return quickFind(nums[cnt+1:], k-cnt-1)
	}
}
