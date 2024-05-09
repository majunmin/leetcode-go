package leetcode_0886

// https://leetcode.cn/problems/possible-bipartition/
// 判断是否是二分图
// 图论：
// 奇数环:
// a dislike b // a  b
// b dislike c // c
// c dislike a //     a
//
// 偶数环:
// a dislike b // a  b
// b dislike c // c
// c dislike d //    d
// d dislike a //
func possibleBipartition(n int, dislikes [][]int) bool {
	//paramcheck

	// build 图
	g := make([][]int, n+1)
	for _, edge := range dislikes {
		g[edge[0]] = append(g[edge[0]], edge[1])
		g[edge[1]] = append(g[edge[1]], edge[0])
	}

	color := make([]int, n+1) // 1: color1 2: color2
	//
	for i := 1; i <= n; i++ {
		if color[i] == 0 { // 染上一种颜色
			color[i] = 1
		}

		if !dfs(i, g, color) {
			return false
		}
	}
	return true
}

// cur 当前 点
func dfs(cur int, g [][]int, color []int) bool {
	for _, next := range g[cur] {
		// terminate
		if color[next] == color[cur] {
			return false
		}
		// 继续向下染色
		if color[next] == 0 {
			color[next] = 3 - color[cur] // 填充相反的颜色
			if !dfs(next, g, color) {
				return false
			}
		}
	}
	return true
}
