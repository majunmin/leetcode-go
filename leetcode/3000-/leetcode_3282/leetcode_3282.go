package leetcode_3282

// 第一想法是 动态规划,类似于之前的跳跃游戏.
// 看数据范围， dp 可能超时,也是硬着头皮写了出来 dp 的解法.
//
// 2. 看题解, 脑筋急转弯, 😹,
// https://leetcode.cn/problems/reach-end-of-array-with-max-score/solutions/2908950/yi-tu-miao-dong-tan-xin-pythonjavacgo-by-tfua/
func findMaximumScore(nums []int) int64 {
	var (
		n         = len(nums)
		result    int
		maxHeight int
	)
	for i := 0; i < n-1; i++ {
		if nums[i] > maxHeight {
			maxHeight = nums[i]
		}
		result += maxHeight
	}
	return int64(result)
}

func dpSolution(nums []int) int64 {
	var (
		n  = len(nums)
		dp = make([]int, n)
	)
	dp[0] = 0
	for i := 1; i < n; i++ {
		dp[i] = -1
	}
	// process
	for i := 0; i < n-1; i++ {
		if nums[i] == -1 {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j != n-1 && nums[i] >= nums[j] {
				continue
			}
			dp[j] = max(dp[j], dp[i]+(j-i)*nums[i])
		}
	}
	if dp[n-1] == -1 {
		return 0
	}
	return int64(dp[n-1])
}
