package leetcode_1793

// https://leetcode.cn/problems/maximum-score-of-a-good-subarray/description/?envType=daily-question&envId=2024-03-19
func maximumScore(nums []int, k int) int {
	// param check
	if k < 0 || k >= len(nums) {
		// todo:invalid param
		return -1
	}
	l, r := k-1, k+1
	result := nums[k]
	cur := nums[k]
	for l >= 0 || r < len(nums) {
		if l < 0 || (r < len(nums) && nums[r] > nums[l]) {
			cur = min(cur, nums[r])
			r++
		} else {
			cur = min(cur, nums[l])
			l--
		}
		result = max(result, cur*(r-l-1))
	}
	return result
}

func solution1(nums []int, k int) int {
	// param check
	if k < 0 || k >= len(nums) {
		// todo:invalid param
		return -1
	}
	l, r := k, k
	var result int
	for {
		if l < 0 && r == len(nums) {
			break
		}
		for l >= 0 && nums[l] >= nums[k] {
			l--
		}
		for r < len(nums) && nums[r] >= nums[k] {
			r++
		}
		result = max(result, (r-l-1)*nums[k])
		if l < 0 {
			nums[k] = nums[r]
		} else if r >= len(nums) {
			nums[k] = nums[l]
		} else {
			nums[k] = max(nums[l], nums[r])
		}
	}
}
