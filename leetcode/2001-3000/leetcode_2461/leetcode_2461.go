package leetcode_2461

func maximumSubarraySum(nums []int, k int) int64 {
	var (
		n         = len(nums)
		result    int
		kind      int //  窗口中数字的种类数
		windowSum int // 窗口数字求和
		cnts      = make(map[int]int, n)
	)
	for r := 0; r < n; r++ {
		cnts[nums[r]]++
		if cnts[nums[r]] == 1 {
			kind++
		}
		windowSum += nums[r]
		if r < k-1 {
			continue
		}
		if kind == k {
			result = max(result, windowSum)
		}
		// r >= k-1
		cnts[nums[r-k+1]]--
		windowSum -= nums[r-k+1]
		if cnts[nums[r-k+1]] == 0 {
			kind--
		}
	}
	return int64(result)
}
