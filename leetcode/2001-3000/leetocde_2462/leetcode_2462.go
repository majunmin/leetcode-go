package leetocde_2462

func totalCost(costs []int, k int, candidates int) int64 {

}

type hp [][]int

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i][0] < h[j][0]
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(x any) {
	*h = append(*h, x.([]int))
}

func (h *hp) Pop() any {
	x := (*h)[:len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
