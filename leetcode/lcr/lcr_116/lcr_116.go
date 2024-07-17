package lcr_116

// https://leetcode.cn/problems/bLyHh0/
func findCircleNum(isConnected [][]int) int {
	var (
		n       = len(isConnected)
		visited = make([]bool, n)
		result  int
	)
	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		visited[i] = true
		dfs(i, isConnected, visited)
		result++
	}
	return result
}

func dfs(cur int, connected [][]int, visited []bool) {
	for j, item := range connected[cur] {
		if item == 0 || visited[item] {
			continue
		}
		visited[j] = true
		dfs(j, connected, visited)
	}
}

// solution 1
func ufSolution(isConnected [][]int) int {
	var (
		n  = len(isConnected)
		uf = newUnionFind(n)
	)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 {
				uf.union(i, j)
			}
		}
	}
	return uf.count()
}

type unionFind struct {
	p    []int
	rank []int
	cnt  int
}

func newUnionFind(n int) *unionFind {
	var (
		p    = make([]int, n)
		rank = make([]int, n)
	)
	for i := 0; i < n; i++ {
		p[i] = i
		rank[i] = 1
	}
	return &unionFind{
		p:    p,
		rank: rank,
		cnt:  n,
	}
}

func (uf *unionFind) find(i int) int {
	if uf.p[i] != i {
		uf.p[i] = uf.find(uf.p[i])
	}
	return uf.p[i]
}

func (uf *unionFind) union(i, j int) {
	rootI, rootJ := uf.find(i), uf.find(j)
	if rootI == rootJ {
		return
	}
	if uf.rank[i] == uf.rank[j] {
		uf.p[rootJ] = rootI
		uf.rank[rootI]++
	} else if uf.rank[i] > uf.rank[j] {
		uf.p[rootJ] = rootI
	} else {
		uf.p[rootI] = rootJ
	}
	uf.cnt--
}

func (uf *unionFind) count() int {
	return uf.cnt
}

// 1. 遍历图的方式,
// 2. unionfind 并查集的方式
