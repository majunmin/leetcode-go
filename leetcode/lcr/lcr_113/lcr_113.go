package lcr_113

// https://leetcode.cn/problems/QA2IGt/
func findOrder(numCourses int, prerequisites [][]int) []int {
	// 拓扑排序
	//
	//构建图 +
	var (
		adj     = make([][]int, numCourses)
		in      = make([]int, numCourses) // 记录节点入度, 入度为 0 才可以修
		result  = make([]int, 0, numCourses)
		visited = make([]bool, numCourses)
	)
	for _, edge := range prerequisites {
		adj[edge[1]] = append(adj[edge[1]], edge[0])
		in[edge[0]]++
	}
	var queue = make([]int, 0)
	for i, p := range in {
		if p == 0 {
			queue = append(queue, i)
			result = append(result, i)
			visited[i] = true
		}
	}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			item := queue[i]
			for _, next := range adj[item] {
				in[next]--
				// 去重?
				if in[next] == 0 && !visited[next] {
					visited[next] = true
					result = append(result, next)
					queue = append(queue, next)
				}
			}
		}
		queue = queue[size:]
	}
	if len(result) < numCourses {
		return nil
	}
	return result
}
