package offer_016

// https://leetcode.cn/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/?envType=study-plan&id=lcof
// 2^10 = 2^(1111) = 2^8 * 2^4 * 2^2 * 2^1 * 2^0
func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}

	res := float64(1) // 2^0
	for n > 0 {
		if n&1 == 1 {
			res *= x
		}
		x *= x
		n >>= 1
	}
	return res
}

func solution1(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		return 1 / myPow(x, -n)
	}

	if n&1 == 1 {
		return myPow(x, n>>1) * myPow(x, n>>1) * x
	} else {
		return myPow(x, n>>1) * myPow(x, n>>1)
	}
}
