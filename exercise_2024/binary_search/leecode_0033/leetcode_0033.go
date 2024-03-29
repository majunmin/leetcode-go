package leecode_0033

// https://leetcode.cn/problems/search-in-rotated-sorted-array/?envType=study-plan-v2&envId=top-100-liked
func search(nums []int, target int) int {
	// param check

	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)>>1
		// mid 在左侧区间里
		if nums[mid] <= nums[r] {
			// nums[mid:r] 有序
			if nums[mid+1] <= target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid
			}
		} else {
			// nums[l:mid] 是有序的
			if target >= nums[l] && target <= nums[mid] {
				r = mid
			} else {
				l = mid + 1
			}
		}
	}
	if nums[l] == target {
		return l
	}
	return -1
}
