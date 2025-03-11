package leetcode_2012

// https://leetcode.cn/problems/sum-of-beauty-in-the-array/?envType=daily-question&envId=2025-03-11
func sumOfBeauties(nums []int) int {
	//题目还得好好读一遍 :cry

	var (
		n = len(nums)
		// 维护一个 单调非递减栈
		leftStack = make([]int, n)
		// 维护一个 单调非递增栈
		rightStack = make([]int, n)
		result     int
	)
	// init stack
	leftStack[0] = nums[0]
	rightStack[n-1] = nums[n-1]
	for i := 1; i < n; i++ {
		leftStack[i] = max(leftStack[i-1], nums[i])
		rightStack[n-i-1] = min(rightStack[n-i], nums[n-i-1])
	}
	for i := 1; i < n-1; i++ {
		if leftStack[i-1] < nums[i] && nums[i] < rightStack[i+1] {
			result += 2
		} else if nums[i-1] < nums[i] && nums[i] < nums[i+1] {
			result += 1
		}
	}
	return result
}
