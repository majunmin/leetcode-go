package floyd

import "math"

const (
	INF = math.MaxInt / 3
)

type Graph struct {
	adj [][]int
	n   int
}

func Constructor(n int, edges [][]int) Graph {
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]int, n)
		for j := 0; j < n; j++ {
			adj[i][j] = INF
		}
		adj[i][i] = 0
	}
	g := Graph{
		adj: adj,
		n:   n,
	}
	for _, e := range edges {
		adj[e[0]][e[1]] = e[2]
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			if g.adj[i][k] == INF {
				continue
			}
			for j := 0; j < n; j++ {
				adj[i][j] = min(adj[i][j], adj[i][k]+adj[k][j])
			}
		}
	}
	return g
}

func (this *Graph) AddEdge(edge []int) {
	x, y, w := edge[0], edge[1], edge[2]
	// 无需更新
	if this.adj[x][y] <= w {
		return
	}
	for i := 0; i < this.n; i++ {
		for j := 0; j < this.n; j++ {
			this.adj[i][j] = min(this.adj[i][j], this.adj[i][x]+w+this.adj[y][j])
		}
	}
}

func (this *Graph) ShortestPath(node1 int, node2 int) int {
	if this.adj[node1][node2] == INF {
		return -1
	}
	return this.adj[node1][node2]
}

/**
 * Your Graph object will be instantiated and called as such:
 * obj := Constructor(n, edges);
 * obj.AddEdge(edge);
 * param_2 := obj.ShortestPath(node1,node2);
 */
