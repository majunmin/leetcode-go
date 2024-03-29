package leetcode_0295

import "container/heap"

type minHeap []int

func newMinHeap() *minHeap {
	minHp := new(minHeap)
	heap.Init(minHp)
	return minHp
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

func (m *minHeap) push(x any) {
	heap.Push(m, x)
}

func (m *minHeap) pop() any {
	return heap.Pop(m)
}
func (m *minHeap) top() any {
	return (*m)[0]
}

type maxHeap []int

func newMaxHeap() *maxHeap {
	maxHp := new(maxHeap)
	heap.Init(maxHp)
	return maxHp
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

func (m *maxHeap) push(x any) {
	heap.Push(m, x)
}

func (m *maxHeap) pop() any {
	return heap.Pop(m)
}

func (m *maxHeap) top() any {
	return (*m)[0]
}

// https://leetcode.cn/problems/find-median-from-data-stream/?envType=study-plan-v2&envId=top-100-liked
type MedianFinder struct {
	left  *maxHeap
	right *minHeap
	size  int
}

func Constructor() MedianFinder {
	return MedianFinder{
		left:  newMaxHeap(),
		right: newMinHeap(),
	}
}

// 因为num 是有序数字， 所以
func (this *MedianFinder) AddNum(num int) {
	llen, rlen := this.left.Len(), this.right.Len()
	if llen == rlen {
		this.right.push(num)
		this.left.push(this.right.pop())
	} else {
		this.left.push(num)
		this.right.push(this.left.pop())
	}
	this.size++
}

func (this *MedianFinder) FindMedian() float64 {
	if this.size == 0 {
		return float64(0)
	}
	if this.size&1 == 1 {
		return float64(this.left.top().(int))
	} else {
		return float64(this.left.top().(int)+this.right.top().(int)) / 2.0
	}
}
