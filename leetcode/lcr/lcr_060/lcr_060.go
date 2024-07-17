package lcr_060

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	// 堆来实现
	var (
		freqs  = make(map[int]int, 0)
		maxHp  = newHp()
		result = make([]int, 0, k)
	)
	for _, num := range nums {
		freqs[num]++
	}
	for num, freq := range freqs {
		heap.Push(maxHp, pair{num: num, freq: freq})
	}
	if k > maxHp.Len() {
		return nil
	}
	for i := 0; i < k; i++ {
		p := heap.Pop(maxHp).(pair)
		result = append(result, p.num)
	}
	return result
}

type pair struct {
	num  int
	freq int
}

type hp []pair

func newHp() *hp {
	h := new(hp)
	heap.Init(h)
	return h
}

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i].freq > h[j].freq
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(x any) {
	*h = append(*h, x.(pair))
}

func (h *hp) Pop() any {
	x := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return x
}
