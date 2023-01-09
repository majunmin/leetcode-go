package offer_064

// https://leetcode.cn/problems/qiu-12n-lcof/?envType=study-plan&id=lcof
// 等差数列  sum = n(n+1)/2  => n(n+1) >> 1
func sumNums(n int) int {

	ans := 0
	num, temp := n, n+1
	// 快速乘法
	// 这个方法经常被用于两数相乘取模的场景，如果两数相乘已经超过数据范围，但取模后不会超过，我们就可以利用这个方法来拆位取模计算贡献，保证每次运算都在数据范围内。
	for temp > 0 {
		if temp&1 == 1 {
			ans += num
		}
		temp >>= 1
		num <<= 1
	}
	return ans >> 1
}

func solution1(n int) int {
	ans := 0
	var sum func(int) bool
	sum = func(n int) bool {
		ans += n
		// terminate
		return n > 0 && sum(n-1)
	}
	sum(n)
	return ans
}
