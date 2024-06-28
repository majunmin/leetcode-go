package leetcode_2741

import "fmt"

// https://leetcode.cn/problems/special-permutations/
// 超时
func specialPerm(nums []int) int {
	// 有条件的全排列
	var (
		n       = len(nums)
		result  int
		visited = make([]bool, n)
		dfs     func([]int)
	)
	dfs = func(path []int) {
		// terminate
		if len(path) == n {
			fmt.Println(path)
			result++
			return
		}
		// for choice in choiceList
		for i, num := range nums {
			if visited[i] {
				continue
			}
			path = append(path, num)
			visited[i] = true
			if len(path) == 1 || (path[len(path)-1]%path[len(path)-2] == 0 || path[len(path)-2]%path[len(path)-1] == 0) {
				dfs(path)
			}
			visited[i] = false
			path = path[:len(path)-1]
		}
	}
	dfs([]int{})
	return result
}

// 与题目 526 类似, 建议先做  526
func specialPerm2(nums []int) int {
	// 有条件的全排列
	var (
		n    = len(nums)
		mask = 1<<n - 1
		dp   = make([][]int, n)
	)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, mask)
	}
	// dp[i][state] : 前 i个数, 选择 state 个,
	for i := 0; i < n; i++ {

	}
}
