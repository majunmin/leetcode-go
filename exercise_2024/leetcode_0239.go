package exercise_2024

// https://leetcode.cn/problems/sliding-window-maximum/?envType=study-plan-v2&envId=top-100-liked
func maxSlidingWindow(nums []int, k int) []int {
	// 单调递减队列解法
	// param check
	if len(nums) < k {
		return nil
	}
	queue := make([]int, 0, k)
	for i := 0; i < k; i++ {
		for len(queue) > 0 && nums[i] > queue[len(queue)-1] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
	}
	result := make([]int, 0, len(nums)-k+1)
	result = append(result, queue[0])
	for i := k; i < len(nums); i++ {
		if queue[0] < i-k+1 {
			queue = queue[1:]
		}
		for len(queue) > 0 && nums[i] > queue[len(queue)-1] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		result = append(result, nums[queue[0]])
	}
	return result
}
