package leetcode_1997

// https://leetcode.cn/problems/first-day-where-you-have-been-in-all-the-rooms/?envType=daily-question&envId=2024-03-28
func firstDayBeenInAllRooms(nextVisit []int) int {
	//首先 步数一定是从左往右走的 ， 因为  nextVisit[i] < i
	// 递推解法
	//dp[i] : 第一次到达 i 所需要的步数
	// dp[i+1] = (dp[i] + 1 - dp[nextVisit[i]]) + 1
	M := int(1e9 + 7)
	dp := make([]int, len(nextVisit))
	for i := 0; i < len(nextVisit)-1; i++ {
		j := nextVisit[i]
		dp[i+1] = (dp[i] + 1 + dp[i] - dp[j] + M) % M
	}
	return dp[len(nextVisit)-1]
}

// 会超时
func moniSlution(nextVisit []int) int {
	// param check
	if len(nextVisit) == 1 {
		return 0
	}
	// process
	var (
		cnts   = make([]int, len(nextVisit))
		cnt    int
		result int
		cur    int
	)

	for cnt < len(nextVisit) {
		cnts[cur]++
		if cnts[cur] == 1 {
			cnt++
		}
		if cnts[cur]&1 == 1 {
			cur = nextVisit[cur]
		} else {
			cur = (cur + 1) % len(nextVisit)
		}
		result = (result + 1) % (1e9 + 7)
	}
	return result
}
