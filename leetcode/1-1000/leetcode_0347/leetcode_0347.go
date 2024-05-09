package leetcode_0347

import "container/heap"

type node struct {
	num  int
	freq int
}

type qHeap []*node

func (q qHeap) Len() int {
	return len(q)
}

func (q qHeap) Less(i, j int) bool {
	return q[i].freq > q[j].freq
}

func (q qHeap) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *qHeap) Push(x any) {
	*q = append(*q, x.(*node))
}

func (q *qHeap) Pop() any {
	x := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return x
}

func (q *qHeap) push(x any) {
	heap.Push(q, x)
}

func (q *qHeap) pop() any {
	return heap.Pop(q)
}

// https://leetcode.cn/problems/top-k-frequent-elements/?envType=study-plan-v2&envId=top-100-liked
func topKFrequent(nums []int, k int) []int {
	var result []int
	if k == 0 {
		return result
	}
	if len(nums) < k {
		panic("invalid param")
	}
	//
	numCnt := make(map[int]*node, 0)
	for _, num := range nums {
		if _, exist := numCnt[num]; !exist {
			numCnt[num] = &node{
				num:  num,
				freq: 0,
			}
		}
		numCnt[num].freq++
	}

	pq := new(qHeap)
	heap.Init(pq)
	for _, n := range numCnt {
		pq.push(n)
	}
	for i := 0; i < k; i++ {
		item := pq.pop().(*node)
		result = append(result, item.num)
	}
	return result
}
