package leetcode_2958

// https://leetcode.cn/problems/length-of-longest-subarray-with-at-most-k-frequency/description/
func maxSubarrayLength(nums []int, k int) int {
	cnts := make(map[int]int) // 窗口内数字出现的次数
	var l int
	var result int
	for r := 0; r < len(nums); r++ {
		cnts[nums[r]]++
		for cnts[nums[r]] > k {
			// shrink window
			cnts[nums[l]]--
			l++
		}
		result = max(result, r-l+1)
	}
	return result
}

func templateSolution(nums []int, k int) int {
	// param check
	if k == 0 {
		return 0
	}
	if len(nums) < k {
		return len(nums)
	}
	// 滑动窗口解法
	var (
		l, r     int
		numsCnts = make(map[int]int) // 窗口内 数字的个数
		result   int
	)
	for r < len(nums) {
		numsCnts[nums[r]]++
		if numsCnts[nums[r]] <= k {
			r++
			result = max(result, r-l)
			continue
		}
		// 窗口内 nums[r] 出现的次数  > k
		// shrink window
		for l < r {
			numsCnts[nums[l]]--
			l++
			if numsCnts[nums[l-1]] == k {
				break
			}
		}
		r++
	}
	return result
}
