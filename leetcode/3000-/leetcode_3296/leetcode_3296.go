package leetcode_3296

import "container/heap"

// https://leetcode.cn/problems/minimum-number-of-seconds-to-make-mountain-height-zero/
func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	var (
		mHp    = newMinHp()
		result int
	)
	for _, item := range workerTimes {
		heap.Push(mHp, worker{
			next: item,
			delt: item,
			base: item,
		})
	}
	for i := 0; i < mountainHeight; i++ {
		top := (*mHp)[0]
		result = max(result, top.next)
		top.delt += top.base
		top.next += top.delt
		heap.Fix(mHp, 0)
	}
	return int64(result)
}

type worker struct {
	next int
	delt int
	base int
}

type minHp []worker

func newMinHp() *minHp {
	hp := new(minHp)
	heap.Init(hp)
	return hp
}

func (m minHp) Len() int {
	return len(m)
}

func (m minHp) Less(i, j int) bool {
	return m[i].next < m[j].next
}

func (m minHp) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *minHp) Push(x any) {
	*m = append(*m, x.(worker))
}

func (m *minHp) Pop() any {
	w := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return w
}
