package leetcode_0547

// https://leetcode.com/problems/number-of-provinces/
func findCircleNum(isConnected [][]int) int {
	// param check
	n := len(isConnected)
	if n == 0 || n != len(isConnected[0]) {
		return 0
	}
	uf := newUnionFind(n)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if i == j || isConnected[i][j] == 0 {
				continue
			}
			//isConnected[i][j] == 1
			uf.union(i, j)
		}
	}
	return uf.Count()
}

type unionFind struct {
	parent []int
	count  int
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &unionFind{
		parent: parent,
		count:  n,
	}
}

func (uf *unionFind) Count() int {
	return uf.count
}

func (uf *unionFind) find(i int) int {
	for uf.parent[i] != i {
		uf.parent[i] = uf.parent[uf.parent[i]]
		i = uf.parent[i]
	}
	return i
}

func (uf *unionFind) union(i, j int) {
	rootI := uf.find(i)
	rootJ := uf.find(j)
	if rootI == rootJ {
		return
	}
	// 按秩合并...
	uf.parent[rootI] = rootJ
	uf.count--
}
