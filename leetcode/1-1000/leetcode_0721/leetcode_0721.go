package leetcode_0721

import "slices"

// https://leetcode.cn/problems/accounts-merge/
func accountsMerge(accounts [][]string) [][]string {
	// hash + 并查集
	var (
		email2Index = make(map[string]int)
		email2Name  = make(map[string]string)
	)
	for _, acc := range accounts {
		for i := 1; i < len(acc); i++ {
			if _, exist := email2Index[acc[i]]; !exist {
				email2Index[acc[i]] = len(email2Index)
				email2Name[acc[i]] = acc[0]
			}
		}
	}
	var (
		uf           = newUnionFind(len(email2Index))
		index2Emails = make(map[int][]string)
	)

	for _, account := range accounts {
		firstEmail := account[1]
		firstIdx := email2Index[firstEmail]
		for i := 2; i < len(account); i++ {
			uf.union(firstIdx, email2Index[account[i]])
		}
	}
	for email, i := range email2Index {
		idx := uf.find(i)
		index2Emails[idx] = append(index2Emails[idx], email)
	}
	// 合并结果
	result := make([][]string, 0)
	for _, emails := range index2Emails {
		cur := make([]string, 0, len(emails))
		slices.Sort(emails)
		cur = append(cur, email2Name[emails[0]])
		cur = append(cur, emails...)
		result = append(result, cur)
	}
	return result
}

type UnionFind struct {
	p     []int
	count int
}

func newUnionFind(cap int) *UnionFind {
	p := make([]int, cap)
	for i := 0; i < len(p); i++ {
		p[i] = i
	}
	return &UnionFind{
		p:     p,
		count: 0,
	}
}
func (uf *UnionFind) union(i, j int) {
	rootI := uf.find(i)
	rootJ := uf.find(j)
	if rootI == rootJ {
		return
	}
	uf.p[rootI] = rootJ
}

func (uf *UnionFind) find(i int) int {
	for uf.p[i] != i {
		uf.p[i] = uf.find(uf.p[i])
		i = uf.p[i]
	}
	return uf.p[i]
}
