package leetcode_1423

// https://leetcode.cn/problems/maximum-points-you-can-obtain-from-cards/
func maxScore(cardPoints []int, k int) int {
	// param check
	n := len(cardPoints)
	if n <= k {
		return sum(cardPoints, 0, n-1)
	}

	var total int
	var minWindowSum int
	windowSize := n - k
	for i := 0; i < windowSize; i++ {
		minWindowSum += cardPoints[i]
		total += cardPoints[i]
	}
	curSum := minWindowSum
	for i := windowSize; i < n; i++ {
		curSum += cardPoints[i]
		curSum -= cardPoints[i-windowSize]
		minWindowSum = min(minWindowSum, curSum)
		total += cardPoints[i]
	}

	return total - minWindowSum
}

// 滑动窗口，逆向思维
// 最后剩下的一定是 连续的 (n-k) 的子数组
func sum(points []int, start, end int) int {
	var total int
	for i := start; i <= end; i++ {
		total += points[i]
	}
	return total
}
