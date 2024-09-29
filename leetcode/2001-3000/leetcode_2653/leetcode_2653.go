package leetcode_2653

var (
	bound = 50
)

// 1. 数据范围较小的场景下， 可以使用暴力方法解决
// https://leetcode.cn/problems/sliding-subarray-beauty/
func getSubarrayBeauty(nums []int, k int, x int) []int {
	var (
		cnts   = make([]int, bound*2+1)
		n      = len(nums)
		result = make([]int, n-k+1)
	)
	for i := 0; i < k; i++ {
		cnts[nums[i]+bound]++
	}
	// find 第x小的数字
	item := findKthMin(cnts, x)
	if item < 0 {
		result[0] = item
	}
	for i := k; i < n; i++ {
		cnts[nums[i]+bound]++
		cnts[nums[i-k]+bound]--
		item = findKthMin(cnts, x)
		if item < 0 {
			result[i-k+1] = item
		}
	}
	return result
}

func findKthMin(cnts []int, k int) int {
	for i := 0; i < len(cnts); i++ {
		if cnts[i] == 0 {
			continue
		}
		if cnts[i] >= k {
			return i - bound
		}
		k -= cnts[i]
	}
	return 0
}
