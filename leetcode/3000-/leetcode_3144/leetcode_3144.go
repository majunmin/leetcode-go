package leetcode_3144

type item struct {
}

// https://leetcode.cn/problems/minimum-substring-partition-of-equal-character-frequency/
func minimumSubstringsInPartition(s string) int {
	//动态规划.
	// dp[i]: s[o:i] 可以分割最少多少个平衡串
	// dp[i] = min(dp[i], dp(j-1)+1 when dp[j:i]是一个平衡串)

	// 判断平衡串 si 的一个小技巧:
	// si 某个字符的最大频率 * 字符数 == len(si)

	var (
		n  = len(s)
		dp = make([]int, n+1)
	)
	// init state
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		// 记录每个字符出现的频率
		freqs := make(map[byte]int)
		dp[i] = dp[i-1] + 1 // dp 最大可能的值就是 dp[i-1]+1
		maxFreq := 0
		for j := i; j > 0; j-- {
			freqs[s[j-1]]++
			maxFreq = max(maxFreq, freqs[s[j-1]])
			if maxFreq*len(freqs) == i-j+1 {
				dp[i] = min(dp[i], dp[j-1]+1)
			}
		}
	}
	return dp[n]
}
