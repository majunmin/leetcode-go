package leetcode_1345

// https://leetcode.cn/problems/jump-game-iv/
// https://leetcode.cn/problems/jump-game-iv/
func minJumps(arr []int) int {
	// bfs 构建图
	// key sourceNum  value: destination position
	edges := make(map[int][]int)
	for i, num := range arr {
		if _, exist := edges[num]; !exist {
			edges[num] = make([]int, 0)
		}
		edges[num] = append(edges[num], i)
	}

	// bfs
	var steps int
	visited := make([]bool, len(arr))
	visited[0] = true
	queue := make([]int, 0)
	queue = append(queue, 0)
	for len(queue) > 0 {
		size := len(queue)
		// process
		for i := 0; i < size; i++ {
			item := queue[i]
			if len(arr)-1 == item {
				// find result
				return steps
			}
			// 相同数字的跳数
			for _, idx := range edges[arr[item]] {
				if !visited[idx] {
					visited[idx] = true
					queue = append(queue, idx)
				}
			}
			// 剪枝， 否则会运行超时
			delete(edges, arr[item])
			// neiborhood
			lIdx, rIdx := item-1, item+1
			if lIdx >= 0 && lIdx < len(arr) && !visited[lIdx] {
				visited[lIdx] = true
				queue = append(queue, lIdx)
			}

			if rIdx >= 0 && rIdx < len(arr) && !visited[rIdx] {
				visited[rIdx] = true
				queue = append(queue, rIdx)
			}
		}
		steps++
		queue = queue[size:]
	}
	return -1
}
