package leetcode_0862

// https://leetcode.cn/problems/shortest-subarray-with-sum-at-least-k/
// 前缀和 + 单调队列
// 符合条件的场景:  nums[i] - nums[j] >= k && i > j
func shortestSubarray(nums []int, k int) int {
	// cal pre sum
	n := len(nums)
	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		s[i] = s[i-1] + nums[i-1]
	}

	result := n + 1
	queue := make([]int, n) // 单调递增队列， 里面存储的是 index
	for i, item := range s {
		for len(queue) > 0 && item-s[queue[0]] >= k { // 相当于缩小窗口,找到更好的结果
			result = minInt(result, i-queue[0])
			queue = queue[1:]
		}
		for len(queue) > 0 && item <= s[queue[len(queue)-1]] { // curItem <= preItem,  pop PreItem
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
	}
	// 没找到结果
	if result > n {
		return -1
	}
	return result
}

func minInt(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
