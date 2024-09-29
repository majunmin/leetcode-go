package leetcode_0307

// https://leetcode.cn/problems/range-sum-query-mutable/description/
type NumArray struct {
	st *segmentTree
}

func Constructor(nums []int) NumArray {
	st := newSegmentTree(len(nums))
	for i, num := range nums {
		st.update(i, num)
	}
	return NumArray{st: st}
}

func (this *NumArray) Update(index int, val int) {
	this.st.update(index, val)
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.st.querySum(left, right)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */

type segment struct {
	// 左闭右开区间
	left  int
	right int
	sum   int
	// cnt, max , min
}

type segmentTree struct {
	m        int
	segments []*segment
}

func newSegmentTree(m int) *segmentTree {
	segments := make([]*segment, m*4)
	st := &segmentTree{
		m:        m,
		segments: segments,
	}
	st.build(0, m-1, 0)
	return st
}

func (st *segmentTree) build(left, right, idx int) {
	st.segments[idx] = &segment{
		left:  left,
		right: right,
	}
	if left == right {
		return
	}
	mid := left + (right-left)>>1
	st.build(left, mid, 2*idx+1)
	st.build(mid+1, right, 2*idx+2)
}

func (st *segmentTree) querySum(left, right int) int {
	return st.querySumInternal(left, right, 0)
}

func (st *segmentTree) querySumInternal(left int, right int, i int) int {
	seg := st.segments[i]
	if seg.left == left && seg.right == right {
		return seg.sum
	}
	mid := seg.left + (seg.right-seg.left)>>1
	if left > mid {
		return st.querySumInternal(left, right, i*2+2)
	} else if right <= mid {
		return st.querySumInternal(left, right, i*2+1)
	} else {
		return st.querySumInternal(left, mid, i*2+1) + st.querySumInternal(mid+1, right, i*2+2)
	}
}

func (st *segmentTree) update(idx int, val int) {
	st.change(idx, val, 0, 0, st.m-1)
}

func (st *segmentTree) change(idx int, val int, node int, left, right int) {
	// terminate
	if left == right {
		st.segments[node].sum = val
		return
	}
	mid := left + (right-left)>>1
	if idx <= mid {
		st.change(idx, val, node*2+1, left, mid)
	} else {
		st.change(idx, val, node*2+2, mid+1, right)
	}
	st.segments[node].sum = st.segments[node*2+1].sum + st.segments[node*2+2].sum
}
