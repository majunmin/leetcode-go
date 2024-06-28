package unifind

// 路径压缩
// 按秩合并
type UniFind struct {
	parent []int `json:"parent,omitempty"`
	// 添加一个 rank(秩)数组,记录树的高度.
	rank  []int `json:"rank,omitempty"`
	count int   `json:"count,omitempty"`
}

func NewUniFind(count int) *UniFind {
	parent := make([]int, count)
	rank := make([]int, count)
	for i := 0; i < count; i++ {
		parent[i] = i
		rank[i] = 1
	}
	return &UniFind{
		parent: parent,
		rank:   rank,
		count:  count,
	}
}

func (uf *UniFind) Find(i int) int {
	// 隔代路径压缩
	//for uf.parent[i] != i {
	//	uf.parent[i] = uf.parent[uf.parent[i]]
	//	i = uf.parent[i]
	//}
	//return i

	// 全路径压缩
	if uf.parent[i] != i {
		uf.parent[i] = uf.Find(uf.parent[i])
	}
	return i
}

func (uf *UniFind) Union(i, j int) {
	rootI := uf.Find(i)
	rootJ := uf.Find(j)
	if rootI == rootJ {
		return
	}
	// 按秩合并
	//if uf.rank[rootI] < uf.rank[rootJ] {
	//	uf.parent[rootI] = rootJ
	//} else if uf.rank[rootI] > uf.rank[rootJ] {
	//	uf.parent[rootJ] = rootI
	//} else {
	//	uf.parent[rootJ] = rootI
	//}
	//uf.rank[rootI] += 1
	uf.parent[rootI] = rootJ
	uf.count--
}

func (uf *UniFind) Count() int {
	return uf.count
}
