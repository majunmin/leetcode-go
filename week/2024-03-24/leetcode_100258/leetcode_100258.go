package leetcode_100258

import "container/heap"

// https://leetcode.cn/contest/weekly-contest-390/problems/most-frequent-ids/
func mostFrequentIDs(nums []int, freq []int) []int64 {
	numsFreq := make(map[int]*pair, 0)
	maxHp := NewHp()
	result := make([]int64, 0, len(nums))
	for i, num := range nums {
		p, exist := numsFreq[num]
		if !exist {
			newP := &pair{
				num:  num,
				freq: freq[i],
				idx:  len(*maxHp),
			}
			heap.Push(maxHp, newP)
			numsFreq[num] = newP
		} else {
			p.freq += freq[i]
			heap.Fix(maxHp, p.idx)
		}
		result = append(result, int64((*maxHp)[0].freq))
	}
	return result
}

type pair struct {
	num  int
	freq int
	idx  int
}
type hp []*pair

func NewHp() *hp {
	maxHp := &hp{}
	heap.Init(maxHp)
	return maxHp

}

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i].freq > h[j].freq
}

func (h hp) Swap(i, j int) {
	h[i].idx, h[j].idx = h[j].idx, h[i].idx
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(x any) {
	*h = append(*h, x.(*pair))
}

func (h *hp) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x

}
