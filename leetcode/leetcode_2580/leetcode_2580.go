package leetcode_2580

import (
	"math"
	"slices"
)

// https://leetcode.cn/problems/count-ways-to-group-overlapping-ranges/
func countWays(ranges [][]int) int {
	// param check
	if len(ranges) == 0 {
		return 0
	}
	slices.SortFunc(ranges, func(a, b []int) int {
		return a[0] - b[0]
	})
	// 将 ranges 按照 start[0]排序
	var (
		end = math.MinInt
		//cnt    int
		result = 1
	)

	for i := 0; i < len(ranges); i++ {
		// 添加一个区间
		if ranges[i][0] > end {
			end = ranges[i][1]
			//cnt++
			result = (result << 1) % (1e9 + 7)
			continue
		}
		// 两个区间 存在交集， 合并
		end = max(end, ranges[i][1])

	}
	return result
}

// 用一种笨方法解决吧
// 并查集 + 快速幂的解法 算是复习一下其他算法
func countWays2(ranges [][]int) int {
	// param check
	if len(ranges) == 0 {
		return 0
	}
	slices.SortFunc(ranges, func(a, b []int) int {
		return a[0] - b[0]
	})
	//
	uf := NewUnionFind(len(ranges))
	for i := 1; i < len(ranges); i++ {
		if ranges[i-1][1] >= ranges[i][0] {
			uf.union(i-1, i)
		}
	}
	return pow(2, uf.count())
}

func pow(base int, p int) int {
	if p == 0 {
		return 1
	}
	cur := pow(base, p/2)
	if p&1 == 1 {
		return cur * cur * base % (1e9 + 7)
	}
	return cur * cur % (1e9 + 7)
}

type UnionFind struct {
	p   []int
	cnt int
}

func NewUnionFind(cnt int) *UnionFind {
	p := make([]int, cnt)
	for i := 0; i < len(p); i++ {
		p[i] = i
	}
	return &UnionFind{
		p:   p,
		cnt: cnt,
	}
}

func (uf *UnionFind) union(i, j int) {
	if uf.find(i) == uf.find(j) {
		return
	}
	uf.p[j] = i
	uf.cnt--
}

func (uf *UnionFind) count() int {
	return uf.cnt
}

func (uf *UnionFind) find(i int) int {
	if i == uf.p[i] {
		return i
	}
	root := uf.find(uf.p[i])
	uf.p[i] = root
	return root
}

func (uf *UnionFind) Find(i, j int) bool {
	// todo: implements me
	return false
}
