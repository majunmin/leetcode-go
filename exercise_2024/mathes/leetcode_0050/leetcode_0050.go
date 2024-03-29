package leetcode_0050

// https://leetcode.cn/problems/powx-n/
func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1.0 / quickMul(x, -n)
	}
	return quickMul(x, n)
}

func quickMul(x float64, n int) float64 {
	var (
		gain   = x
		result = 1.0
	)
	for n > 0 {
		if n&1 == 1 {
			result *= gain
		}
		n = n >> 1
	}
	return result
}
