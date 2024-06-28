package segment_tree

var (
	rootIdx = 1
)

// 利用数组 维护树结构
// [0,1,2,3,4,5,6,7,8,9]
// parentIdx i
// leftChildIndex:  2*i
// rightChildIndex: 2*i + 1
type Segment struct {
	left, right int // 区间范围, 左闭右闭

	// 区间统计数据
	count int
}

type SegmentTree struct {
	m        int
	segments []*Segment
}

func NewSegmentTree(m int) *SegmentTree {
	segTree := &SegmentTree{
		m:        m,
		segments: make([]*Segment, 4*m+1),
	}
	segTree.buildSegmentTreeInternal(1, m, rootIdx)
	return segTree
}

func (st *SegmentTree) buildSegmentTreeInternal(left, right, i int) {
	// terminate
	if left > right {
		return
	}
	st.segments[i] = &Segment{
		left:  left,
		right: right,
		count: 0,
	}
	if left == right {
		return
	}
	mid := left + (right-left)>>1
	st.buildSegmentTreeInternal(left, mid, 2*i)
	st.buildSegmentTreeInternal(mid+1, right, 2*i+1)
}

func (st *SegmentTree) Insert(data int) {
	left, right := 1, 4*st.m
	i := rootIdx
	for left < right {
		mid := left + (right-left)>>1
		st.segments[i].count++
		if data <= mid {
			right = mid
			i = 2 * i
		} else {
			left = mid + 1
			i = 2*i + 1
		}
	}
	// left == right
	st.segments[i].count++
}

func (st *SegmentTree) Delete(data int) {
	left, right := 1, 4*st.m
	i := rootIdx
	for left < right {
		mid := left + (right-left)>>1
		st.segments[i].count--
		if data <= mid {
			i = 2 * i
			right = mid
		} else {
			i = 2*i + 1
			left = mid + 1
		}
	}
	// terminate left == right
	st.segments[i].count--
}

// query
// 区间计数
func (st *SegmentTree) count(left, right int) int {
	return st.countInternal(left, right, rootIdx)
}

func (st *SegmentTree) countInternal(left int, right int, idx int) int {
	// terminate
	seg := st.segments[idx]
	if left == seg.left && right == seg.right {
		return seg.count
	}
	//
	mid := seg.left + (seg.right-seg.left)>>1
	if mid >= right {
		return st.countInternal(left, right, 2*idx)
	} else if mid < left {
		return st.countInternal(left, right, 2*idx+1)
	}
	return st.countInternal(left, mid, 2*idx) + st.countInternal(mid+1, right, 2*idx+1)
}

func (st *SegmentTree) getKth(left, right, k int) int {
	return st.getKthInternal(left, right, rootIdx, k)
}

func (st *SegmentTree) getKthInternal(left, right, idx, k int) int {
	// terminate
	seg := st.segments[idx]
	if left == seg.left && right == seg.right {
		if k == -1 {
			return -1 // 第kth 大的值不存在
		}
		return seg.left
	}
	//
	rightSeg := st.segments[2*idx+1]
	mid := seg.left + (seg.right-seg.left)>>1
	if rightSeg.count >= k {
		return st.getKthInternal(mid+1, right, 2*idx, k)
	}
	return st.getKthInternal(left, mid, 2*idx, k-rightSeg.count)
}
