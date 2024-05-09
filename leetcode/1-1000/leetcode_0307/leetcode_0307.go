package leetcode_0307

// 线段树 节点
type SegmentTree struct {
	tree []int
	data []int
}

// [start,end] 左右是闭区间
func (t *SegmentTree) buildTree(data []int, idx int, start int, end int) int {
	if start == end {
		t.tree[idx] = data[start]
		return data[start]
	}
	mid := start + (end-start)>>1
	lVal := t.buildTree(data, idx*2+1, start, mid)
	rVal := t.buildTree(data, idx*2+2, mid+1, end)
	t.tree[idx] = lVal + rVal
	return t.tree[idx]
}

func NewSegmentTree(data []int) *SegmentTree {
	if len(data) == 0 {
		panic("invalid param")
	}
	st := &SegmentTree{
		data: data,
		tree: make([]int, 4*len(data)),
	}
	st.buildTree(data, 0, 0, len(data)-1)
	return st
}

func (st *SegmentTree) Update(idx, val int) {
	if idx >= len(st.data) {
		//panic("invalid param")
	}
	st.data[idx] = val
	st.update(0, idx, val, 0, len(st.data)-1)
}

func (st *SegmentTree) Sum(left, right int) int {
	// param check
	if left < 0 || left > len(st.data) || right < 0 || right > len(st.data) || left > right {
		panic("invalid param")
	}
	return st.sum(0, 0, len(st.data)-1, left, right)
}

func (t *SegmentTree) update(treeIdx, idx, val, start, end int) int {
	if start == end { // mid == start == end
		t.tree[treeIdx] = val
		return val
	}
	mid := start + (end-start)>>1
	var leftVal, rightVal int
	if idx <= mid {
		leftVal = t.update(2*treeIdx+1, idx, val, start, mid)
		rightVal = t.tree[treeIdx*2+2]
	} else {
		leftVal = t.tree[treeIdx*2+1]
		rightVal = t.update(2*treeIdx+2, idx, val, mid+1, end)
	}
	t.tree[treeIdx] = leftVal + rightVal
	return t.tree[treeIdx]
}

func (t *SegmentTree) sum(idx int, left int, right int, ql int, qr int) int {
	if ql == left && qr == right {
		return t.tree[idx]
	}
	//
	mid := left + (right-left)>>1
	// 查询右区间
	if ql > mid {
		return t.sum(idx*2+2, mid+1, right, ql, qr)
	}
	// 查询左区间
	if qr <= mid {
		return t.sum(idx*2+1, left, mid, ql, qr)
	}
	return t.sum(idx*2+1, left, mid, ql, mid) +
		t.sum(idx*2+2, mid+1, right, mid+1, qr)
}

// https://leetcode.cn/problems/range-sum-query-mutable/description/
type NumArray struct {
	// 因为构建出的  线段树是一个完全二叉树， 所以这里用一个数组表示
	tree *SegmentTree
}

func Constructor(nums []int) NumArray {
	return NumArray{
		tree: NewSegmentTree(nums),
	}
}

func (this *NumArray) Update(index int, val int) {
	this.tree.Update(index, val)
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.tree.Sum(left, right)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
