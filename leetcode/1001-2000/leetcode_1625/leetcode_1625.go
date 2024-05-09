package leetcode_1625

// https://leetcode.cn/problems/defuse-the-bomb/?envType=daily-question&envId=2024-05-05
func decrypt(code []int, k int) []int {
	// param check
	if len(code) == 0 {
		return nil
	}
	var (
		n      = len(code)
		result = make([]int, n)
	)
	if k == 0 {
		return result
	}
	if k > 0 {
		for i := 0; i < n; i++ {
			for j := i + 1; j <= i+k; j++ {
				result[i] += code[(j+n)%n]
			}
		}
		return result
	}
	// k < 0
	for i := 0; i < n; i++ {
		for j := i - 1; j >= i-k; j-- {
			result[i] += code[(j+n)%n]
		}
	}
	return result
}
