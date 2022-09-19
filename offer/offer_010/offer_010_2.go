package offer_010

//https://leetcode.cn/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/?envType=study-plan&id=lcof
func numWays(n int) int {

	if n == 0 || n == 1 {
		return 1
	}

	p1, p2 := 1, 1
	res := 0
	for i := 2; i <= n; i++ {
		res = (p1 + p2) % (1e9 + 7)
		p1, p2 = p2, res
	}
	return res
}
