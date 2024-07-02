package leetcode_3115

var (
	N       = 101
	isPrime = make([]bool, N)
)

// 1 3 5 7 11
func init() {
	// 0, 1 不是质数
	for i := 2; i < N; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i < N; i++ {
		if isPrime[i] {
			for j := i * i; j < N; j += i {
				isPrime[j] = false
			}
		}
	}
}

// https://leetcode.cn/problems/maximum-prime-difference/
func maximumPrimeDifference(nums []int) int {
	var (
		result   int
		startIdx = -1
	)
	for i := 0; i < len(nums); i++ {
		if isPrime[nums[i]] {
			if startIdx == -1 {
				startIdx = i
			} else {
				result = i - startIdx
			}
		}
	}
	return result
}
