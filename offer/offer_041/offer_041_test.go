package offer_041

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"
)

type hpStruct struct {
	sort.IntSlice
}

func (h *hpStruct) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *hpStruct) Pop() interface{} {
	if h.Len() == 0 {
		return nil
	}
	item := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return item
}

func TestName(t *testing.T) {
	finder := Constructor()
	finder.AddNum(-1)
	finder.AddNum(-2)
	finder.AddNum(-3)
	finder.AddNum(-4)
	finder.AddNum(-5)
	fmt.Println(finder.FindMedian())
}
func TestHeap(t *testing.T) {
	hp := &hpStruct{}
	heap.Init(hp)

	heap.Push(hp, 3)
	heap.Push(hp, 1)
	heap.Push(hp, 2)
	heap.Push(hp, 5)
	heap.Push(hp, 4)
	heap.Push(hp, 10)
	heap.Push(hp, 2)

	fmt.Println(heap.Pop(hp))
	fmt.Println(heap.Pop(hp))
	fmt.Println(heap.Pop(hp))
	fmt.Println(heap.Pop(hp))
	fmt.Println(heap.Pop(hp))
	fmt.Println(heap.Pop(hp))
	fmt.Println(heap.Pop(hp))
}
