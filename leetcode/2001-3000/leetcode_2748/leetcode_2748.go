package leetcode_2748

// https://leetcode.cn/problems/number-of-beautiful-pairs/
func countBeautifulPairs(nums []int) int {
	var (
		result int
	)
	for i := 0; i < len(nums); i++ {
		for nums[i] > 10 {
			nums[i] /= 10
		}
		for j := i + 1; j < len(nums); j++ {
			if gcd(nums[i], nums[j]%10) == 1 {
				result++
			}
		}
	}
	return result
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
