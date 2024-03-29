package leetcode_1695

// https://leetcode.cn/problems/maximum-erasure-value/
func maximumUniqueSubarray(nums []int) int {
	// param check
	if len(nums) == 0 {
		return 0
	}
	var (// https://leetcode.cn/problems/maximum-erasure-value/
	func maximumUniqueSubarray(nums []int) int {
		// param check
		if len(nums) == 0 {
		return 0
	}
		var (
		numsCnts = make(map[int]int) // 统计窗口内数字的个数
		result   int
		sum      int // 当前窗口内 数字的和
		l, r     int
	)
		for r < len(nums) {
		numsCnts[nums[r]]++
		sum += nums[r]
		// 窗口合法， 统计结果
		if numsCnts[nums[r]] == 1 {
		r++
		result = max(result, sum)
		continue
	}
		// 窗口出现重复数字
		// shrink window
		for l < r {
		numsCnts[nums[l]]--
		sum -= nums[l]
		l++
		if numsCnts[nums[l-1]] == 1 {
		break
	}
	}
	}
		return result
	}
		numsCnts = make(map[int]int) // 统计窗口内数字的个数
		result   int
		sum      int // 当前窗口内 数字的和
		l, r     int
	)
	for r < len(nums) {
		numsCnts[nums[r]]++
		sum += nums[r]
		// 窗口合法， 统计结果
		if numsCnts[nums[r]] == 1 {
			r++
			result = max(result, sum)
			continue
		}
		// 窗口出现重复数字
		// shrink window
		for l < r {
			numsCnts[nums[l]]--
			sum -= nums[l]
			l++
			if numsCnts[nums[l-1]] == 1 {
				break
			}
		}
	}
	return result
}
