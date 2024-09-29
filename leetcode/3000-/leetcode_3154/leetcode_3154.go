package leetcode_3154

type arg struct {
	i, j    int
	preDown bool
}

// https://leetcode.cn/problems/find-number-of-ways-to-reach-the-k-th-stair/
func waysToReachStair(k int) int {
	return dfs(1, 0, k, false, make(map[arg]int))
}

// i: 当前处于第几级台阶
// j 执行了第几次步骤二(jump):
// preDown: 上一次是否执行了步骤一
func dfs(i, j, k int, preDown bool, memo map[arg]int) int {
	if i > k+1 {
		return 0
	}
	if _, exist := memo[arg{i, j, preDown}]; exist {
		return memo[arg{i, j, preDown}]
	}
	var res int
	if i == k {
		res = 1
	}
	res += dfs(i+pow2(j), j+1, k, false, memo)
	if !preDown {
		res += dfs(i-1, j, k, true, memo)
	}
	memo[arg{i, j, preDown}] = res
	return res
}

func pow2(x int) int {
	return 1 << x
}
