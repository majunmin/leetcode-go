package leetcode_2571

// 贪心做法:
// 高比特位会收到低比特位的影响， 优先处理低比特位.
// 获取低比特位: lb = x & -x
func minOperations(n int) int {
	var cnt int
	for n > 0 {
		lb := lowbit(n)
		if n&(lb<<1) > 0 {
			n += lb // 有两个连续的 bit位, 低位加法
		} else {
			n -= lb // 不存在连个连续的bit位, 减法
		}
		cnt++
	}
	return cnt
}

func lowbit(x int) int {
	return x & (-x)
}
