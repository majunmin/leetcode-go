package tree

var (
	rootIndex = 1
)

// index = 1 开始编号
type Segment struct {
	left  int // 区间起始点
	right int // 区间 结束点
	count int // 统计值
}

type SegmentTree struct {
	m        int
	segments []*Segment
}

func NewSegmentTree(m int) *SegmentTree {
	return &SegmentTree{
		m:        m,
		segments: make([]*Segment, 4*m),
	}
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
	mid := left + (right-left)/2
	st.buildSegmentTreeInternal(left, mid, i*2)
	st.buildSegmentTreeInternal(mid+1, right, i*2+1)
}

func (st *SegmentTree) insert(data int) {
	left, right := 1, st.m
	i := rootIndex
	for left < right {
		mid := left + (right - left) + 1
		st.segments[i].count++
		if data <= mid {
			right = mid
			i = 2 * i
		} else {
			left = mid + 1
			i = 2*i + 1
		}
	}
	st.segments[i].count++
}

func (st *SegmentTree) delete(data int) {
	left, right := 1, st.m
	i := rootIndex

	for left < right {
		st.segments[i].count--
		mid := left + (right-left)/2
		if data <= mid {
			right = mid
			i = 2 * i
		} else {
			left = mid + 1
			i = 2*i + 1
		}
	}
	st.segments[i].count--
}

// query
func (st *SegmentTree) count(left, right int) int {
	return st.countInternal(left, right, rootIndex)
}

func (st *SegmentTree) countInternal(left int, right int, index int) int {
	// terminate
	seg := st.segments[index]
	if seg.left == left && seg.right == right {
		return seg.count
	}

	mid := seg.left + (seg.right-seg.left)>>1
	if mid >= right {
		//区间在左子节点
		return st.countInternal(left, right, 2*index)
	} else if mid < left {
		// 区间在右子节点
		return st.countInternal(left, right, 2*index+1)
	} else {
		return st.countInternal(left, mid, 2*index) +
			st.countInternal(mid+1, right, 2*index+1)
	}

}

func (st *SegmentTree) getKth(left, right, k int) int {
	return st.getKthInternal(left, right, rootIndex, k)
}

func (st *SegmentTree) getKthInternal(left int, right int, index int, k int) int {
	seg := st.segments[index]
	if left == seg.left && right == seg.right {
		if k == -1 {
			//第 Kth 大值不存在
			return -1
		} else {
			return seg.left
		}
	}
	rightSeg := st.segments[2*index+1]
	mid := left + (right-left)/2
	if rightSeg.count > k {
		return st.getKthInternal(mid+1, right, 2*index+1, k)
	} else {
		return st.getKthInternal(left, mid, 2*index, k-rightSeg.count)
	}
}
