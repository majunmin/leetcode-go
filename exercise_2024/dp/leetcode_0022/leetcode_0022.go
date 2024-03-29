package leetcode_0022

func generateParenthesis(n int) []string {
	// dpSolution
	// 1 -> ()
	// 2 -> () + '(' + ')',
	dp := make([][]string, n+1)
	dp[0] = []string{""}
	dp[1] = []string{"()"}
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			for _, left := range dp[j] {
				for _, right := range dp[i-j-1] {
					dp[i] = append(dp[i], left+"("+right+")")
				}
			}
		}
	}
	return dp[n]
}

func dfsSolution(n int) []string {
	var result []string
	dfs(n, 0, 0, "", &result)
	return result
}

func dfs(n int, l int, r int, path string, result *[]string) {
	// terminate
	if r == n {
		*result = append(*result, path)
		return
	}
	if l < n {
		dfs(n, l+1, r, path+"(", result)
	}
	if r < l {
		dfs(n, l, r+1, path+")", result)
	}
}
