package leetcode_1004

// https://leetcode.cn/problems/max-consecutive-ones-iii/
func longestOnes(nums []int, k int) int {
	// param check
	if k >= len(nums) {
		return len(nums)
	}
	//
	var (
		cnt    int //窗口中 0 的个数
		result int
		l, r   int
	)
	for ; r < len(nums); r++ {
		if nums[r] == 0 {
			cnt++
		}
		// 当前窗口符合条件
		if cnt <= k {
			result = max(result, r-l+1)
			continue
		}
		// 不满足条件  shrink window
		for ; cnt > k && l < r; l++ {
			if nums[l] == 0 {
				cnt--
			}
		}
	}
	return result
}

func longestOnes2(nums []int, k int) int {
	// param check
	if k >= len(nums) {
		return len(nums)
	}
	//
	var (
		cnt    int //窗口中 0 的个数
		result int
		l, r   int
	)
	for r < len(nums) {
		if nums[r] == 0 {
			cnt++
		}
		// 当前窗口符合条件
		if cnt <= k {
			result = max(result, r-l+1)
			r++
			continue
		}
		// 不满足条件  shrink window
		for cnt > k && l < r {
			if nums[l] == 0 {
				cnt--
			}
			l++
		}
		r++
	}
	return result
}
