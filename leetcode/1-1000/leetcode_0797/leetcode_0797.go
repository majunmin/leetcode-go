package leetcode_0797

// https://leetcode.cn/problems/all-paths-from-source-to-target/
func allPathsSourceTarget(graph [][]int) [][]int {
	var (
		n      = len(graph)
		result = make([][]int, 0)
		dfs    func(cur int, path []int)
	)
	dfs = func(cur int, path []int) {
		if cur == n-1 {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}
		for _, next := range graph[cur] {
			path = append(path, next)
			dfs(next, path)
			path = path[:len(path)-1]
		}
	}
	return result
}
