package _024_04_07

func longestMonotonicSubarray(nums []int) int {
	// param check
	if len(nums) < 2 {
		return len(nums)
	}
	var (
		maxLen  int
		incrCnt = 1
		decrCnt = 1
	)
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			incrCnt++
			decrCnt = 1
		} else if nums[i] < nums[i-1] {
			decrCnt++
			incrCnt = 1
		} else {
			incrCnt = 1
			decrCnt = 1
		}
		maxLen = max(maxLen, incrCnt, decrCnt)
	}
	return maxLen
}
