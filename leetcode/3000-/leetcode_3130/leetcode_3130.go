package leetcode_3130

// https://leetcode.cn/problems/find-all-possible-stable-binary-arrays-ii/description/
// 题解
// https://leetcode.cn/problems/find-all-possible-stable-binary-arrays-ii/solutions/2758868/dong-tai-gui-hua-cong-ji-yi-hua-sou-suo-37jdi/
const (
	MOD = 1_000_000_007
)

func numberOfStableArrays(zero int, one int, limit int) int {
	var (
		memo = make([][][2]int, zero+1)
	)
	// init memo
	for i := 0; i < zero+1; i++ {
		memo[i] = make([][2]int, one+1)
		for j := 0; j < one+1; j++ {
			memo[i][j][0] = -1
			memo[i][j][1] = -1
		}
	}
	// 第 i+j的位置 填 0 or 1
	// dfs(i,j,k): 由i个+ j个 1 组成的 && 第 i +j 位填 k 的稳定数组的方案数.
	return (dfs(zero, one, 0, limit, memo) + dfs(zero, one, 1, limit, memo)) % MOD
}

// 由 i个 0， j个 1 构成稳定数组的方案数. 其中 第 i + j 个位置填 k(0|1).
func dfs(i, j int, k, limit int, memo [][][2]int) int {
	// terminate i == 0 && j == 0
	if i == 0 { // 0 消耗完了，之后不能再填 0 了 && 剩余的 1的个数 不能 > limit
		if k == 1 && j <= limit { // 剩下的1 的个数 多余 limit, 不符合条件.
			return 1
		}
		return 0
	}
	if j == 0 { // 1 消耗完了.
		if k == 0 && i <= limit {
			return 1
		}
		return 0
	}
	// process
	// 重复计算.
	if memo[i][j][k] != -1 {
		return memo[i][j][k]
	}
	if k == 0 {
		// 第 i+j-1的位置 填 0 or 1
		ret := dfs(i-1, j, 0, limit, memo) + dfs(i-1, j, 1, limit, memo)
		if i > limit {
			// 移除 从i开始往前数 连续limit+1个0的方案数.
			ret += MOD - dfs(i-limit-1, j, 1, limit, memo)
		}
		memo[i][j][k] = ret % MOD
	} else {
		ret := dfs(i, j-1, 1, limit, memo) + dfs(i, j-1, 0, limit, memo)
		if j > limit {
			// 移除 从j开始往前数 连续limit+1个1的方案数.
			ret += MOD - dfs(i, j-limit-1, 0, limit, memo)
		}
		memo[i][j][k] = ret % MOD
	}
	return memo[i][j][k]
}
