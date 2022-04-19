// Created By: junmin.ma
// Description: <description>
// Date: 2022-03-26 15:55
package heap

type BinaryHeap struct {
	d        int // 度
	heap     []int
	heapSize int
}

func NewBinaryHeap(heapSize, d int) *BinaryHeap {
	return &BinaryHeap{
		d:        d,
		heap:     make([]int, heapSize),
		heapSize: heapSize,
	}
}

func (bh *BinaryHeap) insert(item int) {
	if bh.isFull() {
		return
	}
	bh.heap[bh.heapSize] = item
	bh.heapSize++
	bh.heapifyUp(bh.heapSize)

}

func (bh *BinaryHeap) delete(idx int) {
	if bh.isEmpty() {
		return
	}
	bh.heapSize--
	bh.swap(bh.heapSize, idx)
	bh.heapifyDown(idx)

}

func (bh *BinaryHeap) peek() int {
	if bh.isEmpty() {
		return -1
	}
	return bh.heap[0]
}

func (bh *BinaryHeap) isFull() bool {
	return bh.heapSize == len(bh.heap)
}

func (bh *BinaryHeap) isEmpty() bool {
	return bh.heapSize == 0
}

func (bh *BinaryHeap) heapifyUp(idx int) {
	if idx <= 0 {
		return
	}
	pIdx := bh.parent(idx)
	if bh.heap[idx] > bh.heap[pIdx] {
		bh.swap(idx, pIdx)
		bh.heapifyUp(pIdx)
	}
}

func (bh *BinaryHeap) parent(child int) int {
	return (child - 1) / bh.d
}

func (bh *BinaryHeap) child(parent, idx int) int {
	return parent*bh.d + idx
}

func (bh *BinaryHeap) swap(i int, j int) {
	bh.heap[i], bh.heap[j] = bh.heap[j], bh.heap[i]
}

func (bh *BinaryHeap) heapifyDown(idx int) {
	if idx < 0 {
		return
	}
	for i := 1; i <= bh.d; i++ {
		childIdx := bh.child(idx, i)
		//剪枝
		if childIdx >= bh.heapSize {
			break
		}
		if bh.heap[idx] < childIdx {
			bh.swap(idx, childIdx)
			bh.heapifyDown(childIdx)
		}
	}
}
