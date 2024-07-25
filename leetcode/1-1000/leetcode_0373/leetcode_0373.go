package leetcode_0373

import "container/heap"

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var (
		result = make([][]int, k)
		hp     = newMinHeap()
	)
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			heap.Push(hp, [2]int{n1, n2})
		}
	}
	// add to res
	for i := 0; i < k; i++ {
		result = append(result, heap.Pop(hp).([]int))
	}
	return result

}

type pair struct {
	i, j int
}

type minHeap []pair

func newMinHeap() *minHeap {
	hp := new(minHeap)
	heap.Init(hp)
	return hp
}

func (m minHeap) Len() int {
	return len(m)
}

func (m minHeap) Less(i, j int) bool {
	return m[i][0]+m[i][1] < m[j][0]+m[j][1]
}

func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *minHeap) Push(x any) {
	*m = append(*m, x.(pair))
}

func (m *minHeap) Pop() any {
	x := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return x
}
