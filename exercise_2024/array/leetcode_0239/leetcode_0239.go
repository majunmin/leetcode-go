package leetcode_0239

// https://leetcode.cn/problems/sliding-window-maximum/
func maxSlidingWindow(nums []int, k int) []int {
	return solution1(nums, k)
}

// 利用单调递减队列
func solution1(nums []int, k int) []int {
	n := len(nums)
	if n == 1 || k == 1 {
		return nums
	}

	// process
	result := make([]int, 0, n-k+1)
	// 单调非递增队列
	queue := make([]int, 0)
	for i := 0; i < k; i++ {
		enQueue(&queue, nums, i, k)
	}
	result = append(result, nums[queue[0]])
	for i := k; i < n; i++ {
		enQueue(&queue, nums, i, k)
		result = append(result, nums[queue[0]])
	}
	return result
}

func enQueue(queue *[]int, nums []int, i int, k int) {
	// 移动窗口
	if len(*queue) > 0 && (*queue)[0] <= i-k {
		*queue = (*queue)[1:]
	}
	for len(*queue) > 0 && nums[i] > nums[(*queue)[len(*queue)-1]] {
		*queue = (*queue)[:len(*queue)-1]
	}
	*queue = append(*queue, i)
}
