package leetcode_0050

// https://leetcode.cn/problems/powx-n/
func myPow(x float64, n int) float64 {
	return fastPow(x, n)
}

func fastPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		x = 1 / x
		n = -n
		return fastPow(x, n)
	}

	var ans float64 = 1
	// n 用 二进制表示
	// 7 - > 0   1    1     1
	//   - > x^  x^4 x^2  x^1
	//此时的 x 表示 当前 位 的 值
	for n > 0 {
		if n&1 == 1 {
			ans *= x
		}
		x *= x
		n >>= 1
	}
	return ans

}

func recurSolution(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	return powHelp(x, n)
}

func powHelp(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	subPow := myPow(x, n/2)
	if n%2 == 0 {
		return subPow * subPow
	} else {
		return subPow * subPow * x
	}
}
