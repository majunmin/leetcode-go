package leetcode_1845

import "container/heap"

// https://leetcode.cn/problems/seat-reservation-manager/?envType=daily-question&envId=2024-09-30
type SeatManager struct {
	mHp heap.Interface
}

func Constructor(n int) SeatManager {
	mHp := newMinHp()
	for i := 1; i <= n; i++ {
		heap.Push(mHp, i)
	}
	return SeatManager{mHp: mHp}

}

func (this *SeatManager) Reserve() int {
	return heap.Pop(this.mHp).(int)
}

func (this *SeatManager) Unreserve(seatNumber int) {
	heap.Push(this.mHp, seatNumber)
}

type minHp []int

func newMinHp() heap.Interface {
	mHp := new(minHp)
	heap.Init(mHp)
	return mHp
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

/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */
