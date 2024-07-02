package lcr_085

// https://leetcode.cn/problems/IDBivT/
func generateParenthesis(n int) []string {
	// 动态规划
	var (
		dp = make([][]string, n+1)
	)
	//init state
	dp[0] = []string{""}
	dp[1] = []string{"()"}
	for i := 2; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			for _, p1 := range dp[j] {
				for _, p2 := range dp[i-j-1] {
					dp[i] = append(dp[i], "("+p1+")"+p2)
				}
			}
		}
	}
	return dp[n]
}

func dfsSolution(n int) []string {
	var result []string
	generate(0, 0, n, "", &result)
	return result
}

func generate(l int, r int, n int, path string, result *[]string) {
	// terminate
	if l+r == n*2 {
		*result = append(*result, path)
	}
	if l < n {
		generate(l+1, r, n, path+"(", result)
	}
	if r < l {
		generate(l, r+1, n, path+")", result)
	}
}
