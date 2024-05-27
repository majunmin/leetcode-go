package leetcode_0219

// https://leetcode.cn/problems/contains-duplicate-ii/?envType=study-plan-v2&envId=top-interview-150
func containsNearbyDuplicate(nums []int, k int) bool {
	var (
		num2LastIdx = make(map[int]int) //
	)
	for i := range nums {
		if j, exist := num2LastIdx[nums[i]]; exist && i-j <= k {
			return true
		}
		num2LastIdx[nums[i]] = i
	}
	return false
}

// 滑动窗口解法
func windowSolution(nums []int, k int) bool {
	var (
		window = make(map[int]int)
	)
	for i := 0; i <= min(k, len(nums)-1); i++ {
		window[nums[i]]++
		if window[nums[i]] > 1 {
			return true
		}
	}
	for i := k + 1; i < len(nums); i++ {
		window[nums[i]]++
		window[nums[i-k-1]]--
		if window[nums[i]] > 1 {
			return true
		}
	}
	return false
}
