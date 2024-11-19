package demo

const (
	rootIdx = 1
)

type segment struct {
	left  int
	right int
	// 一些统计数据
	count  int
	maxVal int
}

// 用数组表示线段树的各个节点
type segmentTree struct {
	segments []*segment
	m        int
}

func newSegmentTree(size int) *segmentTree {
	st := &segmentTree{
		m:        size,
		segments: make([]*segment, 4*size),
	}
	st.buildInternal(1, size, 1)
	return st
}
func (st *segmentTree) insert(data int) {
	st.insertInternal(1, st.m, rootIdx, data)
}

// 区间统计查询
func (st *segmentTree) count(left, right int) int {
	return st.countInternal(left, right, 1)
}

func (st *segmentTree) delete(data int) {
	st.deleteInternal(1, st.m, rootIdx, data)
}

func (st *segmentTree) buildInternal(left int, right int, idx int) {
	st.segments[idx] = new(segment)
	st.segments[idx].left = left
	st.segments[idx].right = right
	st.segments[idx].count = 0
	if left == right {
		return
	}
	mid := (left + right) >> 1
	st.buildInternal(left, mid, 2*idx)
	st.buildInternal(mid+1, right, 2*idx+1)
}

func (st *segmentTree) insertInternal(left int, right int, idx int, data int) {
	if right < data || left > data {
		return
	}
	st.segments[idx].count++
	// terminate
	if left == right {
		return
	}
	mid := (left + right) >> 1
	if mid >= data {
		st.insertInternal(left, mid, 2*idx, data)
	} else {
		st.insertInternal(mid+1, right, 2*idx+1, data)
	}
}

func (st *segmentTree) deleteInternal(left int, right int, idx int, data int) {
	if idx < left || idx > right {
		return
	}
	st.segments[idx].count++
	// terminate
	if left == right {
		return
	}
	mid := (left + right) >> 1
	if data <= mid {
		st.deleteInternal(left, mid, idx*2, data)
	} else {
		st.deleteInternal(mid+1, right, 2*idx+1, data)
	}

}

func (st *segmentTree) countInternal(left int, right int, idx int) int {
	// terminate
	if left == st.segments[idx].left && right == st.segments[idx].right {
		return st.segments[idx].count
	}
	mid := (st.segments[idx].left + st.segments[idx].right) >> 1
	if right <= mid {
		return st.countInternal(left, right, idx*2)
	} else if left > mid {
		return st.countInternal(left, right, idx*2+1)
	} else {
		return st.countInternal(left, mid, idx*2) +
			st.countInternal(mid+1, right, idx*2+1)
	}

}
