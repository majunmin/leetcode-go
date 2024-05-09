package leetcode_1969

// 题解： https://leetcode.cn/problems/minimum-non-zero-product-of-the-array-elements/solutions/2688955/shu-zu-yuan-su-de-zui-xiao-fei-ling-chen-xti0/
// https://leetcode.cn/problems/minimum-non-zero-product-of-the-array-elements/?envType=daily-question&envId=2024-03-20

var mod = int(1e9 + 7)

func minNonZeroProduct(p int) int {
	x := (1 << p) - 1
	y := x - 1
	return (x * quickPow(y, (1<<(p-1))-1)) % mod
}

func quickPow(x int, n int) int {
	var (
		result = 1
		gain   = x
	)
	for n > 0 {
		result = (result * gain) % mod
		gain = (gain * gain) % mod
		n >>= 1
	}
	return result
}
