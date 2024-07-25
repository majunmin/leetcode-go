package leetcode_0295

import "container/heap"

// https://leetcode.cn/problems/find-median-from-data-stream/?envType=study-plan-v2&envId=top-interview-150
type MedianFinder struct {
	// 用一个大顶堆和一个小顶堆来实现
	leftHp  *maxHeap
	rightHp *minHeap
	cnt     int // 元素数
}

func Constructor() MedianFinder {
	return MedianFinder{
		leftHp:  newMaxHeap(),
		rightHp: newMinHeap(),
	}
}

func (this *MedianFinder) AddNum(num int) {
	defer func() {
		this.cnt++
	}()

	if this.leftHp.Len() == 0 {
		heap.Push(this.leftHp, num)
		return
	}
	if this.leftHp.Len() == this.rightHp.Len() {
		if (*this.rightHp)[0] < num {
			heap.Push(this.leftHp, heap.Pop(this.rightHp))
			heap.Push(this.rightHp, num)
		} else {
			heap.Push(this.leftHp, num)
		}
		return
	}
	if (*this.leftHp)[0] > num {
		heap.Push(this.rightHp, heap.Pop(this.leftHp))
		heap.Push(this.leftHp, num)
		return
	}
	heap.Push(this.rightHp, num)
}

func (this *MedianFinder) FindMedian() float64 {
	if this.cnt == 0 {
		return 0
	}
	if this.cnt&1 == 1 {
		return float64((*this.leftHp)[0])
	}
	return float64((*this.leftHp)[0]+(*this.rightHp)[0]) / 2.0
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

type minHeap []int

func newMinHeap() *minHeap {
	h := new(minHeap)
	heap.Init(h)
	return h
}

func (m minHeap) Len() int {
	return len(m)
}

func (m minHeap) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *minHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *minHeap) Pop() any {
	x := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return x
}

type maxHeap []int

func newMaxHeap() *maxHeap {
	h := new(maxHeap)
	heap.Init(h)
	return h
}

func (m maxHeap) Len() int {
	return len(m)
}

func (m maxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m maxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *maxHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *maxHeap) Pop() any {
	x := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return x
}
