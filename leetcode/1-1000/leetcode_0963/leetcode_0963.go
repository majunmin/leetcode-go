package leetcode_0963

// https://leetcode.cn/problems/maximum-average-subarray-i/
func findMaxAverage(nums []int, k int) float64 {
	var total int
	for i := 0; i < k; i++ {
		total += nums[i]
	}
	avg := float64(total) / float64(k)
	for i := k; i < len(nums); i++ {
		total += nums[i]
		total -= nums[i-k]
		avg = min(avg, float64(total)/float64(k))
	}
	return avg
}
