package leetcode_100228

// https://leetcode.cn/contest/weekly-contest-390/problems/apply-operations-to-make-sum-of-array-greater-than-or-equal-to-k/
func minOperations(k int) int {
	result := k - 1
	for i := k - 1; i >= 1; i-- {
		copyCnt := k/i - 1
		incrCnt := i - 1
		if k%i != 0 {
			copyCnt += 1
		}
		result = min(result, copyCnt+incrCnt)
	}
	return result
}
