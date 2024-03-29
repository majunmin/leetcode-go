package leetcode_0303

// 	前缀和 or 树状数组

type segment struct {
	left, right int
	// 统计信息
	//count, sum, max, min ...
	sum int
}

type SegmentTree struct {
	m        int
	segments []*segment
}

func NewSegmentTree(size int) *SegmentTree {
	segments := make([]*segment, 4*size+1)
	buildSegments(segments, 0, 0, size-1)
	return &SegmentTree{segments: segments, m: size}
}

func buildSegments(segments []*segment, rootIdx int, l int, r int) {
	// terminate
	segments[rootIdx] = &segment{
		left:  l,
		right: r,
	}
	if l == r {
		return
	}
	mid := l + (r-l)>>1
	buildSegments(segments, 2*rootIdx+1, l, mid)
	buildSegments(segments, 2*rootIdx+2, mid+1, r)
}

func (st *SegmentTree) add(i, data int) {
	idx := 0
	root := st.segments[idx]
	l, r := root.left, root.right
	for l < r {
		st.segments[idx].sum += data
		mid := l + (r-l)>>1
		if i <= mid {
			r = mid
			idx = idx*2 + 1
		} else {
			l = mid + 1
			idx = idx*2 + 2
		}
	}
	// l == r
	st.segments[idx].sum += data
}

func (st *SegmentTree) delete(i, data int) {
	idx := 0
	node := st.segments[idx]
	left, right := node.left, node.right
	for left < right {
		st.segments[idx].sum -= data
		mid := left + (right-left)>>1
		if i <= mid {
			right = mid
			idx = idx*2 + 1
		} else {
			left = mid + 1
			idx = idx*2 + 2
		}
	}
	// left == right
	st.segments[idx].sum -= data
}

func (st *SegmentTree) query(l, r int) int {
	return st.queryInternal(l, r, 0, st.segments[0])

}

func (st *SegmentTree) queryInternal(l int, r int, i int, seg *segment) int {
	// terminate
	if seg.left == l && seg.right == r {
		return seg.sum
	}
	left, right := seg.left, seg.right
	mid := left + (right-left)>>1
	if r <= mid {
		return st.queryInternal(l, r, 2*i+1, st.segments[2*i+1])
	} else if l > mid {
		return st.queryInternal(l, r, 2*i+2, st.segments[2*i+2])
	}
	return st.queryInternal(l, mid, 2*i+1, st.segments[2*i+1]) +
		st.queryInternal(mid+1, r, 2*i+2, st.segments[2*i+2])
}

// https://leetcode.cn/problems/range-sum-query-immutable/
type NumArray struct {
	segTree *SegmentTree
}

func Constructor(nums []int) NumArray {
	segTree := NewSegmentTree(len(nums))
	for i, num := range nums {
		segTree.add(i, num)
	}
	return NumArray{segTree: segTree}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.segTree.query(left, right)
}
