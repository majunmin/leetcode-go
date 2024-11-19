package leetcode_3243

// https://leetcode.cn/problems/shortest-distance-after-road-addition-queries-i/
func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	var (
		// 从 0 -> i 的最短距离
		dis = make([]int, n)
		adj = make([][]int, n)
	)
	for i := 1; i < n; i++ {
		dis[i] = i
		adj[i-1] = append(adj[i-1], i)
	}

	//
	var (
		result = make([]int, 0, len(queries))
	)
	for _, q := range queries {
		adj[q[0]] = append(adj[q[0]], q[1])
		//if q[0] >= q[1] {
		//	result = append(result, dis[n-1])
		//	continue
		//}
		// bfs
		queue := make([]int, 0)
		queue = append(queue, q[0])
		for len(queue) > 0 {
			size := len(queue)
			for i := 0; i < size; i++ {
				for _, next := range adj[queue[i]] {
					dis[next] = min(dis[next], dis[queue[i]]+1)
					queue = append(queue, next)
				}
			}
			queue = queue[size:]
		}
		result = append(result, dis[n-1])
	}
	return result
}
