package leetcode_2642

import "math"

type SearchGraph interface {
	AddEdge(edge []int)
	ShortestPath(node1 int, node2 int) int
}

// 有向无环图 求最短路径， 经典的 Dijkstra 算法(不能有负权边) 或者 Floyd算法
type Graph struct {
	//矩阵的实现方式,稀疏的
	n   int
	adj [][]int
}

func Constructor(n int, edges [][]int) Graph {
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]int, n)
		for j := 0; j < n; j++ {
			adj[i][j] = math.MaxInt //标记不可达
		}
	}
	graph := Graph{n: n, adj: adj}
	for _, e := range edges {
		graph.AddEdge(e)
	}
	return graph
}

func (this *Graph) AddEdge(edge []int) {
	this.adj[edge[0]][edge[1]] = edge[2]
}

func (this *Graph) ShortestPath(node1 int, node2 int) int {
	// start node1
	dsts := make(map[int]int)
	dsts[node1] = 0
	queue := make([]int, 0)
	queue = append(queue, node1)
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		for next := 0; next < this.n; next++ {
			if this.adj[item][next] == math.MaxInt {
				continue
			}

			if _, exist := dsts[next]; exist && dsts[next] <= dsts[item]+this.adj[item][next] {
				continue
			}
			dsts[next] = dsts[item] + this.adj[item][next]
			queue = append(queue, next)
		}
	}
	if _, exist := dsts[node2]; !exist {
		return -1
	}
	return dsts[node2]
}

/**
 * Your Graph object will be instantiated and called as such:
 * obj := Constructor(n, edges);
 * obj.AddEdge(edge);
 * param_2 := obj.ShortestPath(node1,node2);
 */
