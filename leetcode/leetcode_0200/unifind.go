package leetcode_0200

type UnionFind struct {
	parent []int
	rank   []int // 秩
	count  int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 1; i < n; i++ {
		parent[i] = i
		rank[i] = 1
	}
	return &UnionFind{
		parent: parent,
		rank:   rank,
		count:  n,
	}
}

func (uf *UnionFind) FindR(i int) int {
	if uf.parent[i] == i {
		return i
	}
	// 路径压缩 / 可选
	r := uf.FindR(uf.parent[i])
	uf.parent[i] = r
	return r
}

func (uf *UnionFind) Find(i, j int) bool {
	return uf.FindR(i) == uf.FindR(j)
}

func (uf *UnionFind) Union(i, j int) {
	if uf.Find(i, j) {
		return
	}
	if uf.rank[i] > uf.rank[j] {
		uf.parent[uf.FindR(j)] = uf.FindR(i)
	} else if uf.rank[i] < uf.rank[j] {
		uf.parent[uf.FindR(i)] = uf.FindR(j)
	} else {
		uf.parent[uf.FindR(i)] = uf.FindR(j)
		uf.rank[j]++
	}
	uf.count--
}

func (uf UnionFind) Count() int {
	return uf.count
}
