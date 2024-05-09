package leetcode_2856

// https://leetcode.cn/problems/minimum-array-length-after-pair-removals/
func minLengthAfterRemovals(nums []int) int {
	// 1. 记录频次
	// 目的是找到众数, 寻找众数的方式:
	// 1. map
	// 2. 摩尔投票
	// 3. 排序后 寻找 nums[n/2] 位置的数
	var (
		n      = len(nums)
		cnts   = make(map[int]int)
		maxCnt int
	)
	for _, num := range nums {
		cnts[num]++
	}
	// find max freqNum
	for _, freq := range cnts {
		maxCnt = max(maxCnt, freq)
	}
	if maxCnt > n/2 {
		// 能消去的数量为  n- maxCnt
		return maxCnt - (n - maxCnt)
	}
	// 剩下的都可以消掉.
	return n % 2
}

func minLengthAfterRemovals2(nums []int) int {
	// 如果 存在众数， 一定是  nums[len(nums)/2]
	var (
		n         = len(nums)
		candidate = nums[n/2]
	)
	left := lowerBound(nums, 0, n-1, candidate)
	right := lowerBound(nums, 0, n-1, candidate+1) - 1
	maxCnt := right - left + 1
	if maxCnt > n/2 {
		return maxCnt - (n - maxCnt)
	}
	return n % 2
}

// 开区间
func lowerBound(nums []int, left int, right int, target int) int {
	left, right = left-1, right+1
	for left+1 < right {
		mid := left + (right-left)>>1
		if target <= nums[mid] {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}
