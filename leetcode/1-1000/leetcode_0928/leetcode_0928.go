package leetcode_0928

import (
	"sort"
)

// https://leetcode.cn/problems/minimize-malware-spread-ii/
func minMalwareSpread(graph [][]int, initial []int) int {
	var (
		n         = len(graph)
		isInitial = make(map[int]bool, len(initial))
		adj       = make([][]int, n)
		visited   = make(map[int]bool, n)
		cnts      = make(map[int]int)
	)
	sort.Ints(initial)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if graph[i][j] == 1 {
				adj[i] = append(adj[i], j)
				adj[j] = append(adj[j], i)
			}
		}
	}
	for _, node := range initial {
		isInitial[node] = true
	}

	var (
		nodeId = -1
		size   int
	)
	var dfs func(cur int)
	dfs = func(cur int) {
		visited[cur] = true
		size++
		for _, next := range adj[cur] {
			if isInitial[next] && nodeId != next {
				// 避免重复计算  感染节点.
				if nodeId == -2 || nodeId == next {
					continue
				}
				// nodeId != -2
				if nodeId == -1 { // 如果没有感染节点
					nodeId = next
				} else { // 已存在一个感染节点
					nodeId = -2
				}
			} else if !visited[next] {
				dfs(next)
			}
		}
	}

	// dfs 处理
	// 状态机 处理 nodeId: -1: 不存在 感染节点  >0 :只有一个被感染节点, 节点id   -2: 有多个被感染节点
	// 找到 只与一个感染节点相连的 好节点并计数(需要遍历好节点)(逆向思维)
	for i := 0; i < n; i++ {
		if visited[i] || isInitial[i] {
			continue
		}
		nodeId = -1
		size = 0
		dfs(i)
		if nodeId >= 0 {
			cnts[nodeId] += size
		}
	}

	result := initial[0]
	maxCnt := cnts[result]
	for k, cnt := range cnts {
		if cnt > maxCnt {
			result = k
			maxCnt = cnt
		}
	}
	return result
}
