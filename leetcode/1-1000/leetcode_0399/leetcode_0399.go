package leetcode_0399

// https://leetcode.cn/problems/evaluate-division/?envType=study-plan-v2&envId=top-interview-150
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 并查集解法
	var (
		flag2Index = make(map[string]int)
		result     = make([]float64, 0, len(queries))
	)
	// 为每一个元素编号,方便使用并查集
	for _, varialbes := range equations {
		var1 := varialbes[0]
		var2 := varialbes[1]
		if _, exist := flag2Index[var1]; !exist {
			flag2Index[var1] = len(flag2Index)
		}
		if _, exist := flag2Index[var2]; !exist {
			flag2Index[var2] = len(flag2Index)
		}
	}
	uf := newUnionFind(len(flag2Index))

	for i, variables := range equations {
		var1, var2 := variables[0], variables[1]
		value := values[i]
		uf.union(flag2Index[var1], flag2Index[var2], value)
	}
	for _, query := range queries {
		var1, var2 := query[0], query[1]
		idx1, exist1 := flag2Index[var1]
		idx2, exist2 := flag2Index[var2]
		if !exist1 || !exist2 {
			result = append(result, -1.0)
		} else {
			result = append(result, uf.isConnected(idx1, idx2))
		}
	}
	return result
}

type unionFind struct {
	parent []int
	weight []float64
}

func newUnionFind(size int) *unionFind {
	p := make([]int, size)
	weight := make([]float64, size)

	for i := 0; i < size; i++ {
		p[i] = i
		weight[i] = 1.0
	}
	return &unionFind{
		parent: p,
		weight: weight,
	}
}

// 1  (2.0) 2 (1.5)   3(1)
// a -> b -> c
// x / y = value
func (uf *unionFind) union(x, y int, value float64) {
	rootX := uf.find(x)
	rootY := uf.find(y)
	if rootX == rootY {
		return
	}
	uf.parent[rootX] = rootY
	//y = weight[y]rootY
	//x = weight[x]rootX = value*y = value * weight[y] rootY
	//==> rootX/rootY = value * weight[y]/ weight[x]
	// 自己画图最容易理解
	uf.weight[rootX] = uf.weight[y] * value / uf.weight[x]
}

func (uf *unionFind) find(x int) int {
	if x != uf.parent[x] {
		ori := uf.parent[x]
		uf.parent[x] = uf.find(uf.parent[x])
		uf.weight[x] *= uf.weight[ori]
		x = uf.parent[x]
	}
	return uf.parent[x]
}

func (uf *unionFind) isConnected(x, y int) float64 {
	rootX := uf.find(x)
	rootY := uf.find(y)
	if rootX == rootY {
		// 在一个连通分量里面
		return uf.weight[x] / uf.weight[y]
	} else {
		return -1.0
	}
}
