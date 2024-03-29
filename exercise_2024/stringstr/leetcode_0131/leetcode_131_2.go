package leetcode_0131

// https://leetcode.cn/problems/palindrome-partitioning/
func partition2(s string) [][]string {
	// 利用动态规划先把结果算出来
	// param check
	if len(s) < 1 {
		return nil
	}
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	// dp 规划找到所有的  回文子串
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			dp[i][j] = s[i] == s[j] && (j-i < 2 || dp[i+1][j-1])
		}
	}
	//
	var result [][]string
	backtrace2(&result, s, dp, []string{}, 0, n-1)
	return result
}

func backtrace2(result *[][]string, s string, dp [][]bool, path []string, left int, right int) {
	// terminate
	if left > right {
		tmp := make([]string, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
		return
	}
	for i := left; i <= right; i++ {
		// 判断前缀是否是  回文串
		if dp[left][i] {
			path = append(path, s[left:i+1])
			backtrace2(result, s, dp, path, i+1, right)
			path = path[:len(path)-1]
		}
	}
}
