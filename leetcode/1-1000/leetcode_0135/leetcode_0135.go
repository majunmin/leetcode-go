package leetcode_0135

// https://leetcode.cn/problems/candy/description/?envType=study-plan-v2&envId=top-interview-150
func candy(ratings []int) int {
	// 按照两种规则发糖.
	// 1. 左规则:
	// 从左到右遍历 ratings, 如果nums[i] > nums[i-1], nums[i] 就需要比 nums[i-1]多发一颗糖.
	// 2. 右规则
	// 从右到左遍历 ratings, 如果nums[i] > nums[i+1], nums[i] 就需要比 nums[i+1]多发一颗糖.
	var (
		n      = len(ratings)
		left   = make([]int, n)
		right  = make([]int, n)
		result int
	)
	// init state
	for i := 0; i < n; i++ {
		left[i] = 1
		right[i] = 1
	}
	// 按规则遍历.
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		}
		if ratings[n-i-1] > ratings[n-i] {
			right[n-i-1] = right[n-i] + 1
		}
	}

	// 统计结果.
	for i := 0; i < n; i++ {
		result += max(left[i], right[i])
	}
	return result
}
