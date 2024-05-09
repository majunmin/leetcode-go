package leetcode_1483

import "math/bits"

// https://leetcode.cn/problems/kth-ancestor-of-a-tree-node/
type TreeAncestor struct {
	pa [][]int
}

// 动态规划
// 状态定义:
// pa[i][j] : 节点 i 的 (2^j) 级父节点
// pa[i][0] = parent[i]  // 父节点
// pa[i][1] = parent[parent[i]] == pa[pa[i][0]][0] // 爷爷节点
// pa[i][2] = parent[parent[parent[parent[i]]]] = pa[pa[i][1]][1] // 爷爷的爷爷
// pa[i][j+1] = ... = pa[pa[i][j]][j]
func Constructor(n int, parent []int) TreeAncestor {
	m := bits.Len(uint(n))
	pa := make([][]int, n)
	for i := 0; i < n; i++ {
		pa[i] = make([]int, m)
		pa[i][0] = parent[i]
	}
	for i := 0; i < m-1; i++ {
		for j := 0; j < n; j++ {
			if pa[j][i] != -1 {
				pa[j][i+1] = pa[pa[j][i]][i]
			} else {
				pa[j][i+1] = -1
			}
		}
	}
	return TreeAncestor{pa: pa}
}

func (this *TreeAncestor) GetKthAncestor(node int, k int) int {
	m := bits.Len(uint(k))
	cur := node
	for i := 0; i < m; i++ {
		if k>>i&1 == 1 {
			cur = this.pa[cur][i]
			if cur < 0 {
				break
			}
		}
	}
	return cur
}

/**
 * Your TreeAncestor object will be instantiated and called as such:
 * obj := Constructor(n, parent);
 * param_1 := obj.GetKthAncestor(node,k);
 */
