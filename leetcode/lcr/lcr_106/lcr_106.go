package lcr_106

const (
	UNKNOWN = 0
	RED     = 1
	BLUE    = 2
)

// https://leetcode.cn/problems/vEAB3K/
func isBipartite(graph [][]int) bool {
	//
	var (
		colors = make([]int, len(graph))
		queue  = make([]int, 0, len(graph))
	)
	// init color
	for i := range colors {
		colors[i] = UNKNOWN
	}
	for i := 0; i < len(graph); i++ {
		if colors[i] != UNKNOWN {
			continue
		}
		queue = append(queue, i)
		// 当前颜色
		cur := BLUE
		for len(queue) > 0 {
			size := len(queue)
			for i := 0; i < size; i++ {
				item := queue[i]
				if colors[item] == UNKNOWN {
					// 将当前节点染色
					queue = append(queue, graph[item]...)
					colors[item] = cur
				} else if colors[item] != cur {
					return false
				}
			}
			queue = queue[size:]
			cur = 3 - cur
		}
	}

	return true
}

// dfsSolution
func isBipartiteDFS(graph [][]int) bool {

	//
	var (
		colors = make([]int, len(graph))
	)
	for i := 0; i < len(graph); i++ {
		if colors[i] != UNKNOWN {
			continue
		}
		if !dfs(graph, colors, i, RED) {
			return false
		}
	}
	return true
}

func dfs(graph [][]int, colors []int, node int, color int) bool {
	// terminate
	if colors[node] == UNKNOWN {
		// colors[node] == UNKNOWN
		colors[node] = color
		for _, next := range graph[node] {
			if !dfs(graph, colors, next, 3-color) {
				return false
			}
		}
		return true
	}
	return colors[node] == color

}
