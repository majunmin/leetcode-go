package leetcode_2312

// https://leetcode.cn/problems/selling-pieces-of-wood/

func sellingWood(m int, n int, prices [][]int) int64 {
	return recursiveSolution(m, n, prices)
}

// 记忆化 +递归实现
func recursiveSolution(m int, n int, prices [][]int) int64 {
	pos2Price := make(map[int]int64)
	for _, items := range prices {
		pos2Price[getHash(items[0], items[1])] = int64(items[2])
	}
	memo := make([][]int64, m+1)
	for i := 0; i <= m; i++ {
		memo[i] = make([]int64, n+1)
		for j := 0; j <= n; j++ {
			memo[i][j] = -1
		}
	}
	return dfs(m, n, pos2Price, memo)
}

func dfs(m int, n int, pos2Price map[int]int64, memo [][]int64) int64 {
	if memo[m][n] != -1 {
		return memo[m][n]
	}
	var ret int64
	key := getHash(m, n)
	if _, exist := pos2Price[key]; exist {
		ret = pos2Price[key]
	}
	for i := 1; i <= m; i++ {
		ret = max(ret, dfs(i, n, pos2Price, memo)+dfs(m-i, n, pos2Price, memo))
	}
	for i := 1; i <= n; i++ {
		ret = max(ret, dfs(m, i, pos2Price, memo)+dfs(m, n-i, pos2Price, memo))
	}
	memo[m][n] = ret
	return ret
}

func getHash(x, y int) int {
	return (x << 32) & y
}

// 动态规划,自低向上
func sellingWood2(m int, n int, prices [][]int) int64 {
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	for _, items := range prices {
		dp[items[0]][items[1]] = items[2]
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			for k := 1; k < i; k++ {
				dp[i][j] = max(dp[i][j], dp[k][j]+dp[i-k][j])
			}
			for k := 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[i][j-k])
			}
		}
	}
	return int64(dp[m][n])
}
