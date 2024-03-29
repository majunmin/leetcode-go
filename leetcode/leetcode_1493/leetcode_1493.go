package leetcode_1493

// https://leetcode.cn/problems/longest-subarray-of-1s-after-deleting-one-element/
func longestSubarray(nums []int) int {
	//枚举删除位置
	preSum := make([]int, len(nums))
	sufSum := make([]int, len(nums))
	preSum[0] = nums[0]
	sufSum[len(nums)-1] = nums[len(nums)-1]
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			preSum[i] = 0
		} else {
			preSum[i] = preSum[i-1] + 1
		}
		if nums[len(nums)-i-1] == 0 {
			sufSum[len(nums)-i-1] = 0
		} else {
			sufSum[len(nums)-i-1] = sufSum[len(nums)-i] + 1
		}
	}
	var result int
	for i := 0; i < len(nums); i++ {
		p, s := 0, 0
		if i == 0 {
			p = 0
		} else {
			p = preSum[i-1]
		}
		if i == len(nums)-1 {
			s = 0
		} else {
			s = sufSum[i+1]
		}
		result = max(result, p+s)
	}
	return result
}

// 滑动窗口解法
func slidingWindow(nums []int) int {
	var (
		zeroCnt int
		l, r    int
		result  int
	)
	for ; r < len(nums); r++ {
		if nums[r] == 0 {
			zeroCnt++
		}
		if nums[r] == 1 || nums[r] == 0 && zeroCnt == 1 {
			result = max(result, r-l)
			continue
		}
		// 窗口内 0 的个数 超过 1， shrink window
		for l < r {
			if nums[l] == 0 {
				l++
				break
			}
			l++
		}
	}
	return result
}
