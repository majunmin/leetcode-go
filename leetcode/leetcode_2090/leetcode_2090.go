package leetcode_2090

// https://leetcode.cn/problems/k-radius-subarray-averages/
func getAverages(nums []int, k int) []int {
	// param check
	if len(nums) == 0 {
		return nil
	}
	var (
		result = make([]int, len(nums))
		sum    int
	)

	for i := 0; i < len(nums); i++ {
		if i < k || i >= len(nums)-k {
			result[i] = -1
			continue
		}
		sum += nums[i]
		if i >= 2*k {
			result[i-k] = sum / (2*k + 1)
			sum -= nums[i-2*k]
		}
	}
	return result
}
