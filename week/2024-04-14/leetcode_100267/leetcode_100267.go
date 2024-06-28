package leetcode_100267

import "container/heap"

// 利用小顶堆
// 利用map 去重
// https://leetcode.cn/contest/weekly-contest-393/problems/kth-smallest-amount-with-single-denomination-combination/
func findKthSmallest(coins []int, k int) int64 {
	var (
		visited = make(map[int]bool) //记录已经出现过的面值
		minHeap = newMinHeap()
	)

	// init minHeap
	for i, coin := range coins {
		minHeap.offer(&item{
			value:   coin,
			coinIdx: i,
			idx:     i,
		})
		visited[coin] = true
	}
	//
	for k-1 > 0 {
		val := minHeap.top()
		//val.value += coins[val.coinIdx]
		for visited[val.value+coins[val.coinIdx]] {
			val.value += coins[val.coinIdx]
		}
		val.value += coins[val.coinIdx]
		heap.Fix(minHeap, val.idx)
		visited[val.value] = true
		k--
	}
	return int64(minHeap.top().value)
}

type item struct {
	value   int
	idx     int
	coinIdx int
}

type minHp []*item

func newMinHeap() *minHp {
	m := new(minHp)
	heap.Init(m)
	return m
}

func (m minHp) Len() int {
	return len(m)
}

func (m minHp) Less(i, j int) bool {
	return m[i].value < m[j].value
}

func (m minHp) Swap(i, j int) {
	m[i].idx, m[j].idx = m[j].idx, m[i].idx
	m[i], m[j] = m[j], m[i]
}

func (m *minHp) Push(x any) {
	*m = append(*m, x.(*item))
}

func (m *minHp) Pop() any {
	x := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return x
}

func (m *minHp) offer(x *item) {
	heap.Push(m, x)
}

func (m *minHp) poll() *item {
	return heap.Pop(m).(*item)
}
func (m *minHp) top() *item {
	return (*m)[0]
}
