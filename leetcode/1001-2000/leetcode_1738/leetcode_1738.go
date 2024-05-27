package leetcode_1738

import (
	"container/heap"
	"slices"
)

// https://leetcode.cn/problems/find-kth-largest-xor-coordinate-value/?envType=daily-question&envId=2024-05-27
func kthLargestValue(matrix [][]int, k int) int {
	var (
		m, n   = len(matrix), len(matrix[0])
		pre    = make([][]int, m+1)
		result = make([]int, 0, m*n)
	)
	for i := 0; i <= m; i++ {
		pre[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			pre[i][j] = pre[i-1][j] ^ pre[i][j-1] ^ pre[i-1][j-1] ^ matrix[i-1][j-1]
			result = append(result, pre[i][j])
		}
	}
	slices.Sort(result)
	return result[m*n-k]
}

// 小顶堆的方式
func kthLargestValue2(matrix [][]int, k int) int {
	var (
		m, n = len(matrix), len(matrix[0])
		pre  = make([][]int, m+1)
		hp   = newHeap()
	)
	for i := 0; i <= m; i++ {
		pre[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			pre[i][j] = pre[i-1][j] ^ pre[i][j-1] ^ pre[i-1][j-1] ^ matrix[i-1][j-1]
			if hp.Len() < k {
				heap.Push(hp, pre[i][j])
			} else {
				if pre[i][j] > hp.Peek() {
					heap.Pop(hp)
					heap.Push(hp, pre[i][j])
				}
			}
		}
	}
	return hp.Pop().(int)
}

type minHp []int

func newHeap() *minHp {
	hp := new(minHp)
	heap.Init(hp)
	return hp
}

func (m minHp) Len() int {
	return len(m)
}

func (m minHp) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m minHp) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *minHp) Push(x any) {

	*m = append(*m, x.(int))
}

func (m *minHp) Pop() any {
	x := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return x
}

func (m *minHp) Peek() int {
	return (*m)[0]
}
